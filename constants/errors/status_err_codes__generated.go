package errors

import (
	"github.com/johnnyeven/libtools/courier/status_error"
)

func init() {
	status_error.StatusErrorCodes.Register("BadRequest", 400000000, "请求参数错误", "", false)
	status_error.StatusErrorCodes.Register("Unauthorized", 401000000, "未授权", "", true)
	status_error.StatusErrorCodes.Register("Forbidden", 403000000, "不允许操作", "", true)
	status_error.StatusErrorCodes.Register("NotFound", 404000000, "未找到", "", false)
	status_error.StatusErrorCodes.Register("Conflict", 409000000, "操作冲突", "", true)
	status_error.StatusErrorCodes.Register("InternalError", 500000000, "内部处理错误", "", false)
}
