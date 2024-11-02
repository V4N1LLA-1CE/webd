package pipeline

import (
	"github.com/v4n1lla-1ce/webd/internal/cli"
	"github.com/v4n1lla-1ce/webd/internal/codec"
)

func Webp2PNG(args cli.Arguments) {
	conversionPipeline := &ConversionPipeline{
		sourceExt: "webp",
		targetExt: "png",
		decoder:   codec.DecodeWebp,
		encoder:   codec.EncodeToPng,
		saver:     codec.SaveToDisk,
	}

	conversionPipeline.Convert(args)
}
