package main

import (
	bmc "github.com/zenlayer/zenlayercloud-sdk-go/zenlayercloud/bmc20221120"
	"github.com/zenlayer/zenlayercloud-sdk-go/zenlayercloud/common"
)

var t = "0"

func main() {
	config := common.NewConfig()
	config.Domain = "10.89.0.11:32070"
	config.Scheme = "HTTP"
	config.Debug = common.Bool(true)

	client, _ := bmc.NewClient(config, "DSxK5RRbouf1jQ2k", "n16ZznQw8Z3XFJcN00PDV7E6gFd4WG")
	request := bmc.NewDescribeInstancesRequest()
	request.PageSize = 2
	request.PageNum = 100
	//request.InstanceIds = []string{"812649911051886092"}

	_, _ = client.DescribeInstances(request)
	print(t)
}
