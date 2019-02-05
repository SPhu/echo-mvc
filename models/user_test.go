package models

import (
	"reflect"
	"testing"
)

func Test_GetUserid(t *testing.T) {
	user := new(User)
	userid := user.GetUserid("test", "test")
	v := reflect.ValueOf(userid)
	t.Log(v.Type())
}
