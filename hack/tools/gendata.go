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

package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"

	"kubefin.dev/go-tools/awspricing"
)

const (
	DefaultRetryCount = 6
)

// Usage:
// go run hack/tools/gendata.go

func main() {
	for region := range awspricing.ReginList {
		url := fmt.Sprintf(awspricing.GlobalPriceDataBaseURL + region + "/index.json")
		if strings.HasPrefix(region, "cn-") {
			url = fmt.Sprintf(awspricing.CNPriceDataBaseURL + region + "/index.json")
		}

		getDataFunc := func() error {
			resp, err := http.Get(url)
			if err != nil {
				return err
			}
			defer resp.Body.Close()

			respData, err := io.ReadAll(resp.Body)
			if err != nil {
				return err
			}

			ec2Data, err := awspricing.ParseData(respData)
			if err != nil {
				return err
			}

			generalData, err := awspricing.ConvertToGeneralPriceData(ec2Data)
			if err != nil {
				return err
			}

			regionData := map[string][]awspricing.EC2GeneralPrice{}
			for _, item := range *generalData {
				if _, ok := regionData[item.Region]; !ok {
					regionData[item.Region] = []awspricing.EC2GeneralPrice{}
				}
				regionData[item.Region] = append(regionData[item.Region], item)
			}

			for region, data := range regionData {
				jsonRet, err := json.Marshal(data)
				if err != nil {
					return err
				}

				err = os.WriteFile("awspricing/data/"+region+".json", jsonRet, 0644)
				if err != nil {
					return err
				}
				fmt.Printf("EC2 price data(%s) download/parse finished....\n", region)
			}

			return nil
		}

		for i := 0; i < DefaultRetryCount; i++ {
			if err := getDataFunc(); err != nil {
				fmt.Println("Failed:", err.Error())
				time.Sleep(time.Second * 30)
				continue
			}
			break
		}
	}
}
