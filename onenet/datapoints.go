/**
 * @Time    :2022/9/3 19:59
 * @Author  :MELF晓宇
 * @Email   :xyzh.melf@petalmail.com
 * @FileName:datapoints.go
 * @Project :go-cm-heclouds
 * @Blog    :https://blog.csdn.net/qq_29537269
 * @Guide   :https://guide.melf.space
 * @Information:
 *
 */

package onenet

import (
	"errors"
	"fmt"
	"github.com/melf-xyzh/go-cm-heclouds/constant"
	onenetmod "github.com/melf-xyzh/go-cm-heclouds/model"
	"strings"
)

// GetDeviceDataPoint
/**
 * @Description: 查询设备历史数据
 * @receiver c
 */
func (c *Client) GetDeviceDataPoint() {

}

// GetDevicesDataPoint
/**
 *  @Description: 批量查询设备最新数据
 *  @receiver c
 *  @param deviceIds
 *  @return res
 *  @return err
 */
func (c *Client) GetDevicesDataPoint(deviceIds []string) (res onenetmod.GetDataPointsRes, err error) {
	// 处理参数
	var build strings.Builder
	for i, deviceId := range deviceIds {
		build.WriteString(deviceId)
		if i != len(deviceIds)-1 {
			build.WriteString(",")
		}
	}
	// 处理Url
	paramMap := make(map[string]string, 1)
	paramMap["devIds"] = build.String()
	var url string
	url, err = c.getUrlPathWithParams(constant.GetDevicesDataPointUrl, paramMap)
	if err != nil {
		return
	}
	url = constant.BaseUrl + url
	// 发起请求
	var body []byte
	body, err = c.RequestClient(url, "GET", ApiKey(c.MasterAPIkey))
	if err != nil {
		return
	}
	// 反序列化ResponseBody
	err = json.Unmarshal(body, &res)
	if err != nil {
		return res, err
	}
	// 抛出非succ响应码造成的错误
	if res.Errno != constant.ErrNoSuccess {
		err = errors.New(res.Error)
		return
	}
	return
}

// CreateDataPoint
/**
 *  @Description: 新增单个数据点
 *  @receiver c
 *  @param deviceId
 *  @param dataStreamName
 *  @param value
 *  @return res
 *  @return err
 */
func (c *Client) CreateDataPoint(deviceId, dataStreamName string, value interface{}) (res onenetmod.CreateDataPointsRes, err error) {
	// 处理Body
	req := onenetmod.CreateDataPointsReq{
		Datastreams: []onenetmod.DataStreams{
			{
				Id: dataStreamName,
				Datapoints: []onenetmod.DataPoint{
					{
						Value: value,
					},
				},
			},
		},
	}
	// 序列化Body
	jsonByte, err := json.Marshal(&req)
	if err != nil {
		return res, err
	}
	// 拼接Url
	url := constant.BaseUrl + fmt.Sprintf(constant.CreateDataPointUrl, deviceId)
	// 发起请求
	var body []byte
	body, err = c.RequestClient(url, "POST",
		ApiKey(c.MasterAPIkey),
		Body(jsonByte),
		ContentType(constant.ContentTypeJson))
	// 抛出错误
	if err != nil {
		return
	}
	// 反序列化ResponseBody
	err = json.Unmarshal(body, &res)
	if err != nil {
		return res, err
	}
	// 抛出非succ响应码造成的错误
	if res.Errno != constant.ErrNoSuccess {
		err = errors.New(res.Error)
		return
	}
	return
}
