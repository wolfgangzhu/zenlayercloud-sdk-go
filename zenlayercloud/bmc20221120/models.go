package bmc

import "github.com/zenlayer/zenlayercloud-sdk-go/zenlayercloud/common"

type CreateInstancesRequest struct {
	*common.BaseRequest
	ZoneId                  string         `json:"zoneId,omitempty"`
	InstanceChargeType      string         `json:"instanceChargeType,omitempty"`
	InstanceChargePrepaid   *ChargePrepaid `json:"instanceChargePrepaid,omitempty"`
	InstanceTypeId          string         `json:"instanceTypeId,omitempty"`
	ImageId                 string         `json:"imageId,omitempty"`
	ResourceGroupId         string         `json:"resourceGroupId,omitempty"`
	InstanceName            string         `json:"instanceName,omitempty"`
	Hostname                string         `json:"hostname,omitempty"`
	Amount                  int            `json:"amount,omitempty"`
	Password                string         `json:"password,omitempty"`
	SshKeys                 []string       `json:"sshKeys,omitempty"`
	InternetChargeType      string         `json:"internetChargeType,omitempty"`
	InternetMaxBandwidthOut int            `json:"internetMaxBandwidthOut,omitempty"`
	TrafficPackageSize      float64        `json:"trafficPackageSize,omitempty"`
	SubnetId                string         `json:"subnetId,omitempty"`
	RaidConfig              *RaidConfig    `json:"raidConfig,omitempty"`
	Partitions              []*Partition   `json:"partitions,omitempty"`
	Nic                     *Nic           `json:"nic,omitempty"`
}

type RaidConfig struct {
	RaidType    *int          `json:"raidType,omitempty"`
	CustomRaids []*CustomRaid `json:"customRaids,omitempty"`
}

type CustomRaid struct {
	RaidType     *int  `json:"raidType"`
	DiskSequence []int `json:"diskSequence,omitempty"`
}

type Partition struct {
	FsPath string `json:"fsPath,omitempty"`
	FsType string `json:"fsType,omitempty"`
	Size   int    `json:"size,omitempty"`
}

type Nic struct {
	WanName string `json:"wanName,omitempty"`
	LanName string `json:"lanName,omitempty"`
}

type CreateInstancesResponse struct {
	*common.BaseResponse
	RequestId string                        `json:"requestId"`
	Response  *CreateInstanceResponseParams `json:"response"`
}

type CreateInstanceResponseParams struct {
	RequestId     *string   `json:"requestId,omitempty"`
	InstanceIdSet []*string `json:"instanceIdSet,omitempty"`
	OrderNumber   *string   `json:"orderNumber,omitempty"`
}

type ChargePrepaid struct {
	Period int `json:"period,omitempty"`
}

type DescribeInstanceTypesRequest struct {
	*common.BaseRequest
	ImageId             string   `json:"imageId,omitempty"`
	InstanceTypeIds     []string `json:"instanceTypeIds,omitempty"`
	MinimumCpuCoreCount *int     `json:"minimumCpuCoreCount,omitempty"`
	MaximumCpuCoreCount *int     `json:"maximumCpuCoreCount,omitempty"`
	MinimumMemorySize   *int     `json:"minimumMemorySize,omitempty"`
	MaximumMemorySize   *int     `json:"maximumMemorySize,omitempty"`
	MinimumBandwidth    *int     `json:"minimumBandwidth,omitempty"`
	SupportRaids        []int    `json:"supportRaids,omitempty"`
	SupportSubnet       *bool    `json:"supportSubnet,omitempty"`
	MinimumDiskSize     *int     `json:"minimumDiskSize,omitempty"`
	MaximumDiskSize     *int     `json:"maximumDiskSize,omitempty"`
	IsHA                *bool    `json:"isHA,omitempty"`
}

type DescribeInstanceTypesResponse struct {
	*common.BaseResponse
	RequestId string                               `json:"requestId"`
	Response  *DescribeInstanceTypesResponseParams `json:"response"`
}

type DescribeInstanceTypesResponseParams struct {
	RequestId     string          `json:"requestId,omitempty"`
	InstanceTypes []*InstanceType `json:"instanceTypes,omitempty"`
}

