package middlewares

import (
	"github.com/henrylee2cn/teleport"
)

type AuthPlugin struct {
}

func (*AuthPlugin) Name() string {
	return "AuthPlugin"
}

func (*AuthPlugin) PostReadCallBody(ctx tp.ReadCtx) *tp.Status {

	return nil
}
