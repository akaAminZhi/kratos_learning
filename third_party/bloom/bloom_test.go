package bloom

import (
	"testing"

	c "github.com/smartystreets/goconvey/convey" // 别名导入
)

func TestBloom(t *testing.T) {
	LocalInit()
	c.Convey("基础用例", t, func() {
		var (
			s string = "akjha"

			expect = true
		)
		GetBloom().AddString(s)

		got := GetBloom().Test([]byte(s))
		c.So(got, c.ShouldEqual, expect) // 断言
	})

}
