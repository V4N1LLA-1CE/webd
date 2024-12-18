package pipeline

import (
	"github.com/v4n1lla-1ce/webd/internal/cli"
	"github.com/v4n1lla-1ce/webd/internal/codec"
)

func Webp2Png(args cli.Arguments) {
	c := &ConversionPipeline{
		sourceExt: "webp",
		targetExt: "png",
		decoder:   codec.DecodeWebp,
		encoder:   codec.EncodeToPng,
		saver:     codec.SaveToDisk,
	}

	c.Convert(args)
}

func PNG2Webp(args cli.Arguments) {
	c := &ConversionPipeline{
		sourceExt: "png",
		targetExt: "webp",
		decoder:   codec.DecodePng,
		encoder:   codec.EncodeToWebp,
		saver:     codec.SaveToDisk,
	}

	c.Convert(args)
}

func Jpg2Png(args cli.Arguments) {
	c := &ConversionPipeline{
		sourceExt: "jpg",
		targetExt: "png",
		decoder:   codec.DecodeJpg,
		encoder:   codec.EncodeToPng,
		saver:     codec.SaveToDisk,
	}

	c.Convert(args)
}

func Png2Jpg(args cli.Arguments) {
	c := &ConversionPipeline{
		sourceExt: "png",
		targetExt: "jpg",
		decoder:   codec.DecodePng,
		encoder:   codec.EncodeToJpg,
		saver:     codec.SaveToDisk,
	}

	c.Convert(args)
}

func Jpg2Webp(args cli.Arguments) {
	c := &ConversionPipeline{
		sourceExt: "jpg",
		targetExt: "webp",
		decoder:   codec.DecodeJpg,
		encoder:   codec.EncodeToWebp,
		saver:     codec.SaveToDisk,
	}

	c.Convert(args)
}

func Webp2Jpg(args cli.Arguments) {
	c := &ConversionPipeline{
		sourceExt: "webp",
		targetExt: "jpg",
		decoder:   codec.DecodeWebp,
		encoder:   codec.EncodeToJpg,
		saver:     codec.SaveToDisk,
	}

	c.Convert(args)
}
