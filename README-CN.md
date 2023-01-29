快速使用[(English)](./README.md)

---

# 介绍

欢迎使用Zenlayer Cloud 开发者工具SDK，SDK是云 Zenlayer API 平台v2版本的配套工具。目前已经支持bmc等产品，后续所有的云服务产品都会接入进来。 为方便 GO 开发者调试和接入 Zenlayer Cloud
产品 API，提供了一些使用SDK的简单示例。让您快速上手调试 GO SDK。

# 依赖环境

1. Go 环境版本必须不低于 1.9 版本以上
2. 从使用 Zenlayer Cloud SDK Go, 您需要在云平台拥有一个云账号，并在 Zenlayer 云平台控制台中的创建和查看您的 Access Key ID 和 Access Key
   Password。如何获取详见 [帮助文档](https://docs.console.zenlayer.com/welcome/platform/team-management/generate-an-api-access-key)

# 安装

# 快速使用

以BMC服务产品下查询实例接口CreateInstances为例：

```go
package main

func main() {

}
```

### 创建客户端

```go
package main

import bmc "github.com/zenlayer/zenlayercloud-sdk-go/zenlayercloud/bmc20221120"

func main() {

	client, err := bmc.NewClientWithSecretKey("<ACCESS_KEY_ID>", "<ACCESS_KEY_PASSWORD>")
	if err != nil {
		// Handle exceptions
		panic(err)
	}
}
```

## 调试

你可以设置环境变量 `DEBUG=on`开启调试模式，调试模式会打印更详细的日志，当您需要进行详细的排查错误时可以开启。默认调试模式为关闭。 你也可以设置配置 config.Debug = Bool(true)  来进行开启，如下所示：

```go
config := common.NewConfig()
cofig.Debug = Bool(true)
```

## 请求重试

当调用发生网络错误时，可以配置重试发起API的重试，默认重试为关闭的。  
您可以全局开启重试，也可以只对某一个接口请求设置重试

```go
package main

import (
	bmc "github.com/zenlayer/zenlayercloud-sdk-go/zenlayercloud/bmc20221120"
	"github.com/zenlayer/zenlayercloud-sdk-go/zenlayercloud/common"
	"os"
)

func main() {
	// 全局开启重试，并设置重试次数为3次
	config := common.NewConfig()
	config.AutoRetry = true
	config.MaxRetryTime = 3

	// 配置接口重试
	client, _ := bmc.NewClient(config, os.Getenv("ZENLAYERCLOUD_SECRET_KEY_ID"), os.Getenv("ZENLAYERCLOUD_SECRET_KEY_PASSWORD"))
	request := bmc.NewDescribeInstanceTypesRequest()
	request.SetAutoRetries(true) // 如果不设置true，则依旧使用全局的配置
	request.SetMaxAttempts(2)    // 覆盖全局的配置3次
	_, _ = client.DescribeInstanceTypes(request)
}

```

---
快速使用[(English)](./README.md)
