package main

import (
	"runtime"
	"testing"
)

func TestGetOSType(t *testing.T) {
	expected := runtime.GOOS
	actual := "linux"

	if(expected != actual) {
		t.Errorf("%q != %q", actual, expected)
	}
}
