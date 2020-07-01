package latency

import (
	"time"
)

type Latency int64

func New() Latency {
	return Latency(time.Now().UnixNano())
}

func (l Latency) Duration() time.Duration {
	return time.Duration(time.Now().UnixNano() - int64(l))
}

func (l Latency) Seconds() float64 {
	return l.Duration().Seconds()
}
