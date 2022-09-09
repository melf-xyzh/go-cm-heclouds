/**
 * @Time    :2022/9/7 16:20
 * @Author  :Xiaoyu.Zhang
 */

package onenetmod

// GetDataPointsRes 批量查询设备最新数据响应参数
type GetDataPointsRes struct {
	CommonRes
	Data struct { // 接口调用成功之后返回的设备相关信息
		Devices []struct { // 设备信息的json数组
			Id          string `json:"id"`    // 设备ID
			Title       string `json:"title"` // 设备名
			Datastreams []struct { // 设备数据流信息的json数组
				Id    string      `json:"id"`    // 数据流Id
				At    string      `json:"at"`    // 数据点最新时间，北京时间
				Value interface{} `json:"value"` // 数据点的值
			} `json:"datastreams"`
		} `json:"devices"`
	} `json:"data"`
}

// CreateDataPointsReq 新增数据点请求参数
type CreateDataPointsReq struct {
	Datastreams []DataStreams `json:"datastreams"`
}

// CreateDataPointsRes 新增数据点响应参数
type CreateDataPointsRes struct {
	CommonRes
}

type DataStreams struct {
	Id         string      `json:"id,omitempty"`
	Datapoints []DataPoint `json:"datapoints,omitempty"`
}

type DataPoint struct {
	Value interface{} `json:"value"`        // 数据的值
	At    string      `json:"at,omitempty"` // 上传数据点时间
}
