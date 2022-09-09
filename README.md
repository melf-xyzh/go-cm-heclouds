# go-cm-heclouds
中国移动OneNET物联网平台Go - SDK

### 安装

```bash
go get github.com/melf-xyzh/go-cm-heclouds
```

### 示例

初始化 `OneNET Client`

```go
package main

import (
	"github.com/melf-xyzh/go-cm-heclouds/onenet"
	"time"
)

const (
	MasterAPIkey = "d=5jC3w********SL7NOiPOSFMM="
	AccessKey    = "gW6br2S********79MNuVVE0IkIIjth9+kbCrtgNhr0="
	ProductId    = "5***17"
	TimeOut      = time.Second * 5
)

func main() {
	// 创建一个OneNET请求器
	client := onenet.NewOneNetRequestClient(MasterAPIkey, AccessKey, ProductId)
	// 设置请求超时时间
	client.SetTimeOut(TimeOut)
}
```

### 功能

- [ ] 设备管理
  - [x] 新增设备（CreateDevice）
  - [ ] 注册设备
  - [x] 更新设备信息（UpdateDevice）
  - [x] 查询设备详情（GetDeviceInfo）
  - [x] 批量查询设备信息（GetDevicesInfo）
  - [x] 批量查询设备状态（GetDevicesStatus）
  - [x] 删除设备（DeleteDevice）
- [x] 数据流管理
  - [x] 新增数据流（CreateDataStream）
  - [x] 更新数据流（UpdateDataStream）
  - [x] 查询数据流（GetDataStream）
  - [x] 批量查询数据流（GetDataStreams）
  - [x] 删除数据流（DeleteDataStream）
- [ ] 数据点
  - [ ] 查询设备历史数据
  - [x] 批量查询设备数据（GetDevicesDataPoint）
  - [x] 上传数据点（CreateDataPoint）
  - [ ] 上传文件
  - [ ] 获取文件
- [x] 命令
  - [x] 发送命令（SendCmd）
  - [x] 查询命令状态（GetCmdStatus）
  - [x] 查询命令响应（GetCmdResponse）
  - [x] 查询设备历史命令（GetDeviceHistoryCmd）
- [ ] 触发器
  - [ ] 新增触发器
  - [ ] 更新触发器
  - [ ] 查询触发器
  - [ ] 删除触发器
- [ ] APIKey
  - [ ] 新增 APIKey
  - [ ] 更新 APIKey
  - [ ] 查询 APIKey
  - [ ] 删除 APIKey
- [ ] 发布订阅
  - [ ] 发布消息
  - [ ] 查询订阅 topic 设备
  - [ ] 查询设备订阅 topic
  - [ ] 查询产品 topic

### 参考文档

> 旧版文档：https://open.iot.10086.cn/doc/multiprotocol/
>
> 新版文档：https://open.iot.10086.cn/doc/v5/develop/detail/multiprotocol
