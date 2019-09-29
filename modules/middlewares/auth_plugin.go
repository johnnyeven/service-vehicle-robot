package middlewares

import (
	"github.com/henrylee2cn/teleport"
	"github.com/johnnyeven/service-vehicle-robot/constants/errors"
	"reflect"
)

type AuthPlugin struct {
}

func (*AuthPlugin) Name() string {
	return "AuthPlugin"
}

func (p *AuthPlugin) PostReadCallBody(ctx tp.ReadCtx) *tp.Status {
	method := ctx.Input().ServiceMethod()
	if method == "/authorization/auth" {
		return nil
	}
	return p.checkToken(ctx.Input().Body())
}

func (p *AuthPlugin) PostReadPushBody(ctx tp.ReadCtx) *tp.Status {
	return p.checkToken(ctx.Input().Body())
}

func (p *AuthPlugin) checkToken(body interface{}) *tp.Status {
	rv := reflect.ValueOf(body)
	rv = rv.Elem()
	if rv.Kind() == reflect.Struct {
		tokenV := rv.FieldByName("Token")
		if !tokenV.IsValid() {
			return tp.NewStatus(int32(errors.Forbidden), "", errors.Forbidden.StatusError())
		}
		if _, ok := tokenV.Interface().(string); ok {

		} else {
			return tp.NewStatus(int32(errors.BadRequest), "", errors.BadRequest.StatusError())
		}
	}
	return nil
}
