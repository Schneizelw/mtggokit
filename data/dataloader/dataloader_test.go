package dataloader

import (
	"github.com/smartystreets/goconvey/convey"
	"mtggokits/data/container"
	"testing"
)

type FakeStreamer struct {
}

func (*FakeStreamer) SetContainer(container.Container) {

}

func (*FakeStreamer) GetContainer() container.Container {
	return nil
}

func (*FakeStreamer) UpdateData() error {
	return nil
}

func TestLoader_Register(t *testing.T) {
	convey.Convey("Test register duplicate name", t, func() {
		loader := NewLoader()
		convey.So(loader, convey.ShouldNotBeNil)
	})

	convey.Convey("Test register name", t, func() {
		loader := NewLoader()
		convey.So(loader, convey.ShouldNotBeNil)
		loader.Register("abc", &FakeStreamer{})
	})

	convey.Convey("Test register duplicate name", t, func() {
		loader := NewLoader()
		convey.So(loader, convey.ShouldNotBeNil)
		convey.So(loader.Register("abc", &FakeStreamer{}), convey.ShouldBeNil)
		e := loader.Register("abc", &FakeStreamer{})
		convey.So(e, convey.ShouldNotBeNil)
		convey.So(e.Error(), convey.ShouldEqual, "streamer[abc] has already exist")
	})
}
