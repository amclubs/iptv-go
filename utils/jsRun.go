// Package list
// @Time:2024/07/22 10:00
// @File:mian.go
// @SoftWare:Goland
// @Author:AM科技
// @Contact:https://t.me/AM_CLUBS

package utils

import (
	"fmt"
	js "github.com/dop251/goja"
	"sync"
)

type JsUtil struct {
	pool sync.Pool
}

func (j *JsUtil) getVm() *js.Runtime {
	v := j.pool.Get()
	if v != nil {
		return v.(*js.Runtime)
	}
	return js.New()
}

func (j *JsUtil) putVm(vm *js.Runtime) {
	vm.Set("global", nil) // 清除全局对象
	j.pool.Put(vm)
}

func (j *JsUtil) JsRun(funcContent []string, params ...any) any {
	vm := j.getVm()
	defer j.putVm(vm)
	_, err := vm.RunString(funcContent[0])
	if err != nil {
		return err
	}
	jsfn, ok := js.AssertFunction(vm.Get(funcContent[1]))
	if !ok {
		return fmt.Errorf("执行函数失败")
	}
	jsValues := make([]js.Value, 0, len(params))
	for _, v := range params {
		jsValues = append(jsValues, vm.ToValue(v))
	}
	result, err := jsfn(
		js.Undefined(),
		jsValues...,
	)
	if err != nil {
		return err
	}
	return result
}