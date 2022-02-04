package utils

import (
	"errors"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestRequestUtils(t *testing.T) {

	Convey("#GetFirstError()", t, func() {
		err1 := errors.New("#1")
		err2 := errors.New("#2")
		err3 := errors.New("#3")
		expect1 := GetFirstError(err1, nil, nil)
		expect2 := GetFirstError(nil, nil, err3)
		expect3 := GetFirstError(nil, err2, err3)
		expect4 := GetFirstError(err1, err2, err3)
		expect5 := GetFirstError(nil, nil, nil)
		So(expect1, ShouldEqual, err1)
		So(expect2, ShouldEqual, err3)
		So(expect3, ShouldEqual, err2)
		So(expect4, ShouldEqual, err1)
		So(expect5, ShouldEqual, nil)
	})
}
