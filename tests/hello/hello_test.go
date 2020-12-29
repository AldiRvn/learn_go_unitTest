package hello_test

import (
	"hello_world/hello"
	"testing"
)

func TestHelloWorld(t *testing.T) {
	expect := "Hello World!"
	actual := hello.HelloWorld()
	if actual != expect {
		t.Errorf("expect %s and not %s", expect, actual)
	}
}
