package latency

import (
	"testing"
	"time"
)

func TestLatency(t *testing.T) {

	l := New()

	time.Sleep(time.Millisecond * 10)

	if secs := int(l.Seconds() * 1000); secs < 10 {
		t.Fatal("Seconds failed")
	}

}
