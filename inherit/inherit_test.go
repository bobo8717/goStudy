package inherit_test

import (
	"goTest/inherit"
	"testing"
)

func Test_1(t *testing.T) {
	m := inherit.CreateM2()
	m.Test2()
	m.ShowT1()
}
