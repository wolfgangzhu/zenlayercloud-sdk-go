package bmc

import (
	"github.com/zenlayer/zenlayercloud-sdk-go/zenlayercloud/common"
)

const (
	APIVersion = "2022-11-20"
	SERVICE    = "bmc"
)

type Client struct {
	common.Client
}

func NewClientWithSecretKey(secretKeyId, secretKeyPassword string) (client *Client, err error) {
	return NewClient(common.NewConfig(), secretKeyId, secretKeyPassword)
}

func NewClient(config *common.Config, secretKeyId, secretKeyPassword string) (client *Client, err error) {
	client = &Client{}

	err = client.InitWithCredential(common.NewCredential(secretKeyId, secretKeyPassword))
	if err != nil {
		return nil, err
	}
	err = client.WithConfig(config)

	if err != nil {
		return nil, err
	}
	return client, nil
}

func NewCreateInstancesRequest() (request *CreateInstancesRequest) {
	request = &CreateInstancesRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "CreateInstances")

	return
}

func NewCreateInstancesResponse() (response *CreateInstancesResponse) {
	response = &CreateInstancesResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

func (c *Client) CreateInstances(request *CreateInstancesRequest) (response *CreateInstancesResponse, err error) {
	response = NewCreateInstancesResponse()
	err = c.ApiCall(request, response)
	return
}

func NewDescribeInstanceTypesRequest() (request *DescribeInstanceTypesRequest) {
	request = &DescribeInstanceTypesRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "DescribeInstanceTypes")

	return
}

func NewDescribeInstanceTypesResponse() (response *DescribeInstanceTypesResponse) {
	response = &DescribeInstanceTypesResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

func (c *Client) DescribeInstanceTypes(request *DescribeInstanceTypesRequest) (response *DescribeInstanceTypesResponse, err error) {
	response = NewDescribeInstanceTypesResponse()
	err = c.ApiCall(request, response)
	return
}

func NewDescribeAvailableResourcesRequest() (request *DescribeAvailableResourcesRequest) {
	request = &DescribeAvailableResourcesRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "DescribeAvailableResources")

	return
}

func NewDescribeAvailableResourcesResponse() (response *DescribeAvailableResourcesResponse) {
	response = &DescribeAvailableResourcesResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

func (c *Client) DescribeAvailableResources(request *DescribeAvailableResourcesRequest) (response *DescribeAvailableResourcesResponse, err error) {
	response = NewDescribeAvailableResourcesResponse()
	err = c.ApiCall(request, response)
	return
}

func NewDescribeImagesRequest() (request *DescribeImagesRequest) {
	request = &DescribeImagesRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "DescribeImages")

	return
}

func NewDescribeImagesResponse() (response *DescribeImagesResponse) {
	response = &DescribeImagesResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

func (c *Client) DescribeImages(request *DescribeImagesRequest) (response *DescribeImagesResponse, err error) {
	response = NewDescribeImagesResponse()
	err = c.ApiCall(request, response)
	return
}

func NewDescribeInstancesRequest() (request *DescribeInstancesRequest) {
	request = &DescribeInstancesRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "DescribeInstances")

	return
}

func NewDescribeInstancesResponse() (response *DescribeInstancesResponse) {
	response = &DescribeInstancesResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

func (c *Client) DescribeInstances(request *DescribeInstancesRequest) (response *DescribeInstancesResponse, err error) {
	response = NewDescribeInstancesResponse()
	err = c.ApiCall(request, response)
	return
}

func NewStartInstancesRequest() (request *StartInstancesRequest) {
	request = &StartInstancesRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "StartInstances")

	return
}

func NewStartInstancesResponse() (response *StartInstancesResponse) {
	response = &StartInstancesResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

func (c *Client) StartInstances(request *StartInstancesRequest) (response *StartInstancesResponse, err error) {
	response = NewStartInstancesResponse()
	err = c.ApiCall(request, response)
	return
}

func NewStopInstancesRequest() (request *StopInstancesRequest) {
	request = &StopInstancesRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "StopInstances")

	return
}

func NewStopInstancesResponse() (response *StopInstancesResponse) {
	response = &StopInstancesResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

func (c *Client) StopInstances(request *StopInstancesRequest) (response *StopInstancesResponse, err error) {
	response = NewStopInstancesResponse()
	err = c.ApiCall(request, response)
	return
}

func NewRebootInstancesRequest() (request *RebootInstancesRequest) {
	request = &RebootInstancesRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "RebootInstances")

	return
}

