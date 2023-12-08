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

const (
	CNPriceDataBaseURL     = "https://pricing.cn-north-1.amazonaws.com.cn/offers/v1.0/cn/AmazonEC2/current/"
	GlobalPriceDataBaseURL = "https://pricing.us-east-1.amazonaws.com/offers/v1.0/aws/AmazonEC2/current/"
)

var ReginList = map[string]string{
	"ap-southeast-1-bkk-1":    "Fri, 17 Nov 2023 17:44:44 GMT",
	"af-south-1-los-1":        "Fri, 17 Nov 2023 17:44:44 GMT",
	"eu-north-1-hel-1":        "Fri, 17 Nov 2023 17:44:44 GMT",
	"us-west-2-sea-1":         "Fri, 17 Nov 2023 17:44:45 GMT",
	"us-west-2-las-1":         "Fri, 17 Nov 2023 17:44:45 GMT",
	"us-east-1-dfw-1":         "Fri, 17 Nov 2023 17:44:44 GMT",
	"eu-central-2":            "Fri, 17 Nov 2023 17:44:44 GMT",
	"eu-south-1":              "Fri, 17 Nov 2023 17:44:44 GMT",
	"us-east-1-phl-1":         "Fri, 17 Nov 2023 17:44:44 GMT",
	"us-east-1-wl1-msp1":      "Fri, 17 Nov 2023 17:44:45 GMT",
	"us-west-2":               "Fri, 17 Nov 2023 17:44:46 GMT",
	"ap-southeast-3":          "Fri, 17 Nov 2023 17:44:44 GMT",
	"eu-west-3":               "Fri, 17 Nov 2023 17:44:44 GMT",
	"ap-south-1-ccu-1":        "Fri, 17 Nov 2023 17:44:44 GMT",
	"us-west-2-wl1-sea1":      "Fri, 17 Nov 2023 17:44:46 GMT",
	"eu-west-2-wl1-man1":      "Fri, 17 Nov 2023 17:44:44 GMT",
	"us-east-1-wl1-iah1":      "Fri, 17 Nov 2023 17:44:44 GMT",
	"ap-northeast-2":          "Fri, 17 Nov 2023 17:44:44 GMT",
	"us-west-2-phx-1":         "Fri, 17 Nov 2023 17:44:46 GMT",
	"ap-northeast-1":          "Fri, 17 Nov 2023 17:44:44 GMT",
	"us-east-1-msp-1 ":        "Fri, 17 Nov 2023 17:44:44 GMT",
	"ap-southeast-2-per-1":    "Fri, 17 Nov 2023 17:44:44 GMT",
	"ap-northeast-1-tpe-1":    "Fri, 17 Nov 2023 17:44:44 GMT",
	"us-east-1-chi-1":         "Fri, 17 Nov 2023 17:44:44 GMT",
	"eu-central-1-wl1-ber1":   "Fri, 17 Nov 2023 17:44:44 GMT",
	"us-gov-east-1":           "Fri, 17 Nov 2023 17:44:45 GMT",
	"eu-west-2":               "Fri, 17 Nov 2023 17:44:44 GMT",
	"us-west-2-pdx-1":         "Fri, 17 Nov 2023 17:44:46 GMT",
	"us-east-1-wl1-nyc1":      "Fri, 17 Nov 2023 17:44:45 GMT",
	"us-east-1-atl-1":         "Fri, 17 Nov 2023 17:44:44 GMT",
	"us-east-1-wl1-dfw1":      "Fri, 17 Nov 2023 17:44:44 GMT",
	"eu-central-1-wl1-dtm1":   "Fri, 17 Nov 2023 17:44:44 GMT",
	"eu-west-2-wl1-lon1":      "Fri, 17 Nov 2023 17:44:44 GMT",
	"us-gov-west-1":           "Fri, 17 Nov 2023 17:44:45 GMT",
	"ca-central-1-wl1-yto1":   "Fri, 17 Nov 2023 17:44:44 GMT",
	"ap-southeast-2-akl-1":    "Fri, 17 Nov 2023 17:44:44 GMT",
	"eu-central-1-waw-1":      "Fri, 17 Nov 2023 17:44:44 GMT",
	"us-east-1-wl1-atl1":      "Fri, 17 Nov 2023 17:44:44 GMT",
	"ap-northeast-2-wl1-sel1": "Fri, 17 Nov 2023 17:44:44 GMT",
	"us-west-2-wl1-phx1":      "Fri, 17 Nov 2023 17:44:48 GMT",
	"us-west-2-wl1-las1":      "Fri, 17 Nov 2023 17:44:45 GMT",
	"eu-west-1":               "Fri, 17 Nov 2023 17:44:44 GMT",
	"eu-north-1":              "Fri, 17 Nov 2023 17:44:44 GMT",
	"ap-southeast-4":          "Fri, 17 Nov 2023 17:44:44 GMT",
	"ap-northeast-2-wl1-cjj1": "Fri, 17 Nov 2023 17:44:44 GMT",
	"us-east-1-wl1-was1":      "Fri, 17 Nov 2023 17:44:45 GMT",
	"eu-central-1-ham-1":      "Fri, 17 Nov 2023 17:44:44 GMT",
	"eu-south-2":              "Fri, 17 Nov 2023 17:44:47 GMT",
	"ap-northeast-3":          "Fri, 17 Nov 2023 17:44:44 GMT",
	"us-west-2-wl1-den1":      "Fri, 17 Nov 2023 17:44:45 GMT",
	"us-west-2-lax-1":         "Fri, 17 Nov 2023 17:44:45 GMT",
	"ap-northeast-1-wl1-kix1": "Fri, 17 Nov 2023 17:44:44 GMT",
	"eu-north-1-cph-1":        "Fri, 17 Nov 2023 17:44:44 GMT",
	"us-east-1-mia-1":         "Fri, 17 Nov 2023 17:44:44 GMT",
	"ap-northeast-1-wl1-nrt1": "Fri, 17 Nov 2023 17:44:44 GMT",
	"us-east-2":               "Fri, 17 Nov 2023 17:44:45 GMT",
	"me-south-1":              "Fri, 17 Nov 2023 17:44:44 GMT",
	"il-central-1":            "Fri, 17 Nov 2023 17:44:44 GMT",
	"us-east-1-bue-1":         "Fri, 17 Nov 2023 17:44:44 GMT",
	"us-east-1-bos-1":         "Fri, 17 Nov 2023 17:44:44 GMT",
	"us-east-1-nyc-1":         "Fri, 17 Nov 2023 17:44:44 GMT",
	"us-east-1-wl1-clt1":      "Fri, 17 Nov 2023 17:44:44 GMT",
	"sa-east-1":               "Fri, 17 Nov 2023 17:44:44 GMT",
	"ap-east-1":               "Fri, 17 Nov 2023 17:44:44 GMT",
	"us-east-1":               "Fri, 17 Nov 2023 17:44:45 GMT",
	"us-east-1-wl1":           "Fri, 17 Nov 2023 17:44:45 GMT",
	"us-west-2-wl1-lax1":      "Fri, 17 Nov 2023 17:44:45 GMT",
	"us-east-1-qro-1":         "Fri, 17 Nov 2023 17:44:44 GMT",
	"us-west-2-den-1":         "Fri, 17 Nov 2023 17:44:45 GMT",
	"us-east-1-wl1-tpa1":      "Fri, 17 Nov 2023 17:44:45 GMT",
	"ap-south-1-del-1":        "Fri, 17 Nov 2023 17:44:44 GMT",
	"eu-central-1":            "Fri, 17 Nov 2023 17:44:44 GMT",
	"me-central-1":            "Fri, 17 Nov 2023 17:44:44 GMT",
	"us-east-1-lim-1":         "Fri, 17 Nov 2023 17:44:44 GMT",
	"us-east-1-scl-1":         "Fri, 17 Nov 2023 17:44:45 GMT",
	"me-south-1-mct-1":        "Fri, 17 Nov 2023 17:44:44 GMT",
	"us-east-1-wl1-bna1":      "Fri, 17 Nov 2023 17:44:44 GMT",
	"us-east-1-iah-1":         "Fri, 17 Nov 2023 17:44:44 GMT",
	"us-east-1-wl1-chi1":      "Fri, 17 Nov 2023 17:44:44 GMT",
	"cn-northwest-1":          "Sat, 18 Nov 2023 01:05:01 GMT",
	"ap-southeast-1":          "Fri, 17 Nov 2023 17:44:44 GMT",
	"us-east-1-mci-1":         "Fri, 17 Nov 2023 17:44:45 GMT",
	"eu-central-1-wl1-muc1":   "Fri, 17 Nov 2023 17:44:44 GMT",
	"us-east-1-wl1-mia1":      "Fri, 17 Nov 2023 17:44:44 GMT",
	"ap-south-1":              "Fri, 17 Nov 2023 17:44:44 GMT",
	"us-west-1":               "Fri, 17 Nov 2023 17:44:45 GMT",
	"ap-south-2":              "Fri, 17 Nov 2023 17:44:44 GMT",
	"af-south-1":              "Fri, 17 Nov 2023 17:44:44 GMT",
	"us-west-2-wl1":           "Fri, 17 Nov 2023 17:44:46 GMT",
	"cn-north-1":              "Sat, 18 Nov 2023 01:05:01 GMT",
	"ca-central-1":            "Fri, 17 Nov 2023 17:44:44 GMT",
	"ap-southeast-2":          "Fri, 17 Nov 2023 17:44:44 GMT",
	"ap-southeast-1-mnl-1":    "Fri, 17 Nov 2023 17:44:44 GMT",
	"us-east-1-wl1-dtw1":      "Fri, 17 Nov 2023 17:44:44 GMT",
}

