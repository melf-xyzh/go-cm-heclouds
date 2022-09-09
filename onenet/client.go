/**
 * @Time    :2022/9/3 19:36
 * @Author  :MELF晓宇
 * @Email   :xyzh.melf@petalmail.com
 * @FileName:client.go
 * @Project :go-cm-heclouds
 * @Blog    :https://blog.csdn.net/qq_29537269
 * @Guide   :https://guide.melf.space
 * @Information:
 *
 */

package onenet

import (
	"github.com/json-iterator/go"
	"github.com/melf-xyzh/go-cm-heclouds/constant"
	"time"
)

var (
	json = jsoniter.ConfigCompatibleWithStandardLibrary
)

// Client 请求客户端
type Client struct {
	MasterAPIkey string // Master-APIkey
	AccessKey    string // access_key
	ProductId    string // 产品ID
	TimeOut      time.Duration
}

// NewOneNetRequestClient
/**
 *  @Description: 创建一个OneNET请求器
 *  @param apiKey
 *  @param accessKey
 *  @param productId
 *  @return client
 */
func NewOneNetRequestClient(apiKey, accessKey, productId string) (client *Client) {
	client = &Client{
		MasterAPIkey: apiKey,
		AccessKey:    accessKey,
		ProductId:    productId,
		TimeOut:      constant.DefaultTimeOut,
	}
	return
}

// SetTimeOut
/**
 *  @Description: 设置超时时间
 *  @receiver c
 *  @param timeOut
 */
func (c *Client) SetTimeOut(timeOut time.Duration) {
	c.TimeOut = timeOut
}
