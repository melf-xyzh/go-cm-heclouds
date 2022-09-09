/**
 * @Time    :2022/9/7 14:17
 * @Author  :Xiaoyu.Zhang
 */

package onenetmod

// CreateDataStreamReq 新增数据流请求参数
type CreateDataStreamReq struct {
	Id         string   `json:"id"`                    // 数据流ID，即数据流名称
	Tags       []string `json:"tags,omitempty"`        // 数据流标签，可为一个或者多个
	Unit       string   `json:"unit,omitempty"`        // 数据单位
	UnitSymbol string   `json:"unit_symbol,omitempty"` // 数据单位符号
}

// CreateDataStreamRes 新增数据流响应参数
type CreateDataStreamRes struct {
	CommonRes
	Data struct {
		DsUuid string `json:"ds_uuid"` // 数据流平台内部唯一ID
	} `json:"data"`
}

// UpdateDataStreamReq 更新数据流请求参数
type UpdateDataStreamReq struct {
	Tags       []string `json:"tags,omitempty"`        // 数据流标签，可为一个或者多个
	Unit       string   `json:"unit,omitempty"`        // 数据单位
	UnitSymbol string   `json:"unit_symbol,omitempty"` // 数据单位符号
}

// UpdateDataStreamRes 更新数据流响应参数
type UpdateDataStreamRes struct {
	CommonRes
}

// GetDataStreamRes 查询数据流响应数据
type GetDataStreamRes struct {
	CommonRes
	Data DataStream `json:"data"` // 接口调用成功之后返回的设备相关信息
}

// DataStream 数据流信息
type DataStream struct {
	Id           string      `json:"id"`            // 数据流ID
	CreateTime   string      `json:"create_time"`   // 数据流创建时间
	UpdateAt     string      `json:"update_at"`     // 最新数据上传时间
	CurrentValue interface{} `json:"current_value"` // 最新数据点
}

// GetDataStreamsRes 批量查询数据流响应数据
type GetDataStreamsRes struct {
	CommonRes
	Data []DataStream `json:"data,omitempty"` // 接口调用成功之后返回的设备相关信息
}

// DeleteDataStreamRes 删除设备响应数据
type DeleteDataStreamRes struct {
	CommonRes
}
