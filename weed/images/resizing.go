package images

import (
	"bytes"
	"image"
	"github.com/chrislusf/seaweedfs/weed/glog"
	"github.com/disintegration/imaging"
	"image/jpeg"
	"image/gif"
	"image/png"
	"fmt"
)

func Resized(ext string, data []byte, width, height int, mode string) (resized []byte, w int, h int) {
	if width == 0 && height == 0 {
		return data, 0, 0
	}
	srcImage, _, err := image.Decode(bytes.NewReader(data))
	if err == nil {
		bounds := srcImage.Bounds()
		var dstImage *image.NRGBA
		if bounds.Dx() > width && width != 0 || bounds.Dy() > height && height != 0 {
			switch mode {
			case "fit":
				dstImage = imaging.Fit(srcImage, width, height, imaging.Lanczos)
			case "fill":
				dstImage = imaging.Fill(srcImage, width, height, imaging.Center, imaging.Lanczos)
			default:
				if width == height && bounds.Dx() != bounds.Dy() {
					dstImage = imaging.Thumbnail(srcImage, width, height, imaging.Lanczos)
					w, h = width, height
				} else {
					dstImage = imaging.Resize(srcImage, width, height, imaging.Lanczos)
				}
			}
		} else {
			return data, bounds.Dx(), bounds.Dy()
		}
		fmt.Println("ext:",ext)
		var buf bytes.Buffer
		if ext == "" {
			err = png.Encode(&buf, dstImage)
			if err!=nil{
				jpeg.Encode(&buf, dstImage, nil)
			}
		}else{
			switch ext {
			case ".png":
				png.Encode(&buf, dstImage)
			case ".jpg", ".jpeg":
				jpeg.Encode(&buf, dstImage, nil)
			case ".gif":
				gif.Encode(&buf, dstImage, nil)
			}
		}

		return buf.Bytes(), dstImage.Bounds().Dx(), dstImage.Bounds().Dy()
	} else {
		glog.Error(err)
	}
	return data, 0, 0
}
