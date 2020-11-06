package strlib

import (
	"testing"

	"github.com/smartystreets/goconvey/convey"
)

func TestJoinInts(t *testing.T) {
	convey.Convey("TestJoinInts int", t, func(ctx convey.C) {
		var (
			data = []int{1, 2, 3}
		)
		ctx.Convey("When everything gose positive-1", func(ctx convey.C) {
			res := JoinInts(data)
			ctx.So(res, convey.ShouldEqual, "1,2,3")
		})
		ctx.Convey("When everything gose positive-2", func(ctx convey.C) {
			res := JoinInts(&data)
			ctx.So(res, convey.ShouldEqual, "1,2,3")
		})
	})

	convey.Convey("TestJoinInts uint", t, func(ctx convey.C) {
		var (
			data = []uint{1, 2, 3}
		)
		ctx.Convey("When everything gose positive", func(ctx convey.C) {
			res := JoinInts(data)
			ctx.So(res, convey.ShouldEqual, "1,2,3")
		})
		ctx.Convey("When everything gose positive-2", func(ctx convey.C) {
			res := JoinInts(&data)
			ctx.So(res, convey.ShouldEqual, "1,2,3")
		})
	})
}

func TestSplitInts(t *testing.T) {
	convey.Convey("TestSplitInts int", t, func(ctx convey.C) {
		var (
			data = "1,2,3"
		)
		ctx.Convey("When everything gose positive-1", func(ctx convey.C) {
			var ret []int
			err := SplitInts(data, &ret)
			ctx.So(err, convey.ShouldBeNil)
		})
		ctx.Convey("When everything gose positive-2", func(ctx convey.C) {
			var ret []int
			_ = SplitInts(data, &ret)
			ctx.So(len(ret), convey.ShouldEqual, 3)
		})
	})

	convey.Convey("TestSplitInts uint", t, func(ctx convey.C) {
		var (
			data = "1,2,3"
		)
		ctx.Convey("When everything gose positive-1", func(ctx convey.C) {
			var ret []uint
			err := SplitInts(data, &ret)
			ctx.So(err, convey.ShouldBeNil)
		})
		ctx.Convey("When everything gose positive-2", func(ctx convey.C) {
			var ret []uint
			_ = SplitInts(data, &ret)
			ctx.So(len(ret), convey.ShouldEqual, 3)
		})
	})
}