func NewRebootInstancesResponse() (response *RebootInstancesResponse) {
	response = &RebootInstancesResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

func (c *Client) RebootInstances(request *RebootInstancesRequest) (response *RebootInstancesResponse, err error) {
	response = NewRebootInstancesResponse()
	err = c.ApiCall(request, response)
	return
}

func NewReinstallInstancesRequest() (request *ReinstallInstanceRequest) {
	request = &ReinstallInstanceRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "ReinstallInstance")

	return
}

func NewReinstallInstanceResponse() (response *ReInstallInstanceResponse) {
	response = &ReInstallInstanceResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

func (c *Client) ReinstallInstance(request *ReinstallInstanceRequest) (response *ReInstallInstanceResponse, err error) {
	response = NewReinstallInstanceResponse()
	err = c.ApiCall(request, response)
	return
}

func NewTerminateInstancesRequest() (request *TerminateInstanceRequest) {
	request = &TerminateInstanceRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "TerminateInstance")

	return
}

func NewTerminateInstanceResponse() (response *TerminateInstanceResponse) {
	response = &TerminateInstanceResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

func (c *Client) TerminateInstance(request *TerminateInstanceRequest) (response *TerminateInstanceResponse, err error) {
	response = NewTerminateInstanceResponse()
	err = c.ApiCall(request, response)
	return
}

func NewReleaseInstancesRequest() (request *ReleaseInstancesRequest) {
	request = &ReleaseInstancesRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "ReleaseInstances")

	return
}

func NewReleaseInstancesResponse() (response *ReleaseInstancesResponse) {
	response = &ReleaseInstancesResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

func (c *Client) ReleaseInstances(request *ReleaseInstancesRequest) (response *ReleaseInstancesResponse, err error) {
	response = NewReleaseInstancesResponse()
	err = c.ApiCall(request, response)
	return
}

func NewRenewInstanceRequest() (request *RenewInstanceRequest) {
	request = &RenewInstanceRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "RenewInstance")

	return
}

func NewRenewInstanceResponse() (response *RenewInstanceResponse) {
	response = &RenewInstanceResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

func (c *Client) RenewInstance(request *RenewInstanceRequest) (response *RenewInstanceResponse, err error) {
	response = NewRenewInstanceResponse()
	err = c.ApiCall(request, response)
	return
}

func NewModifyInstancesAttributeRequest() (request *ModifyInstancesAttributeRequest) {
	request = &ModifyInstancesAttributeRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "ModifyInstanceAttribute")

	return
}

func NewModifyInstanceAttributeResponse() (response *ModifyInstanceAttributeResponse) {
	response = &ModifyInstanceAttributeResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

func (c *Client) ModifyInstancesAttribute(request *ModifyInstancesAttributeRequest) (response *ModifyInstanceAttributeResponse, err error) {
	response = NewModifyInstanceAttributeResponse()
	err = c.ApiCall(request, response)
	return
}

func NewInquiryPriceCreateInstanceRequest() (request *InquiryPriceCreateInstanceRequest) {
	request = &InquiryPriceCreateInstanceRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "InquiryPriceCreateInstance")

	return
}

func NewInquiryPriceCreateInstanceRequestResponse() (response *InquiryPriceCreateInstanceResponse) {
	response = &InquiryPriceCreateInstanceResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

func (c *Client) InquiryPriceCreateInstance(request *InquiryPriceCreateInstanceRequest) (response *InquiryPriceCreateInstanceResponse, err error) {
	response = NewInquiryPriceCreateInstanceRequestResponse()
	err = c.ApiCall(request, response)
	return
}

func NewModifyInstanceBandwidthRequest() (request *ModifyInstanceBandwidthRequest) {
	request = &ModifyInstanceBandwidthRequest{
		BaseRequest: &common.BaseRequest{},
	}
	request.Init().InitWithApiInfo(SERVICE, APIVersion, "ModifyInstanceBandwidth")

	return
}

func NewModifyInstanceBandwidthResponse() (response *ModifyInstanceBandwidthResponse) {
	response = &ModifyInstanceBandwidthResponse{
		BaseResponse: &common.BaseResponse{},
	}
	return
}

func (c *Client) ModifyInstanceBandwidth(request *ModifyInstanceBandwidthRequest) (response *ModifyInstanceBandwidthResponse, err error) {
	response = NewModifyInstanceBandwidthResponse()
	err = c.ApiCall(request, response)
	return
}
