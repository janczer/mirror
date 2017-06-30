package mirror

import "image"

func MirrorImage(in image.Image, ox bool, oy bool) image.Image {
	bounds := in.Bounds()
	out := image.NewRGBA(bounds)

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			c := in.At(x, y)
			nx, ny := x, y
			if ox {
				nx = bounds.Max.X - 1 - x
			}
			if oy {
				ny = bounds.Max.Y - 1 - y
			}
			out.Set(nx, ny, c)
		}
	}

	return out
}