type AWSEC2Data struct {
	OfferCode string             `json:"offerCode"`
	Products  map[string]Product `json:"products"`
	Terms     Term               `json:"terms"`
}

type Product struct {
	Attributes Attributes `json:"attributes"`
}

type Attributes struct {
	VCPU            string      `json:"vcpu"`
	Memory          string      `json:"memory"`
	InstanceType    string      `json:"instanceType"`
	RegionCode      string      `json:"regionCode"`
	Tenancy         TenancyType `json:"tenancy"`
	UsageType       string      `json:"usagetype"`
	InstanceSKU     string      `json:"instancesku"`
	OperatingSystem string      `json:"operatingSystem"`
	Operation       string      `json:"operation"`
}

type TenancyType string

const (
	Shared TenancyType = "Shared"
)

type Term struct {
	OnDemand map[string]OfferDetail `json:"OnDemand"`
	Reserved map[string]OfferDetail `json:"Reserved"`
}

type OfferDetail map[string]OfferTerm

type OfferTerm struct {
	PriceDimensions map[string]PriceDimension `json:"priceDimensions"`
}

type PriceDimension struct {
	Description  string       `json:"description"`
	Unit         string       `json:"unit"`
	PricePerUnit PricePerUnit `json:"pricePerUnit"`
}

type PricePerUnit struct {
	USD float32 `json:"USD"`
	CNY float32 `json:"CNY"`
}

type GeneralPriceData []EC2GeneralPrice

type EC2GeneralPrice struct {
	InstanceType string       `json:"instanceType"`
	Region       string       `json:"region"`
	PriceModel   string       `json:"priceModel"`
	VCPU         float64      `json:"vcpu"`
	Memory       float64      `json:"memory"`
	PriceUnit    string       `json:"unit"`
	PricePerUnit PricePerUnit `json:"pricePerUnit"`
}

const (
	PriceModelOnDemand string = "onDemand"
	PriceModelReserved string = "reserved"
)
