package urltool

import (
	"testing"

	c "github.com/smartystreets/goconvey/convey" // 别名导入
)

func TestGetBasePath(t *testing.T) {
	c.Convey("基础用例", t, func() {
		var (
			s string = "https://www.baidu.com/akjha"

			expect = "akjha"
		)
		got, err := GetBasePath(s)
		c.So(err, c.ShouldBeEmpty)
		c.So(got, c.ShouldEqual, expect) // 断言
	})

	c.Convey("用例2", t, func() {
		var (
			s string = "www.baidu.com/akjha"

			expect = ""
		)
		got, err := GetBasePath(s)
		c.So(err, c.ShouldBeError)
		c.So(got, c.ShouldEqual, expect) // 断言
	})

}
