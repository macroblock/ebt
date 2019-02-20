package anim

import (
	"time"

	"github.com/hajimehoshi/ebiten"
)

// const ticksPerSecond = 100000

type (
	// List -
	List struct {
		list map[string]Timeline
	}

	// Timeline -
	Timeline struct {
		name         string
		duration     time.Duration
		time         time.Duration
		speed        float64
		initialState State
		states       []State
	}

	// State -
	State struct {
		Time time.Duration
		Tile *ebiten.Image
	}
)

// DeltaTimeFunc -
func DeltaTimeFunc() func() time.Duration {
	lastUpdate := time.Now()
	dt := time.Since(lastUpdate)
	return func() time.Duration {
		dt = time.Since(lastUpdate)
		lastUpdate = time.Now()
		return dt
	}
}

// NewList -
func NewList() *List {
	return &List{list: map[string]Timeline{}}
}

// Add -
func (o *List) Add(timeline Timeline) {
	o.list[timeline.name] = timeline
}

// NewTimeline -
func NewTimeline(name string, dur float64, tile *ebiten.Image, states []State) *Timeline {
	return &Timeline{
		name:         name,
		duration:     time.Duration(dur * float64(time.Second)),
		initialState: State{Tile: tile},
		states:       states,
		speed:        1.0,
	}
}

// SetSpeed -
func (o *Timeline) SetSpeed(speed float64) {
	if speed < 0 {
		return
	}
	o.speed = speed
}

// State -
func (o *Timeline) State(delta time.Duration) State {
	d := time.Duration(float64(delta) * o.speed)
	o.time += d
	if o.time >= o.duration {
		o.time = o.time % o.duration //math.Remainder(o.time, o.duration)
	}
	ret := &o.initialState
	// str := "init"
	for i := range o.states {
		state := &o.states[i]
		if o.time < state.Time {
			break
		}
		// str = strconv.Itoa(i)
		ret = state
	}
	// fmt.Println("state: ", str, "time ", o.time)
	return *ret
}

// NewState -
func NewState(t float64, tile *ebiten.Image) State {
	return State{Time: time.Duration(t * float64(time.Second)), Tile: tile}
}
