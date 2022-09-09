/**
 * @Time    :2022/9/3 20:04
 * @Author  :MELF晓宇
 * @Email   :xyzh.melf@petalmail.com
 * @FileName:cmds.go
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
	"github.com/melf-xyzh/go-cm-heclouds/model"
	"strconv"
	"time"
)

// SendCmd
/**
 *  @Description: 发送命令
 *  @receiver c
 *  @param deviceId 接收该数据的设备ID
 *  @param qos 是否需要设备应答，默认为0。
	0：最多发送一次，不关心设备是否响应
	1：至少发送一次，如果设备收到命令后没有应答，则会在下一次设备登录时若命令在有效期内（有效期定义参见timeout参数）则会重发该命令
 *  @param timeout 命令有效时间，默认0。有效范围：0~2678400
 *  @param cmd 用户自定义数据：json、string、二进制数据（小于64K）
 *  @return res 响应数据
 *  @return err 错误
*/
func (c *Client) SendCmd(deviceId string, qos, timeout int, cmd string) (res onenetmod.SendCMDRes, err error) {
	// 处理Url
	paramMap := make(map[string]string, 1)
	paramMap["device_id"] = deviceId
	if qos != 0 {
		paramMap["qos"] = strconv.Itoa(qos)
	}
	if timeout != 0 {
		paramMap["timeout"] = strconv.Itoa(timeout)
	}
	var url string
	url, err = c.getUrlPathWithParams(constant.SendCmd, paramMap)
	if err != nil {
		return
	}
	// 发起请求
	var body []byte
	body, err = c.RequestClient(url, "POST",
		ApiKey(c.MasterAPIkey),
		Body(cmd),
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

// GetCmdStatus
/**
 *  @Description: 查询命令状态
 *  @receiver c
 *  @param cmdUUID 命令ID
 *  @return res 响应参数
 *  @return err 错误
 */
func (c *Client) GetCmdStatus(cmdUUID string) (res onenetmod.GetCMDStatusRes, err error) {
	// 拼接Url
	url := constant.BaseUrl + fmt.Sprintf(constant.GetCmdStatus, cmdUUID)
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

// GetCmdResponse
/**
 *  @Description: 查询命令响应（只有当命令状态为“设备正常响应”时，API才有效）
 *  @receiver c
 *  @param cmdUUID 命令ID
 *  @return res 响应参数
 *  @return err 错误
 */
func (c *Client) GetCmdResponse(cmdUUID string) (res string, err error) {
	// 拼接Url
	url := constant.BaseUrl + fmt.Sprintf(constant.GetCmdResponse, cmdUUID)
	// 发起请求
	var body []byte
	body, err = c.RequestClient(url, "GET", ApiKey(c.MasterAPIkey))
	// 抛出错误
	if err != nil {
		return
	}
	// 将返回数据转换为字符串
	res = string(body)
	return
}

// GetDeviceHistoryCmd
/**
 *  @Description: 查询设备历史命令
 *  @receiver c
 *  @param deviceId 设备ID
 *  @param start 查询的开始时间（必填）
 *  @param end 查询的结束时间
 *  @param page 指定页码
 *  @param perPage 指定每页输出设备个数，默认30，最多100
 *  @return res 响应参数
 *  @return err 错误
 */
func (c *Client) GetDeviceHistoryCmd(deviceId string, start, end *time.Time, page, perPage int) (res onenetmod.GetHistoryCMDRes, err error) {
	// 处理Url
	paramMap := make(map[string]string, 1)
	if start != nil {
		paramMap["start"] = start.Format("2006-01-02T15:04:05")
	}
	if end != nil {
		paramMap["start"] = end.Format("2006-01-02T15:04:05")
	}
	if page> 1{
		paramMap["page"] = strconv.Itoa(page)
	}
	if perPage> 1{
		paramMap["per_page"] = strconv.Itoa(perPage)
	}
	var url string
	url, err = c.getUrlPathWithParams(fmt.Sprintf(constant.GetCmdHistory, deviceId), paramMap)
	if err != nil {
		return
	}
	// 拼接Url
	url = constant.BaseUrl + url
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
