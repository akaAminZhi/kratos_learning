package base62

import (
	"testing"

	c "github.com/smartystreets/goconvey/convey" // 别名导入
)

func TestBase10ToBase62(t *testing.T) {
	c.Convey("基础用例", t, func() {
		var (
			s uint64 = 63

			expect = "11"
		)
		got := Base10ToBase62(s)
		c.So(got, c.ShouldEqual, expect) // 断言
	})

	c.Convey("用例2", t, func() {
		var (
			s uint64 = 6347

			expect = "1En"
		)
		got := Base10ToBase62(s)
		c.So(got, c.ShouldEqual, expect) // 断言
	})

}

func TestBase62StrToBase10(t *testing.T) {
	c.Convey("基础用例", t, func() {
		var (
			s string = "11"

			expect = 63
		)
		got := Base62StrToBase10(s)
		c.So(got, c.ShouldEqual, expect) // 断言
	})

	c.Convey("用例2", t, func() {
		var (
			s string = "1En"

			expect = 6347
		)
		got := Base62StrToBase10(s)
		c.So(got, c.ShouldEqual, expect) // 断言
	})

}
