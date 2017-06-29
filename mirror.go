package mirror

import "image"

func MirrorImage(in image.Image, ox bool, oy bool) image.Image {
	bounds := in.Bounds()
	out := image.NewRGBA(bounds)

	for y := bounds.Min.Y; y <= bounds.Max.Y; y++ {
		for x := bounds.Min.X; x <= bounds.Max.X; x++ {
			c := in.At(x, y)
			nx, ny := x, y
			if ox {
				nx = bounds.Max.X - x - 1
			}
			if oy {
				ny = bounds.Max.Y - y - 1
			}
			out.Set(nx, ny, c)
		}
	}

	return out
}
