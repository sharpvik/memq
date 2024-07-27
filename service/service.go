package service

import (
	"sync"

	"github.com/zyedidia/generic/queue"
)

type Message = []byte

type Service struct {
	mutex      sync.RWMutex
	queue      *queue.Queue[Message]
	subscriber string
}

func New() *Service {
	return &Service{
		queue: queue.New[Message](),
	}
}

func (s *Service) Enqueue(msg Message) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.queue.Enqueue(msg)
}

func (s *Service) SetSubscriber(subscriber string) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.subscriber = subscriber
}
