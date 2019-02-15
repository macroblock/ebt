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

// NewPoint2i -
func NewPoint2i(x, y int) Point2i {
	return Point2i{x, y}
}

// Random -
func Random(from, to Point2i) Point2i {
	diff := Diff(from, to)
	return from.AddInt(rand.Intn(diff.X+1), rand.Intn(diff.Y+1))
}

// RandomInt -
func RandomInt(from, to int) int {
	diff := to - from
	from += rand.Intn(diff + 1)
	return from
}

// Min -
func Min(a, b Point2i) Point2i {
	return Point2i{misc.MinInt(a.X, b.X), misc.MinInt(a.Y, b.Y)}
}

// Max -
func Max(a, b Point2i) Point2i {
	return Point2i{misc.MaxInt(a.X, b.X), misc.MaxInt(a.Y, b.Y)}
}

// MinMax -
func MinMax(a, b Point2i) (min Point2i, max Point2i) {
	min = Min(a, b)
	max = Max(a, b)
	return
}

// LessThan -
func LessThan(a, b Point2i) bool {
	if a.X >= b.X {
		return false
	}
	if a.Y >= b.Y {
		return false
	}
	return true
}

// LessOrEqual -
func LessOrEqual(a, b Point2i) bool {
	if a.X > b.X {
		return false
	}
	if a.Y > b.Y {
		return false
	}
	return true
}

// GreaterThan -
func GreaterThan(a, b Point2i) bool {
	if a.X <= b.X {
		return false
	}
	if a.Y <= b.Y {
		return false
	}
	return true
}

// GreaterOrEqual -
func GreaterOrEqual(a, b Point2i) bool {
	if a.X < b.X {
		return false
	}
	if a.Y < b.Y {
		return false
	}
	return true
}

// Diff - calculates: max(pt1, pt2) - min(pt1, pt2)
func Diff(a, b Point2i) Point2i {
	return Max(a, b).Sub(Min(a, b))
}

// Set -
func (o Point2i) Set(pt Point2i) Point2i {
	o.X = pt.X
	o.Y = pt.Y
	return o
}

// Add -
func (o Point2i) Add(pt Point2i) Point2i {
	o.X += pt.X
	o.Y += pt.Y
	return o
}

// Sub -
func (o Point2i) Sub(pt Point2i) Point2i {
	o.X -= pt.X
	o.Y -= pt.Y
	return o
}

// Mul -
func (o Point2i) Mul(pt Point2i) Point2i {
	o.X *= pt.X
	o.Y *= pt.Y
	return o
}

// Div -
func (o Point2i) Div(pt Point2i) Point2i {
	o.X /= pt.X
	o.Y /= pt.Y
	return o
}

// Volume -
func (o Point2i) Volume() int {
	return o.X * o.Y
}

// AddInt -
func (o Point2i) AddInt(dx, dy int) Point2i {
	o.X += dx
	o.Y += dy
	return o
}

// SubInt -
func (o Point2i) SubInt(dx, dy int) Point2i {
	o.X -= dx
	o.Y -= dy
	return o
}

// MulInt -
func (o Point2i) MulInt(kx, ky int) Point2i {
	o.X *= kx
	o.Y *= ky
	return o
}

// DivInt -
func (o Point2i) DivInt(kx, ky int) Point2i {
	o.X /= kx
	o.Y /= ky
	return o
}

// Scale -
func (o Point2i) Scale(kx, ky float64) Point2i {
	o.X = int(math.Round(kx * float64(o.X)))
	o.Y = int(math.Round(ky * float64(o.Y)))
	return o
}
