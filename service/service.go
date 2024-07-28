package service

import (
	"bytes"
	"log"
	"net/http"
	"time"

	"github.com/sharpvik/memq/queue"
	"github.com/sharpvik/memq/retry"
)

type Message struct {
	contentType string
	content     []byte
}

type Service struct {
	queue      *queue.Queue[Message]
	subscriber string
}

func New(cap uint, subscriber string) *Service {
	return &Service{
		queue:      queue.New[Message](cap),
		subscriber: subscriber,
	}
}

func (s *Service) ForwardMessages() {
	log.Println("forwarder started")

	for sleeper := retry.NewSleeper(2, time.Minute); ; sleeper.Sleep() {
		resp, err := s.SendMessage()
		s.HandleConsumerResponse(resp, err, sleeper)
	}
}

func (s *Service) SendMessage() (*http.Response, error) {
	msg := s.queue.Peek()
	log.Printf("sending: %dB (%s)", len(msg.content), msg.contentType)
	return http.Post(
		s.subscriber,
		msg.contentType,
		bytes.NewReader(msg.content),
	)
}

func (s *Service) HandleConsumerResponse(
	resp *http.Response,
	err error,
	sleeper *retry.Sleeper,
) {
	if err != nil {
		log.Printf("failed to send message: %s", err)
		sleeper.Failed()
	} else if resp.StatusCode != http.StatusOK {
		log.Println("consumer responded: status != 200")
		sleeper.Failed()
	} else {
		s.queue.Flush()
		sleeper.Reset()
	}
}
