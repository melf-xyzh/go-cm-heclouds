/**
 * @Time    :2022/9/3 19:55
 * @Author  :MELF晓宇
 * @Email   :xyzh.melf@petalmail.com
 * @FileName:datastreams.go
 * @Project :go-cm-heclouds
 * @Blog    :https://blog.csdn.net/qq_29537269
 * @Guide   :https://guide.melf.space
 * @Information:
 *		数据流相关
 */

package onenet

import (
	"errors"
	"fmt"
	"github.com/melf-xyzh/go-cm-heclouds/constant"
	"github.com/melf-xyzh/go-cm-heclouds/model"
)

// CreateDataStream
/**
 *  @Description: 新增数据流
 *  @receiver c
 *  @param deviceId 设备ID
 *  @param req 请求参数
 *  @return res 响应参数
 *  @return err 异常
 */
func (c *Client) CreateDataStream(deviceId string, req onenetmod.CreateDataStreamReq) (res onenetmod.CreateDataStreamRes, err error) {
	// 序列化Body
	jsonByte, err := json.Marshal(&req)
	if err != nil {
		return res, err
	}
	// 拼接Url
	url := constant.BaseUrl + fmt.Sprintf(constant.CreateDataStreamsUrl, deviceId)
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

// UpdateDataStream
/**
 *  @Description: 更新数据流属性
 *  @receiver c
 *  @param deviceId 设备ID
 *  @param dataStreamId 数据流ID
 *  @param req 请求参数
 *  @return res 响应参数
 *  @return err 错误
 */
func (c *Client) UpdateDataStream(deviceId, dataStreamId string, req onenetmod.UpdateDataStreamReq) (res onenetmod.UpdateDataStreamRes, err error) {
	// 序列化Body
	jsonByte, err := json.Marshal(&req)
	if err != nil {
		return res, err
	}
	// 拼接Url
	url := constant.BaseUrl + fmt.Sprintf(constant.UpdateDataStreamsUrl, deviceId, dataStreamId)
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

// GetDataStream
/**
 *  @Description:查询数据流详情
 *  @receiver c
 *  @param deviceId 设备ID
 *  @param dataStreamId 数据流ID
 *  @return res 响应参数
 *  @return err 错误
 */
func (c *Client) GetDataStream(deviceId, dataStreamId string) (res onenetmod.GetDataStreamRes, err error) {
	// 拼接Url
	url := constant.BaseUrl + fmt.Sprintf(constant.GetDataStreamUrl, deviceId, dataStreamId)
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

// GetDataStreams
/**
 *  @Description: 批量查询数据流信息
 *  @receiver c
 *  @param deviceId 设备ID
 *  @return res 响应参数
 *  @return err 错误
 */
func (c *Client) GetDataStreams(deviceId string) (res onenetmod.GetDataStreamsRes, err error) {
	// 拼接Url
	url := constant.BaseUrl + fmt.Sprintf(constant.GetDataStreamsUrl, deviceId)
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

// DeleteDataStream
/**
 *  @Description: 删除数据流
 *  @receiver c
 *  @param deviceId 设备ID
 *  @param dataStreamId 数据流ID
 *  @return res 响应参数
 *  @return err 错误
 */
func (c *Client) DeleteDataStream(deviceId, dataStreamId string) (res onenetmod.DeleteDataStreamRes, err error) {
	// 拼接Url
	url := constant.BaseUrl + fmt.Sprintf(constant.DeleteDataStreamUrl, deviceId, dataStreamId)
	// 发起请求
	var body []byte
	body, err = c.RequestClient(url, "DELETE",
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
