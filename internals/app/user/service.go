package user

import (
	"sync"
	"time"
)

type service struct {
	mu sync.Mutex
}

type Service interface {
	MethodOne(waitTime int)
	MethodTwo(waitTime int)
}

func NewService() Service {
	return &service{}
}

func (s *service) MethodOne(waitTime int) {
	s.mu.Lock()
	defer s.mu.Unlock()
	time.Sleep(time.Duration(waitTime) * time.Second)
}

func (s *service) MethodTwo(waitTime int) {
	time.Sleep(time.Duration(waitTime) * time.Second)
}
