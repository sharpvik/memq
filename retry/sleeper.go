package retry

import "time"

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
	time.Sleep(s.BackoffDuration())
}

func (s *Sleeper) BackoffDuration() time.Duration {
	return min(time.Second*time.Duration(s.count*s.multiple), s.cap)
}
