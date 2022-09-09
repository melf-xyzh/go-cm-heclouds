/**
 * @Time    :2022/9/7 14:11
 * @Author  :Xiaoyu.Zhang
 */

package onenetmod

// CommonRes 公共返回参数
type CommonRes struct {
	Errno int    `json:"errno"` // 调用错误码，为0表示调用成功
	Error string `json:"error"` // 错误描述，为"succ"表示调用成功
}

// PageRes 分页返回参数
type PageRes struct {
	Page       int `json:"page"`        // 当前页码
	PerPage    int `json:"per_page"`    // 当前每页总数
	TotalCount int `json:"total_count"` // 总数
}
