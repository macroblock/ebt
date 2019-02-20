package interpolator

import (
	"time"
)

// Interpolator -
type Interpolator struct {
	time     time.Duration
	duration time.Duration
	v0, v1   int
	pval     *int
}

// NewInterpolator -
func NewInterpolator(pval *int, dur time.Duration, v0, v1 int) *Interpolator {
	return &Interpolator{
		duration: dur,
		v0:       v0,
		v1:       v1,
		pval:     pval,
	}
}

// Reset -
func (o *Interpolator) Reset(pval *int, dur time.Duration, v0, v1 int) {
	o.time = 0
	o.duration = dur
	o.v0 = v0
	o.v1 = v1
	o.pval = pval
}

// Process -
func (o *Interpolator) Process(delta time.Duration) (time.Duration, bool) {
	o.time += delta
	if o.time >= o.duration {
		return o.time - o.duration, true
	}
	k := float64(o.time) / float64(time.Second)
	// fmt.Printf("k %v\n", k)
	*o.pval = int(float64(o.v1-o.v0)*k) + o.v0
	return 0, false
}
