package pipeline

import (
	"github.com/v4n1lla-1ce/webd/internal/cli"
	"github.com/v4n1lla-1ce/webd/internal/codec"
)

func Webp2PNG(args cli.Arguments) {
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
