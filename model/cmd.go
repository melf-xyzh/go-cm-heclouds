/**
 * @Time    :2022/9/7 16:34
 * @Author  :Xiaoyu.Zhang
 */

package onenetmod

// SendCMDRes 发送命令响应参数
type SendCMDRes struct {
	CommonRes
	Data struct {
		CmdUuid string `json:"cmd_uuid"`
	} `json:"data"`
}

// GetCMDStatusRes 查询命令状态响应参数
type GetCMDStatusRes struct {
	CommonRes
	Data struct {
		Status int    `json:"status"` // 命令状态，见命令状态说明
		Desc   string `json:"desc"`   // 状态描述
	} `json:"data"`
}

// GetHistoryCMDRes 查询设备历史命令响应参数
type GetHistoryCMDRes struct {
	CommonRes
	Data struct {
		PageRes
		Items []struct {
			CmdUuid     string `json:"cmd_uuid"`     // 命令ID
			DeviceId    int    `json:"device_id"`    // 设备ID
			ExpireTime  string `json:"expire_time"`  // 过期时间
			Status      int    `json:"status"`       // 命令状态
			SendTime    string `json:"send_time"`    // 发送时间
			ConfirmTime string `json:"confirm_time"` // 响应时间
			ConfirmBody string `json:"confirm_body"` // 请求内容， 16进制字符
			Body        string `json:"body"`         // 请求内容， 16进制字符
		} `json:"items"`
	} `json:"data"`
}
