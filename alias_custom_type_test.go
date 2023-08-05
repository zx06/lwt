package lwt

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

// 类型别名
type AliasInt = int
// 自定义类型
type MyInt int


func TestAliasCustomType(t *testing.T) {
	t.Run("AliasInt", func(t *testing.T) {
		// AliasInt和int是同一个类型
		var (
			a AliasInt = 1
			b int      = 1
		)
		aType:=reflect.TypeOf(a) // int
		bType:=reflect.TypeOf(b) // int
		assert.Equal(t, aType, bType)

		// AliasInt和int的Kind都是int
		vA:=reflect.ValueOf(a)
		vB:=reflect.ValueOf(b)
		assert.Equal(t, reflect.Int, vA.Kind())
		assert.Equal(t, reflect.Int, vB.Kind())
	})
	t.Run("MyInt", func(t *testing.T) {
		// MyInt和int是不同的类型
		var (
			a MyInt = 1
			b int   = 1
		)
		aType:=reflect.TypeOf(a) // lwt.MyInt
		bType:=reflect.TypeOf(b) // int
		assert.NotEqual(t, aType, bType)

		// MyInt和int的Kind都是int
		vA:=reflect.ValueOf(a)
		vB:=reflect.ValueOf(b)
		assert.Equal(t, reflect.Int, vA.Kind())
		assert.Equal(t, reflect.Int, vB.Kind())
	})
}
