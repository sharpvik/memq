package service

import (
	"bytes"
	"log"
	"net/http"
	"time"

	"github.com/sharpvik/memq/retry"
)

type Message = []byte

type Service struct {
	queue      chan Message
	subscriber string
}

func New(cap uint, subscriber string) *Service {
	return &Service{
		queue:      make(chan Message, cap),
		subscriber: subscriber,
	}
}

func (s *Service) EnqueueMessage(msg Message) {
	s.queue <- msg
}

func (s *Service) ForwardMessages() {
	log.Println("forwarder started")
	for sleeper := retry.NewSleeper(2, time.Minute); ; sleeper.Sleep() {
		if _, err := s.SendMessage(); err != nil {
			sleeper.Failed()
		}
	}
}

func (s *Service) SendMessage() (*http.Response, error) {
	msg := <-s.queue
	log.Printf("sending: %dB", len(msg))
	return http.Post(
		s.subscriber,
		"application/octet-stream",
		bytes.NewReader(msg),
	)
}