type InstanceType struct {
	ImageIds         []string          `json:"imageIds,omitempty"`
	InstanceTypeId   string            `json:"instanceTypeId,omitempty"`
	Description      string            `json:"description,omitempty"`
	CpuCoreCount     int               `json:"cpuCoreCount,omitempty"`
	MemorySize       int               `json:"memorySize,omitempty"`
	MaximumBandwidth int               `json:"maximumBandwidth,omitempty"`
	SupportRaids     []int             `json:"supportRaids,omitempty"`
	SupportSubnet    bool              `json:"supportSubnet,omitempty"`
	IsHA             bool              `json:"isHA,omitempty"`
	DiskInfo         *InstanceDiskInfo `json:"diskInfo,omitempty"`
}

type InstanceDiskInfo struct {
	TotalDiskSize   int     `json:"totalDiskSize,omitempty"`
	DiskDescription string  `json:"diskDescription,omitempty"`
	Disks           []*Disk `json:"disks,omitempty"`
}

type Disk struct {
	DiskSize  int `json:"diskSize,omitempty"`
	DiskCount int `json:"diskCount,omitempty"`
}

type DescribeAvailableResourcesRequest struct {
	*common.BaseRequest
	InstanceTypeId     string `json:"instanceTypeId,omitempty"`
	InstanceChargeType string `json:"instanceChargeType,omitempty"`
	ZoneId             string `json:"zoneId,omitempty"`
	SellStatus         string `json:"sellStatus,omitempty"`
}

type DescribeAvailableResourcesResponse struct {
	*common.BaseResponse
	RequestId string                                    `json:"requestId"`
	Response  *DescribeAvailableResourcesResponseParams `json:"response"`
}

type DescribeAvailableResourcesResponseParams struct {
	RequestId          string `json:"requestId"`
	AvailableResources []struct {
		ZoneId                    string   `json:"zoneId,omitempty"`
		SellStatus                string   `json:"sellStatus,omitempty"`
		InternetChargeTypes       []string `json:"internetChargeTypes,omitempty"`
		InstanceTypeId            string   `json:"instanceTypeId,omitempty"`
		MaximumBandwidthOut       int      `json:"maximumBandwidthOut,omitempty"`
		DefaultBandwidthOut       int      `json:"defaultBandwidthOut,omitempty"`
		DefaultTrafficPackageSize float64  `json:"defaultTrafficPackageSize,omitempty"`
	} `json:"availableResources"`
}

type DescribeImagesRequest struct {
	*common.BaseRequest
	ImageIds       []string `json:"imageIds,omitempty"`
	ImageName      string   `json:"imageName,omitempty"`
	Catalog        string   `json:"catalog,omitempty"`
	ImageType      string   `json:"imageType,omitempty"`
	OsType         string   `json:"osType,omitempty"`
	InstanceTypeId string   `json:"instanceTypeId,omitempty"`
}

type DescribeImagesResponse struct {
	*common.BaseResponse
	RequestId string                        `json:"requestId"`
	Response  *DescribeImagesResponseParams `json:"response"`
}

type DescribeImagesResponseParams struct {
	RequestId  string       `json:"requestId"`
	TotalCount int          `json:"totalCount"`
	DataSet    []*ImageInfo `json:"dataSet"`
}

type ImageInfo struct {
	ImageId   string `json:"imageId"`
	ImageName string `json:"imageName"`
	Catalog   string `json:"catalog"`
	ImageType string `json:"imageType"`
	OsType    string `json:"osType"`
}

type DescribeInstancesRequest struct {
	*common.BaseRequest
	InstanceIds        []string `json:"instanceIds,omitempty"`
	ZoneId             string   `json:"zoneId,omitempty"`
	ResourceGroupId    string   `json:"resourceGroupId,omitempty"`
	InstanceTypeId     string   `json:"instanceTypeId,omitempty"`
	InternetChargeType string   `json:"internetChargeType,omitempty"`
	ImageId            string   `json:"imageId,omitempty"`
	SubnetId           string   `json:"subnetId,omitempty"`
	InstanceStatus     string   `json:"instanceStatus,omitempty"`
	InstanceName       string   `json:"instanceName,omitempty"`
	Hostname           string   `json:"hostname,omitempty"`
	PublicIpAddresses  []string `json:"publicIpAddresses,omitempty"`
	PrivateIpAddresses []string `json:"privateIpAddresses,omitempty"`
	PageNum            int      `json:"pageNum,omitempty"`
	PageSize           int      `json:"pageSize,omitempty"`
}

