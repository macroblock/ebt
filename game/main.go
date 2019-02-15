package game

import (
	"github.com/macroblock/imed/pkg/misc"
)

const (
	minFieldSize = 10
)

// Field -
type Field struct {
	size Point2i
	grid [][]int
}

// NewField -
func NewField(size Point2i) *Field {
	size = Max(size, Point2i{minFieldSize, minFieldSize})
	grid := make([][]int, size.Y)
	for j := range grid {
		line := make([]int, size.X)
		for i := range line {
			line[i] = -1
		}
		grid[j] = line
	}
	field := &Field{size: size, grid: grid}
	return field
}

// NewFieldInt -
func NewFieldInt(w, h int) *Field {
	return NewField(Point2i{w, h})
}

// Grid -
func (o *Field) Grid() [][]int {
	return o.grid
}

// Size -
func (o *Field) Size() Point2i {
	return o.size
}

// Put -
func (o *Field) Put(p Point2i, id int) {
	o.grid[p.Y][p.X] = id
}

// FillRect -
func (o *Field) FillRect(p1, p2 Point2i, id int) {
	for j := p1.Y; j < p2.Y+1; j++ {
		for i := p1.X; i < p2.X+1; i++ {
			o.grid[j][i] = id
		}
	}
}

// DrawLink -
func (o *Field) DrawLink(p1, p2 Point2i, id int) {
	x1 := misc.MinInt(p1.X, p2.X)
	x2 := misc.MaxInt(p1.X, p2.X)
	for x := x1; x < x2+1; x++ {
		o.grid[p1.Y][x] = id
	}

	y1 := misc.MinInt(p1.Y, p2.Y)
	y2 := misc.MaxInt(p1.Y, p2.Y)
	for y := y1; y < y2; y++ {
		o.grid[y][p2.X] = id
	}
}
