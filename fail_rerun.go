package assert

import (
	"reflect"
	"testing"
	"unsafe"
)

func (a *Assertions) Run(name string, f func(a *Assertions)) {
	t, ok := a.t.(*testing.T)
	if !ok {
		a.FailNow("Assertions.Run must run on *testing.T")
		return
	}
	t.Run(name, func(_ *testing.T) {
		for i := a.failRerun; i > 0; i-- {
			f(a)
			if !t.Failed() {
				break
			}
			if i > 1 {
				setTestingNotFail(t)
			}
		}
	})
}

func setTestingNotFail(t *testing.T) {
	commonV := reflect.ValueOf(t).Elem().FieldByName("common")
	failedV := commonV.FieldByName("failed")
	parentV := commonV.FieldByName("parent")
	failedV = reflect.NewAt(failedV.Type(), unsafe.Pointer(failedV.UnsafeAddr())).Elem()
	if failedV.Bool() {
		failedV.Set(reflect.ValueOf(false))
	}
	for parentV.IsValid() && !parentV.IsZero() {
		if parentV.Kind() == reflect.Ptr {
			parentV = parentV.Elem()
		}
		failedV = parentV.FieldByName("failed")
		parentV = parentV.FieldByName("parent")
		failedV = reflect.NewAt(failedV.Type(), unsafe.Pointer(failedV.UnsafeAddr())).Elem()
		if failedV.Bool() {
			failedV.Set(reflect.ValueOf(false))
		}
	}
	return
}
