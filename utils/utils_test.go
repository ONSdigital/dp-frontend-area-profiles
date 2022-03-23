package utils

import (
	"testing"

	"github.com/ONSdigital/dp-frontend-area-profiles/mapper"
	. "github.com/smartystreets/goconvey/convey"
)

func TestSetErrorDetails(t *testing.T) {
	Convey("set correct error details", t, func() {
		var actual mapper.ErrorDetails
		SetErrorDetails(404, &actual)
		expected := mapper.ErrorDetails{
			Description: ErrorDescriptionFor404,
			Title:       ErrorTitleFor404,
		}
		So(actual, ShouldResemble, expected)
	})
	Convey("set correct error details", t, func() {
		var actual mapper.ErrorDetails
		SetErrorDetails(500, &actual)
		expected := mapper.ErrorDetails{
			Description: ErrorDescriptionFor500,
			Title:       ErrorTitleFor500,
		}
		So(actual, ShouldResemble, expected)
	})
}