type InstanceInfo struct {
	InstanceId         string       `json:"instanceId,omitempty"`
	ZoneId             string       `json:"zoneId,omitempty"`
	InstanceName       string       `json:"instanceName,omitempty"`
	Hostname           string       `json:"hostname,omitempty"`
	InstanceTypeId     string       `json:"instanceTypeId,omitempty"`
	ImageId            *string      `json:"imageId,omitempty"`
	ImageName          string       `json:"imageName,omitempty"`
	InstanceChargeType string       `json:"instanceChargeType,omitempty"`
	BandwidthOutMbps   *int         `json:"bandwidthOutMbps,omitempty"`
	TrafficPackageSize *float64  	`json:"trafficPackageSize,omitempty"`
	InternetChargeType string       `json:"internetChargeType,omitempty"`
	Period             *int         `json:"period,omitempty"`
	PublicIpAddresses  []string     `json:"publicIpAddresses,omitempty"`
	PrivateIpAddresses []string     `json:"privateIpAddresses,omitempty"`
	Ipv6Addresses      []string     `json:"ipv6Addresses,omitempty"`
	SubnetIds          []string     `json:"subnetIds,omitempty"`
	CreateTime         string       `json:"createTime,omitempty"`
	ExpiredTime        *string      `json:"expiredTime,omitempty"`
	ResourceGroupId    string       `json:"resourceGroupId,omitempty"`
	ResourceGroupName  string       `json:"resourceGroupName,omitempty"`
	InstanceStatus     string       `json:"instanceStatus,omitempty"`
	Partitions         []*Partition `json:"partitions,omitempty"`
	RaidConfig         *RaidConfig  `json:"raidConfig,omitempty"`
	Nic                *Nic         `json:"nic,omitempty"`
}

type DescribeInstancesResponseParams struct {
	RequestId  string          `json:"requestId"`
	TotalCount int             `json:"totalCount"`
	DataSet    []*InstanceInfo `json:"dataSet"`
}

type DescribeInstancesResponse struct {
	*common.BaseResponse
	RequestId string                           `json:"requestId"`
	Response  *DescribeInstancesResponseParams `json:"response"`
}

type StartInstancesRequest struct {
	*common.BaseRequest
	InstanceIds []string `json:"instanceIds,omitempty"`
}

type StartInstancesResponse struct {
	*common.BaseResponse
	RequestId string `json:"requestId"`
	Response  struct {
		RequestId string `json:"requestId"`
	} `json:"response"`
}

type StopInstancesRequest struct {
	*common.BaseRequest
	InstanceIds []string `json:"instanceIds,omitempty"`
}

type StopInstancesResponse struct {
	*common.BaseResponse
	RequestId string `json:"requestId"`
	Response  struct {
		RequestId string `json:"requestId"`
	} `json:"response"`
}

type RebootInstancesRequest struct {
	*common.BaseRequest
	InstanceIds []string `json:"instanceIds,omitempty"`
}

type RebootInstancesResponse struct {
	*common.BaseResponse
	RequestId string `json:"requestId"`
	Response  struct {
		RequestId string `json:"requestId"`
	} `json:"response"`
}

type ReinstallInstanceRequest struct {
	*common.BaseRequest
	InstanceId string       `json:"instanceId,omitempty"`
	ImageId    string       `json:"imageId,omitempty"`
	Hostname   string       `json:"hostname,omitempty"`
	Password   string       `json:"password,omitempty"`
	SshKeys    []string     `json:"sshKeys,omitempty"`
	RaidConfig *RaidConfig  `json:"raidConfig,omitempty"`
	Partitions []*Partition `json:"partitions,omitempty"`
	Nic        *Nic         `json:"nic,omitempty"`
}

