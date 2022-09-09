/**
 * @Time    :2022/9/4 16:43
 * @Author  :MELF晓宇
 * @Email   :xyzh.melf@petalmail.com
 * @FileName:device.go
 * @Project :go-cm-heclouds
 * @Blog    :https://blog.csdn.net/qq_29537269
 * @Guide   :https://guide.melf.space
 * @Information:
 *
 */

package onenetmod

// CreateOrUpdateDeviceReq 新增设备请求参数
type CreateOrUpdateDeviceReq struct {
	Title    string                 `json:"title"`               // 设备名称
	Desc     string                 `json:"desc,omitempty"`      // 设备描述
	Tags     []string               `json:"tags,omitempty"`      // 设备标签，可为一个或者多个，见示例
	Location Location               `json:"location,omitempty"`  // 设备位置坐标信息，以经纬度键值对表示:{"lon":xx,"lat":xx}
	Private  bool                   `json:"private,omitempty"`   // 设备私密性，决定应用编辑器分享链接中设备信息的可见度，默认为true
	AuthInfo string                 `json:"auth_info,omitempty"` // 鉴权信息，建议携带，并设置为设备的产品序列号
	Other    map[string]interface{} `json:"other,omitempty"`     // 其他设备自定义信息，以键值对格式表示，见示例
}

// Location 定位
type Location struct {
	Lon float64 `json:"lon"` // 经度
	Lat float64 `json:"lat"` // 纬度
}

// GetDevicesReq 批量查询设备信息请求参数
type GetDevicesReq struct {
	KeyWords  string `json:"key_words,omitempty"` // 匹配关键字，从id和title字段中左匹配
	AuthInfo  string `json:"auth_info,omitempty"` // 鉴权信息，建议携带，并设置为设备的产品序列号
	Tag       string `json:"tag,omitempty"`       // 设备标签，可为一个或者多个，见示例
	Online    string `json:"online,omitempty"`    // 设备在线状态
	Private   bool   `json:"private,omitempty"`   // 设备私密性，决定应用编辑器分享链接中设备信息的可见度，默认为true
	Page      int    `json:"page,omitempty"`      // 指定页码，最大页数为10000
	PerPage   string `json:"per_page,omitempty"`  // 指定每页输出设备个数，默认30，最多100
	DeviceId  string `json:"device_id,omitempty"` // 指定设备ID，多个用逗号分隔，最多100个
	Begin     string `json:"begin,omitempty"`     // 起始时间，北京时间，示例：2016-06-20
	End       string `json:"end,omitempty"`       // 结束时间，北京时间，示例：2016-06-20
	DeviceIds []string
}

// UpdateDeviceRes 更新设备响应参数
type UpdateDeviceRes struct {
	CommonRes
}

// CreateDeviceRes 新增设备响应参数
type CreateDeviceRes struct {
	CommonRes
	Data struct {
		DeviceId string `json:"device_id"`
	} `json:"data"`
}

// GetDevicesStatusRes 批量查询设备状态接口响应
type GetDevicesStatusRes struct {
	CommonRes
	Data struct {
		Devices []struct {
			Title  string `json:"title"`
			Online bool   `json:"online"`
			Id     string `json:"id"`
		} `json:"devices"`
		TotalCount int `json:"total_count"`
	} `json:"data"`
}

// DeviceInfoRes 设备详情响应参数
type DeviceInfoRes struct {
	CommonRes
	Data struct {
		Protocol string `json:"protocol"`
		Other    struct {
			Version      string `json:"version"`
			Manufacturer string `json:"manufacturer"`
		} `json:"other"`
		CreateTime string `json:"create_time"`
		Online     bool   `json:"online"`
		Location   struct {
			Lat float64 `json:"lat"`
			Lon int     `json:"lon"`
		} `json:"location"`
		Id          string `json:"id"`
		AuthInfo    string `json:"auth_info"`
		Datastreams []struct {
			Unit       string `json:"unit"`
			Id         string `json:"id"`
			UnitSymbol string `json:"unit_symbol"`
			CreateTime string `json:"create_time"`
		} `json:"datastreams"`
		Title string   `json:"title"`
		Desc  string   `json:"desc"`
		Tags  []string `json:"tags"`
	} `json:"data"`
}

type DevicesInfoRes struct {
	CommonRes
	Data struct {
		Devices []struct {
			Protocol string `json:"protocol"`
			Private  bool   `json:"private,omitempty"`
			Other    struct {
				Name string `json:"name"`
				V    int    `json:"v"`
			} `json:"other,omitempty"`
			CreateTime string `json:"create_time"`
			Online     bool   `json:"online"`
			Location   struct {
				Lat float64 `json:"lat"`
				Lon float64 `json:"lon"`
			} `json:"location"`
			Id        string      `json:"id"`
			AuthInfo  interface{} `json:"auth_info"`
			Title     string      `json:"title"`
			Desc      string      `json:"desc"`
			Tags      []string    `json:"tags"`
			ActTime   string      `json:"act_time,omitempty"`
			LastLogin string      `json:"last_login,omitempty"`
		} `json:"devices"`
		PerPage    int `json:"per_page"`
		TotalCount int `json:"total_count"`
		Page       int `json:"page"`
	} `json:"data"`
}
