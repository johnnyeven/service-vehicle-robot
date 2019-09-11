package middlewares

import (
	"fmt"
	"github.com/henrylee2cn/teleport"
	"github.com/johnnyeven/service-vehicle-robot/constants/errors"
	"reflect"
)

type AuthPlugin struct {
}

func (*AuthPlugin) Name() string {
	return "AuthPlugin"
}

func (*AuthPlugin) PostReadCallBody(ctx tp.ReadCtx) *tp.Status {
	return nil
}

func (*AuthPlugin) PostReadPushBody(ctx tp.ReadCtx) *tp.Status {
	rv := reflect.ValueOf(ctx.Input().Body())
	rv = rv.Elem()
	fmt.Println(rv.Interface())
	if rv.Kind() == reflect.Struct {
		tokenV := rv.FieldByName("Token")
		if !tokenV.IsValid() {
			return tp.NewStatus(int32(errors.Forbidden), "", errors.Forbidden.StatusError())
		}
		if token, ok := tokenV.Interface().(string); ok {
			fmt.Println("token: ", token)
		} else {
			return tp.NewStatus(int32(errors.BadRequest), "", errors.BadRequest.StatusError())
		}
	}
	return nil
}
