package main

import (
	"encoding/json"
	"fmt"
	bmc "github.com/zenlayer/zenlayercloud-sdk-go/zenlayercloud/bmc20221120"
	"os"
)

func main1() {

	client, _ := bmc.NewClientWithSecretKey(os.Getenv("ZENLAYERCLOUD_SECRET_KEY_ID"), os.Getenv("ZENLAYERCLOUD_SECRET_KEY_PASSWORD"))
	request := bmc.NewCreateInstancesRequest()
	request.InstanceTypeId = "M6C"
	request.InstanceChargeType = "POSTPAID"
	request.InternetChargeType = "ByBandwidth"
	request.ZoneId = "SEL-A"

	response, err := client.CreateInstances(request)
	if err != nil {
		panic(err)
	}
	b, _ := json.Marshal(response.Response)
	fmt.Printf("%s", b)
}
