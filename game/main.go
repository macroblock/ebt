package game

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
	size.MaxInt(minFieldSize, minFieldSize)
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
