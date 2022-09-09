/**
 * @Time    :2022/9/4 14:47
 * @Author  :MELF晓宇
 * @Email   :xyzh.melf@petalmail.com
 * @FileName:request.go
 * @Project :go-cm-heclouds
 * @Blog    :https://blog.csdn.net/qq_29537269
 * @Guide   :https://guide.melf.space
 * @Information:
 *
 */

package onenet

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"hash"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

type (
	// ContentType is an option used for RequestClientOptions
	ContentType string
	// ApiKey is an option used for RequestClientOptions
	ApiKey string
	// AccessKey is an option used for RequestClientOptions
	AccessKey string
	// SafeSign is an option used for RequestClientOptions
	SafeSign bool // 使用安全方式签名
	// Body is an option used for RequestClientOptions
	Body string
)

type RequestClientOptions interface {
	setRequestClientOption(formatPr *RequestClientPr)
}

type RequestClientPr struct {
	ContentType string // ContentType
	Body        string // RequestBody
	ApiKey      string // ApiKey
	AccessKey   string // AccessKey
	SafeSign    bool   // 是否使用安全方式签名
}

func (o ContentType) setRequestClientOption(pr *RequestClientPr) {
	pr.ContentType = string(o)
}

func (o ApiKey) setRequestClientOption(pr *RequestClientPr) {
	pr.ApiKey = string(o)
}

func (o AccessKey) setRequestClientOption(pr *RequestClientPr) {
	pr.AccessKey = string(o)
}

func (o SafeSign) setRequestClientOption(pr *RequestClientPr) {
	pr.SafeSign = bool(o)
}

func (o Body) setRequestClientOption(pr *RequestClientPr) {
	pr.Body = string(o)
}

// RequestClient
/**
 * @Description: http请求Client
 * @receiver c
 * @param url
 * @param Method
 * @param opts
 * @return body
 * @return err
 */
func (c *Client) RequestClient(url string, Method string, opts ...RequestClientOptions) (body []byte, err error) {
	pr := new(RequestClientPr)
	for _, opt := range opts {
		opt.setRequestClientOption(pr)
	}

	// 创建一个HttpClient
	var client http.Client
	if c.TimeOut == 0 {
		client = http.Client{}
	} else {
		// 设置超时时间
		client = http.Client{
			Timeout: c.TimeOut,
		}
	}

	payload := strings.NewReader(pr.Body)
	var req *http.Request
	switch Method {
	case "GET":
		// 发送Get请求
		req, err = http.NewRequest("GET", url, nil)
	case "POST":
		// 发送Get请求
		req, err = http.NewRequest("POST", url, payload)
	case "PUT":
		// 发送Get请求
		req, err = http.NewRequest("PUT", url, payload)
	case "DELETE":
		// 发送Get请求
		req, err = http.NewRequest("DELETE", url, payload)
	case "PATCH":
		// 发送Get请求
		req, err = http.NewRequest("PATCH", url, payload)
	default:
		err = errors.New("暂不支持的请求")
		return
	}
	if err != nil {
		return
	}
	if Method != "GET" {
		// 设置Content-Type
		req.Header.Add("Content-Type", pr.ContentType)
	}
	if pr.SafeSign {
		// 使用安全模式签名
		req.Header.Add("Authorization", pr.ApiKey)
	} else {
		// 使用普通模式签名
		req.Header.Add("api-key", pr.ApiKey)
	}
	var res *http.Response
	// 发起Http请求
	res, err = client.Do(req)
	if err != nil {
		return
	}
	// 这步是必要的，防止以后的内存泄漏，切记
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		err = errors.New("Http状态码非200，" + res.Status)
		return
	}
	// 读取响应 body, 返回为 []byte
	body, err = ioutil.ReadAll(res.Body)
	return
}

// getUrlPathWithParams
/**
 * @Description: 获取带参数的URL
 * @receiver c
 * @param urlPath
 * @param paramMap
 * @return urlPathNew
 * @return err
 */
func (c *Client) getUrlPathWithParams(urlPath string, paramMap map[string]string) (urlPathNew string, err error) {
	params := url.Values{}
	var parseURL *url.URL
	parseURL, err = url.Parse(urlPath)
	if err != nil {
		return "", err
	}
	if paramMap != nil {
		for k, v := range paramMap {
			params.Set(k, v)
		}
	}
	// 如果参数中有中文参数,这个方法会进行URLEncode
	parseURL.RawQuery = params.Encode()
	urlPathNew = parseURL.String()
	return urlPathNew, nil
}

// Sign
/**
 * @Description: 签名
 * @param version
 * @param res
 * @param method
 * @param apiKey
 * @param et
 * @return signature
 */
func Sign(version, res, method, apiKey string, et int64) (signature string) {
	// 参数组版本号，日期格式，目前仅支持"2018-10-31"
	version = "2018-10-31"
	// 对access_key进行decode
	keyBytes, _ := base64.StdEncoding.DecodeString(apiKey)
	// 计算签名的字符串
	var build strings.Builder
	build.WriteString(strconv.FormatInt(et, 10))
	build.WriteString("\n")
	build.WriteString(method)
	build.WriteString("\n")
	build.WriteString(res)
	build.WriteString("\n")
	build.WriteString(version)
	stringForSignature := build.String()
	textBytes := []byte(stringForSignature)
	// 调用签名方法
	var h hash.Hash
	switch method {
	case "md5":
		h = hmac.New(md5.New, keyBytes)
	case "sha1":
		h = hmac.New(sha1.New, keyBytes)
	case "sha256":
		h = hmac.New(sha256.New, keyBytes)
	}
	if h == nil {
		return
	}
	h.Write(textBytes)
	result := h.Sum(nil)
	signStr := base64.StdEncoding.EncodeToString(result)

	// URL编码
	signStr = quote(signStr)
	res = quote(res)

	var signatureBuild strings.Builder
	signatureBuild.WriteString("version=")
	signatureBuild.WriteString(version)
	signatureBuild.WriteString("&res=")
	signatureBuild.WriteString(res)
	signatureBuild.WriteString("&et=")
	signatureBuild.WriteString(strconv.FormatInt(et, 10))
	signatureBuild.WriteString("&method=")
	signatureBuild.WriteString(method)
	signatureBuild.WriteString("&sign=")
	signatureBuild.WriteString(signStr)
	signature = signatureBuild.String()
	return
}

func quote(str string) (answer string) {
	str = strings.ReplaceAll(str, "+", "%2B")
	str = strings.ReplaceAll(str, " ", "%20")
	str = strings.ReplaceAll(str, "/", "%2F")
	str = strings.ReplaceAll(str, "?", "%3F")
	str = strings.ReplaceAll(str, "%", "%25")
	str = strings.ReplaceAll(str, "#", "%23")
	str = strings.ReplaceAll(str, "&", "%26")
	str = strings.ReplaceAll(str, "=", "%3D")
	return str
}
