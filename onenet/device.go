/**
 * @Time    :2022/9/3 19:51
 * @Author  :MELF晓宇
 * @Email   :xyzh.melf@petalmail.com
 * @FileName:device.go
 * @Project :go-cm-heclouds
 * @Blog    :https://blog.csdn.net/qq_29537269
 * @Guide   :https://guide.melf.space
 * @Information:
 *		设备相关
 */

package onenet

import (
	"errors"
	"fmt"
	"github.com/melf-xyzh/go-cm-heclouds/constant"
	onenetmod "github.com/melf-xyzh/go-cm-heclouds/model"
	"strings"
)

// CreateDevice
/**
 *  @Description: 新增设备
 *  @receiver c
 *  @param req 请求参数
 *  @return res 响应参数
 *  @return err 错误
 */
func (c *Client) CreateDevice(req onenetmod.CreateOrUpdateDeviceReq) (res onenetmod.CreateDeviceRes, err error) {
	// 序列化Body
	jsonByte, err := json.Marshal(&req)
	if err != nil {
		return res, err
	}
	// 拼接Url
	url := constant.BaseUrl + constant.CreateDeviceUrl
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

// UpdateDevice
/**
 *  @Description: 更新设备
 *  @receiver c
 *  @param deviceId 设备ID
 *  @param req 请求参数
 *  @return res 响应参数
 *  @return err 错误
 */
func (c *Client) UpdateDevice(deviceId string, req onenetmod.CreateOrUpdateDeviceReq) (res onenetmod.UpdateDeviceRes, err error) {
	// 序列化Body
	jsonByte, err := json.Marshal(&req)
	if err != nil {
		return res, err
	}
	// 拼接Url
	url := constant.BaseUrl + fmt.Sprintf(constant.UpdateDeviceUrl, deviceId)
	// 发起请求
	var body []byte
	body, err = c.RequestClient(url, "PUT",
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

// GetDeviceInfo
/**
 *  @Description: 查询设备详情
 *  @receiver c
 *  @param deviceId 设备ID
 *  @return res 响应参数
 *  @return err 错误
 */
func (c *Client) GetDeviceInfo(deviceId string) (res onenetmod.DeviceInfoRes, err error) {
	// 拼接Url
	url := constant.BaseUrl + fmt.Sprintf(constant.GetDeviceInfoUrl, deviceId)
	// 发起请求
	var body []byte
	body, err = c.RequestClient(url, "GET", ApiKey(c.MasterAPIkey))
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

// GetDevicesStatus
/**
 *  @Description: 批量查询设备状态
 *  @receiver c
 *  @param deviceIds 指定设备ID，多个用逗号分隔，最多1000个
 *  @return res 响应参数
 *  @return err 错误
 */
func (c *Client) GetDevicesStatus(deviceIds []string) (res onenetmod.GetDevicesStatusRes, err error) {
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
	url, err = c.getUrlPathWithParams(constant.GetDevicesStatusUrl, paramMap)
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

// GetDevicesInfo
/**
 *  @Description: 批量查询设备详情
 *  @receiver c
 *  @param req
 *  @return res
 *  @return err
 */
func (c *Client) GetDevicesInfo(req onenetmod.GetDevicesReq) (res onenetmod.DevicesInfoRes, err error) {
	// 处理DeviceId
	var build strings.Builder
	for i, deviceId := range req.DeviceIds {
		build.WriteString(deviceId)
		if i != len(req.DeviceIds)-1 {
			build.WriteString(",")
		}
	}
	req.DeviceId = build.String()
	req.DeviceIds = nil

	// 处理请求参数
	paramMap := make(map[string]string)
	paramMap["devIds"] = req.DeviceId
	// 处理Url
	var url string
	url, err = c.getUrlPathWithParams(constant.GetDevicesInfoUrl, paramMap)
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

// DeleteDevice
/**
 * @Description: 删除设备
 * @receiver c
 */
func (c *Client) DeleteDevice(deviceId string) (res onenetmod.UpdateDeviceRes, err error) {
	// 拼接Url
	url := constant.BaseUrl + fmt.Sprintf(constant.DeleteDeviceUrl, deviceId)
	// 发起请求
	var body []byte
	body, err = c.RequestClient(url, "POST",
		ApiKey(c.MasterAPIkey),
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

// RegisterDevice
/**
 * @Description: 注册设备
 * @receiver c
 */
func (c *Client) RegisterDevice() {

}
