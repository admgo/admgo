package response

import (
	"context"
	"reflect"
)

type okRsp struct {
	Code int32       `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data,optional,omitempty"`
}

func OKResponse(ctx context.Context, v any) any {
	val := reflect.ValueOf(v)

	// 如果传入的是指针，则解引用
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	switch val.Kind() {
	case reflect.Struct:
		// 如果是结构体，则遍历其字段进行检查
		if val.NumField() == 0 {
			v = nil
		}
		for i := 0; i < val.NumField(); i++ {
			field := val.Field(i)
			// 检查字段是否为空值
			if reflect.DeepEqual(field.Interface(), reflect.Zero(field.Type()).Interface()) {
				v = nil
			}
		}
	case reflect.Interface:
		v = nil
	case reflect.Invalid:
		v = nil
	default:
		// 如果不是结构体，直接检查该值是否为空值
		if reflect.DeepEqual(val.Interface(), reflect.Zero(val.Type()).Interface()) {
			v = nil
		}
	}
	return &okRsp{
		Code: 0,
		Msg:  "success",
		Data: v,
	}
}
