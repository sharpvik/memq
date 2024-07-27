package retry

import (
	"log"
	"math"
	"time"
)

type Sleeper struct {
	count    uint
	multiple uint
	cap      time.Duration
}

func NewSleeper(multiple uint, cap time.Duration) *Sleeper {
	return &Sleeper{
		multiple: multiple,
		cap:      cap,
	}
}

func (s *Sleeper) Failed() {
	s.count++
}

func (s *Sleeper) Reset() {
	s.count = 0
}

func (s *Sleeper) Sleep() {
	if backoff := s.BackoffDuration(); backoff != 0 {
		log.Printf("backoff: %v", backoff)
		time.Sleep(backoff)
	}
}

// Geometric progression formula: a_n = ar^(n-1)
// a = 1
// n = count
// r = multiple
func (s *Sleeper) BackoffDuration() time.Duration {
	a := 1.0
	n := float64(s.count)
	r := float64(s.multiple)
	an := uint(a * math.Pow(r, n-1))
	return min(time.Second*time.Duration(an), s.cap)
}