type ReInstallInstanceResponse struct {
	*common.BaseResponse
	RequestId string `json:"requestId"`
	Response  struct {
		RequestId string `json:"requestId"`
	} `json:"response"`
}

type TerminateInstanceRequest struct {
	*common.BaseRequest
	InstanceId string `json:"instanceIds,omitempty"`
}

type TerminateInstanceResponse struct {
	*common.BaseResponse
	RequestId string `json:"requestId"`
	Response  struct {
		RequestId string `json:"requestId"`
	} `json:"response"`
}

type ReleaseInstancesRequest struct {
	*common.BaseRequest
	InstanceIds []string `json:"instanceIds,omitempty"`
}

type ReleaseInstancesResponse struct {
	*common.BaseResponse
	RequestId string `json:"requestId"`
	Response  struct {
		RequestId string `json:"requestId"`
	} `json:"response"`
}

type RenewInstanceRequest struct {
	*common.BaseRequest
	InstanceId string `json:"instanceIds,omitempty"`
}

type RenewInstanceResponse struct {
	*common.BaseResponse
	RequestId string `json:"requestId"`
	Response  struct {
		RequestId string `json:"requestId"`
	} `json:"response"`
}

type ModifyInstancesAttributeRequest struct {
	*common.BaseRequest
	InstanceIds  []string `json:"instanceIds,omitempty"`
	InstanceName string   `json:"instanceName,omitempty"`
}

type ModifyInstanceAttributeResponse struct {
	*common.BaseResponse
	RequestId string `json:"requestId"`
	Response  struct {
		RequestId string `json:"requestId"`
	} `json:"response"`
}

type InquiryPriceCreateInstanceRequest struct {
	*common.BaseRequest
	ZoneId                  string         `json:"zoneId,omitempty"`
	InstanceTypeId          string         `json:"instanceTypeId,omitempty"`
	InstanceChargeType      string         `json:"instanceChargeType,omitempty"`
	InstanceChargePrepaid   *ChargePrepaid `json:"instanceChargePrepaid,omitempty"`
	TrafficPackageSize      float64        `json:"trafficPackageSize,omitempty"`
	InternetMaxBandwidthOut int            `json:"internetMaxBandwidthOut,omitempty"`
	InternetChargeType      string         `json:"internetChargeType,omitempty"`
}

type InquiryPriceCreateInstanceResponseParams struct {
	RequestId      string   `json:"requestId"`
	InstancePrice  *Price   `json:"instancePrice"`
	BandwidthPrice []*Price `json:"bandwidthPrice"`
}

type InquiryPriceCreateInstanceResponse struct {
	*common.BaseResponse
	RequestId string                                    `json:"requestId"`
	Response  *InquiryPriceCreateInstanceResponseParams `json:"response"`
}

type Price struct {
	Discount          float64      `json:"discount"`
	DiscountPrice     float64      `json:"discountPrice"`
	OriginalPrice     float64      `json:"originalPrice"`
	UnitPrice         float64      `json:"unitPrice"`
	DiscountUnitPrice float64      `json:"discountUnitPrice"`
	ChargeUnit        string       `json:"chargeUnit"`
	StepPrices        []*StepPrice `json:"stepPrices"`
}

type StepPrice struct {
	StepStart         float64 `json:"stepStart"`
	StepEnd           float64 `json:"stepEnd"`
	UnitPrice         float64 `json:"unitPrice"`
	DiscountUnitPrice float64 `json:"discountUnitPrice"`
}

type ModifyInstanceBandwidthRequest struct {
	*common.BaseRequest
	InstanceId       string `json:"instanceId,omitempty"`
	BandwidthOutMbps *int   `json:"bandwidthOutMbps,omitempty"`
}

type ModifyInstanceBandwidthResponseParams struct {
	RequestId   string `json:"requestId"`
	OrderNumber string `json:"orderNumber,omitempty"`
}

type ModifyInstanceBandwidthResponse struct {
	*common.BaseResponse
	RequestId string                                 `json:"requestId"`
	Response  *ModifyInstanceBandwidthResponseParams `json:"response"`
}
