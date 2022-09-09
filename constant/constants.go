/**
 * @Time    :2022/9/4 16:39
 * @Author  :MELF晓宇
 * @Email   :xyzh.melf@petalmail.com
 * @FileName:constants.go
 * @Project :go-cm-heclouds
 * @Blog    :https://blog.csdn.net/qq_29537269
 * @Guide   :https://guide.melf.space
 * @Information:
 *
 */

package constant

import "time"

const (
	ErrNoSuccess = 0
)

const (
	DefaultTimeOut = 5 * time.Second
)

const (
	BaseUrl = "http://api.heclouds.com"

	CreateDeviceUrl     = "/devices"        // 新增设备
	UpdateDeviceUrl     = "/devices/%s"     // 更新设备信息
	GetDeviceInfoUrl    = "/devices/%s"     // 查询设备详情
	GetDevicesInfoUrl   = "/devices"        // 批量查询设备信息
	GetDevicesStatusUrl = "/devices/status" // 批量查询设备状态
	DeleteDeviceUrl     = "/devices/%s"     // 删除设备

	CreateDataStreamsUrl = "/devices/%s/datastreams"    // 新增数据流
	UpdateDataStreamsUrl = "/devices/%s/datastreams/%s" // 更新数据流
	GetDataStreamUrl     = "/devices/%s/datastreams/%s" // 查询数据流
	GetDataStreamsUrl    = "/devices/%s/datastreams"    // 批量查询数据流
	DeleteDataStreamUrl  = "/devices/%s/datastreams/%s" // 删除数据流

	CreateDataPointUrl     = "/devices/%s/datapoints" // 新增数据点（上传数据点）
	GetDevicesDataPointUrl = "/devices/datapoints"    // 批量查询设备最新数据
	GetDeviceDataPointsUrl = "/devices/%s/datapoints" // 查询设备历史数据

	SendCmd        = "/cmds"            // 发送命令
	GetCmdStatus   = "/cmds/%s"         // 查询命令状态
	GetCmdResponse = "/cmds/%s/resp"    // 查询命令响应
	GetCmdHistory  = "/cmds/history/%s" // 查询设备历史命令
)

const (
	ContentTypeJson       = "application/json;charset=UTF-8"
	ContentTypeUrlencoded = "application/x-www-form-urlencoded"
	ContentTypeFromData   = "multipart/form-data"
)
