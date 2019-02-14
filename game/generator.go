package game

import (
	"fmt"
	"math/rand"
)

type (
	// Room -
	Room struct {
		field    *Field
		id       int
		min, max Point2i
		filter   *CheckFilter
	}

	// CheckFilter -
	CheckFilter struct {
		size  Point2i
		pivot Point2i
	}
)

// Generate -
func Generate(minSize, maxSize Point2i, minRooms, maxRooms int) *Field {
	numRooms := RandomInt(minRooms, maxRooms)
	size := Random(minSize, maxSize)
	fmt.Printf("delta: %v\n", size)
	field := NewField(size)
	rooms := []*Room{}
	id := 1
	for i := 0; i < numRooms; i++ {
		size := Random(Point2i{1, 1}, Point2i{7, 7})
		if size.X == 4 {
			size.X = 5
		}
		if size.Y == 4 {
			size.Y = 5
		}
		pivot := size
		pivot.DivInt(2, 2)
		filter := NewCheckFilter(size, pivot)
		for try := 0; try < 50; try++ {
			pos := Random(Point2i{}, field.size)
			// fmt.Printf("x, y: %v, %v\n", x, y)
			if room, ok := NewRoom(field, pos, id, filter); ok {
				rooms = append(rooms, room)

				for i := 0; i < RandomInt(10, 30); i++ {
					dir := rand.Intn(4)
					for i := 0; i < 4; i++ {
						d := (dir + i) % 4
						// if room.canExtend(d) {
						// 	room.Extend(d)
						// 	break
						// }
						if room.Extend(d) {
							break
						}
					}
				}

				fmt.Printf("room %v: min %v, max %v size %v pivot %v\n", id, room.min, room.max, room.filter.size, room.filter.pivot)
				id++
				break
			}
		}
	}

	// for _, room := range rooms[0 : len(rooms)-1] {
	// 	room.filter.size.SetInt(7, 7)
	// 	room.filter.pivot.SetInt(3, 3)
	// }
	// rooms[len(rooms)-1].filter.size.SetInt(3, 3)
	// rooms[len(rooms)-1].filter.pivot.SetInt(1, 1)

	// for _, room := range rooms {
	// 	for i := 0; i < 10+rand.Intn(20); i++ {
	// 		dir := rand.Intn(4)
	// 		for i := 0; i < 4; i++ {
	// 			d := (dir + i) % 4
	// 			// if room.canExtend(d) {
	// 			// 	room.Extend(d)
	// 			// 	break
	// 			// }
	// 			if room.Extend(d) {
	// 				break
	// 			}
	// 		}
	// 	}
	// }

	// for true {
	// 	room := rooms[len(rooms)-1]
	// 	ok := room.Extend(0)
	// 	ok = ok || room.Extend(1)
	// 	ok = ok || room.Extend(2)
	// 	ok = ok || room.Extend(3)
	// 	if !ok {
	// 		break
	// 	}
	// }

	return field
}

// NewRoom -
func NewRoom(field *Field, pos Point2i, id int, filter *CheckFilter) (*Room, bool) {
	if !filter.canSet(field, pos, id) {
		return nil, false
	}
	room := &Room{}
	room.field = field
	room.id = id
	room.min = pos
	room.max = pos
	field.grid[pos.Y][pos.X] = id
	room.filter = filter //NewCheckFilter(Point2i{3, 3}, Point2i{1, 1})
	return room, true
}

// NewRoomInt -
func NewRoomInt(field *Field, x, y int, id int, filter *CheckFilter) (*Room, bool) {
	return NewRoom(field, Point2i{x, y}, id, filter)
}

func (o *Room) initParamsForExtend(dir int) (start, offset Point2i, len int) {
	switch dir {
	default:
		panic(fmt.Sprintf("unsupported dir value: %v", dir))
	case 0:
		start = o.min
		start.AddInt(0, -1)
		offset.SetInt(1, 0)
		len = Diff(o.min, o.max).X
	case 2:
		start = o.max
		start.AddInt(0, 1)
		offset.SetInt(-1, 0)
		len = Diff(o.min, o.max).X
	case 1:
		start = o.max
		start.AddInt(1, 0)
		offset.SetInt(0, -1)
		len = Diff(o.min, o.max).Y
	case 3:
		start = o.min
		start.AddInt(-1, 0)
		offset.SetInt(0, 1)
		len = Diff(o.min, o.max).Y
	}
	return start, offset, len + 1
}

// dir 0..3 clockwise from north
func (o *Room) canExtend(dir int) bool {
	pt, offs, len := o.initParamsForExtend(dir)
	for i := 0; i < len; i++ {
		if !o.filter.hasSelf(o.field, pt, o.id) || !o.filter.canSet(o.field, pt, o.id) {
			return false
		}
		pt.Add(offs)
	}
	return true
}

// Extend - dir 0..3 clockwise from north
func (o *Room) Extend(dir int) bool {
	start, offs, len := o.initParamsForExtend(dir)
	pt := start
	grid := o.field.grid
	extended := false
	for i := 0; i < len; i++ {
		if o.filter.canSet(o.field, pt, o.id) {
			grid[pt.Y][pt.X] = o.id
			extended = true
		}
		pt.Add(offs)
	}
	if !extended {
		return false
	}
	o.min.Min(start)
	o.max.Max(start)
	return true
}

// NewCheckFilter -
func NewCheckFilter(size, pivot Point2i) *CheckFilter {
	return &CheckFilter{size: size, pivot: pivot}
}

func (o *CheckFilter) hasSelf(field *Field, pos Point2i, id int) bool {
	grid := field.grid
	ok := false
	if pos.X-1 >= 0 {
		ok = ok || grid[pos.Y][pos.X-1] == id
	}
	if pos.X+1 <= field.size.X {
		ok = ok || grid[pos.Y][pos.X+1] == id
	}
	if pos.Y-1 >= 0 {
		ok = ok || grid[pos.Y-1][pos.X] == id
	}
	if pos.Y+1 <= field.size.Y {
		ok = ok || grid[pos.Y+1][pos.X] == id
	}
	return ok
}

func (o *CheckFilter) canSet(field *Field, pos Point2i, id int) bool {
	pt1, pt2 := pos, pos
	pt1.Sub(o.pivot)
	pt2.Add(o.size).Sub(o.pivot)
	if !pt1.GreaterOrEqual(Point2i{}) {
		return false
	}
	if !pt2.LessOrEqual(field.size) {
		return false
	}
	// ok := false
	for j := pt1.Y; j < pt2.Y; j++ {
		line := field.grid[j]
		for i := pt1.X; i < pt2.X; i++ {
			v := line[i]
			// if v == id {
			// 	ok = true
			// }
			if v > 0 && v != id {
				return false
			}
		}
	}

	// if !ok {
	// 	return false
	// }
	// field.grid[pos.Y][pos.X] = id
	return true
}

func (o *CheckFilter) canSetInt(field *Field, x, y int, id int) bool {
	return o.canSet(field, Point2i{x, y}, id)
}

func (o *CheckFilter) canExtendHorLine(field *Field, pos Point2i, len int, id int) bool {
	return true
}
