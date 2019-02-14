package game

import (
	"math"
	"math/rand"

	"github.com/macroblock/imed/pkg/misc"
)

type (
	// Point2i -
	Point2i struct {
		X, Y int
	}
)

// Random -
func Random(from, to Point2i) Point2i {
	diff := Diff(from, to)
	from.AddInt(rand.Intn(diff.X+1), rand.Intn(diff.Y+1))
	return from
}

// RandomInt -
func RandomInt(from, to int) int {
	diff := to - from
	from += rand.Intn(diff + 1)
	return from
}

// Diff - calculates (pt2 - pt1)
func Diff(pt1, pt2 Point2i) Point2i {
	pt2.Sub(pt1)
	return pt2
}

// Set -
func (o *Point2i) Set(pt Point2i) *Point2i {
	o.X = pt.X
	o.Y = pt.Y
	return o
}

// Add -
func (o *Point2i) Add(pt Point2i) *Point2i {
	o.X += pt.X
	o.Y += pt.Y
	return o
}

// Sub -
func (o *Point2i) Sub(pt Point2i) *Point2i {
	o.X -= pt.X
	o.Y -= pt.Y
	return o
}

// Mul -
func (o *Point2i) Mul(pt Point2i) *Point2i {
	o.X *= pt.X
	o.Y *= pt.Y
	return o
}

// Div -
func (o *Point2i) Div(pt Point2i) *Point2i {
	o.X /= pt.X
	o.Y /= pt.Y
	return o
}

// Min -
func (o *Point2i) Min(pt Point2i) *Point2i {
	o.X = misc.MinInt(pt.X, o.X)
	o.Y = misc.MinInt(pt.Y, o.Y)
	return o
}

// Max -
func (o *Point2i) Max(pt Point2i) *Point2i {
	o.X = misc.MaxInt(pt.X, o.X)
	o.Y = misc.MaxInt(pt.Y, o.Y)
	return o
}

// Volume -
func (o *Point2i) Volume() int {
	return o.X * o.Y
}

// LessThan -
func (o *Point2i) LessThan(pt Point2i) bool {
	if o.X >= pt.X {
		return false
	}
	if o.Y >= pt.Y {
		return false
	}
	return true
}

// LessOrEqual -
func (o *Point2i) LessOrEqual(pt Point2i) bool {
	if o.X > pt.X {
		return false
	}
	if o.Y > pt.Y {
		return false
	}
	return true
}

// GreaterThan -
func (o *Point2i) GreaterThan(pt Point2i) bool {
	if o.X <= pt.X {
		return false
	}
	if o.Y <= pt.Y {
		return false
	}
	return true
}

// GreaterOrEqual -
func (o *Point2i) GreaterOrEqual(pt Point2i) bool {
	if o.X < pt.X {
		return false
	}
	if o.Y < pt.Y {
		return false
	}
	return true
}

// SetInt -
func (o *Point2i) SetInt(x, y int) *Point2i {
	o.X = x
	o.Y = y
	return o
}

// AddInt -
func (o *Point2i) AddInt(dx, dy int) *Point2i {
	o.X += dx
	o.Y += dy
	return o
}

// SubInt -
func (o *Point2i) SubInt(dx, dy int) *Point2i {
	o.X -= dx
	o.Y -= dy
	return o
}

// MulInt -
func (o *Point2i) MulInt(kx, ky int) *Point2i {
	o.X *= kx
	o.Y *= ky
	return o
}

// DivInt -
func (o *Point2i) DivInt(kx, ky int) *Point2i {
	o.X /= kx
	o.Y /= ky
	return o
}

// MinInt -
func (o *Point2i) MinInt(x, y int) *Point2i {
	o.X = misc.MinInt(x, o.X)
	o.Y = misc.MinInt(y, o.Y)
	return o
}

// MaxInt -
func (o *Point2i) MaxInt(x, y int) *Point2i {
	o.X = misc.MaxInt(x, o.X)
	o.Y = misc.MaxInt(y, o.Y)
	return o
}

// Scale -
func (o *Point2i) Scale(kx, ky float64) *Point2i {
	o.X = int(math.Round(kx * float64(o.X)))
	o.Y = int(math.Round(ky * float64(o.Y)))
	return o
}
