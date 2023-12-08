/*
Copyright 2023 The KubeFin Authors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package awspricing

import (
	"embed"
	"encoding/json"
	"fmt"
	"io"
	"io/fs"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"

	"k8s.io/klog/v2"
)

//go:embed data/*.json
var file embed.FS

type AWSEC2PriceClient struct {
	latestPriceDataTime time.Time
	dataLock            sync.Mutex
	priceData           *GeneralPriceData
	priceDataURL        string
}

func NewAWSEC2PriceClient(region string) (*AWSEC2PriceClient, error) {
	if region == "" {
		return nil, fmt.Errorf("parameter region is needed")
	}

	client := &AWSEC2PriceClient{
		priceDataURL: GlobalPriceDataBaseURL + region + "/index.json",
		priceData:    &GeneralPriceData{},
		dataLock:     sync.Mutex{},
	}

	if _, ok := ReginList[region]; !ok {
		return nil, fmt.Errorf("unknown region: %s", region)
	}
	lastUpdatedTime, err := time.Parse(time.RFC1123, ReginList[region])
	if err != nil {
		return nil, err
	}
	client.latestPriceDataTime = lastUpdatedTime

	if strings.HasPrefix(region, "cn-") {
		client.priceDataURL = CNPriceDataBaseURL + region + "/index.json"
	}

	data, err := fs.ReadFile(file, "data/"+region+".json")
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(data, client.priceData)
	if err != nil {
		return nil, err
	}

	return client, nil
}

func (a *AWSEC2PriceClient) Start() {
	for range time.Tick(time.Second) {
		a.dataLock.Lock()
		a.updatePriceData()
		a.dataLock.Unlock()
	}
}

func (a *AWSEC2PriceClient) updatePriceData() {
	latestUpdateTime, err := getPriceDataLatestUpdateTime(a.priceDataURL)
	if err != nil {
		klog.Errorf("get latest price data time failed: %v", err)
		return
	}
	if err == nil && !latestUpdateTime.Equal(a.latestPriceDataTime) {
		klog.Infof("EC2 price data updated, start downloading the latest data....")
		resp, err := http.Get(a.priceDataURL)
		if err != nil {
			klog.Errorf("get price data failed: %v", err)
			return
		}
		klog.Infof("EC2 price data downloading finished....")
		defer resp.Body.Close()

		data, err := io.ReadAll(resp.Body)
		if err != nil {
			klog.Errorf("read price data failed: %v", err)
			return
		}
		awsData, err := ParseData(data)
		if err != nil {
			klog.Errorf("parse price data failed: %v", err)
			return
		}
		generalData, err := ConvertToGeneralPriceData(awsData)
		if err != nil {
			klog.Errorf("convert price data failed: %v", err)
			return
		}
		a.priceData = generalData
		a.latestPriceDataTime = latestUpdateTime
	}
}

func (a *AWSEC2PriceClient) GetOnDemandEC2PriceInfo(instanceType string) (*EC2GeneralPrice, error) {
	if instanceType == "" {
		return nil, fmt.Errorf("invalid input, parameter instanceType is needed")
	}

	a.dataLock.Lock()
	defer a.dataLock.Unlock()
	for _, item := range *a.priceData {
		//return &item, nil
		if item.InstanceType == instanceType {
			return &item, nil
		}
	}

	return nil, fmt.Errorf("no matched price info found, instanceType: %s", instanceType)
}

func (a *AWSEC2PriceClient) ListEC2PriceInfo() []EC2GeneralPrice {
	ret := []EC2GeneralPrice{}

	a.dataLock.Lock()
	defer a.dataLock.Unlock()
	ret = append(ret, *a.priceData...)

	return ret
}

func getPriceDataLatestUpdateTime(priceDataUrl string) (time.Time, error) {
	req, err := http.NewRequest("HEAD", priceDataUrl, nil)
	if err != nil {
		return time.Time{}, err
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return time.Time{}, err
	}
	defer resp.Body.Close()

	t := resp.Header.Get("Last-Modified")
	if t == "" {
		return time.Time{}, fmt.Errorf("no last modified time")
	}

	ret, err := time.Parse(time.RFC1123, t)
	if err != nil {
		return time.Time{}, err
	}

	return ret, nil
}

func ParseData(data []byte) (*AWSEC2Data, error) {
	ec2Data := &AWSEC2Data{}
	if err := json.Unmarshal(data, ec2Data); err != nil {
		return nil, err
	}

	return ec2Data, nil
}

func ConvertToGeneralPriceData(data *AWSEC2Data) (*GeneralPriceData, error) {
	generalData := GeneralPriceData{}
	for id, resourceItem := range data.Products {
		if resourceItem.Attributes.InstanceType == "" {
			continue
		}

		// In most scenarios, there is shared VM
		if resourceItem.Attributes.Tenancy != Shared {
			continue
		}
		// We only care about Linux VM price
		if resourceItem.Attributes.OperatingSystem != "Linux" {
			continue
		}

		// This means it installed other sf
		if resourceItem.Attributes.InstanceSKU != "" {
			continue
		}

		// Following code is referring to: https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/billing-info-fields.html
		// We only care about standard Linux
		if resourceItem.Attributes.Operation != "RunInstances" {
			continue
		}

		// Following code is referring to: https://docs.aws.amazon.com/cost-management/latest/userguide/ce-filtering.html
		usageType := strings.Split(resourceItem.Attributes.UsageType, "-")
		if len(usageType) < 2 {
			continue
		}
		usageTypeWithOutInstanceType := strings.Split(usageType[1], ":")
		if len(usageTypeWithOutInstanceType) < 2 || (usageTypeWithOutInstanceType[0] != "BoxUsage" && usageTypeWithOutInstanceType[0] != "Reservation") {
			continue
		}

		cpu, err := strconv.ParseFloat(resourceItem.Attributes.VCPU, 64)
		if err != nil {
			cpu = 0.0
		}
		memory, err := extractMemory(resourceItem.Attributes.Memory)
		if err != nil {
			memory = 0.0
		}
		item := EC2GeneralPrice{
			InstanceType: resourceItem.Attributes.InstanceType,
			Region:       resourceItem.Attributes.RegionCode,
			VCPU:         cpu,
			Memory:       memory,
		}

		item.PriceModel = PriceModelOnDemand
		for key, priceItem := range data.Terms.OnDemand[id] {
			if !strings.HasPrefix(key, id) {
				continue
			}
			for _, priceDimensions := range priceItem.PriceDimensions {
				item.PriceUnit = priceDimensions.Unit
				item.PricePerUnit = priceDimensions.PricePerUnit
				generalData = append(generalData, item)
			}
		}
	}

	return &generalData, nil
}

func extractMemory(memory string) (float64, error) {
	s := strings.TrimSuffix(memory, " GiB")
	return strconv.ParseFloat(s, 64)
}
