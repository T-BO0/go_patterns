package main

import (
	"log"

	"github.com/gofrs/uuid"
)

// Interfaces definition ----->
type IPublisher interface {
	AddSubscriber(*ISubscriber)
	RemoveSubscriber(id uuid.UUID)
	Broadcast(message string)
}

type ISubscriber interface {
	GetId() uuid.UUID
	React(message string)
}

// <----- Interfaces definition

// Publisher implementation ----->
type Publisher struct {
	Subscribers map[uuid.UUID]ISubscriber
}

func NewPublisher() *Publisher {
	return &Publisher{
		Subscribers: make(map[uuid.UUID]ISubscriber),
	}
}

func (p *Publisher) AddSubscriber(s ISubscriber) {
	p.Subscribers[s.GetId()] = s
}

func (p *Publisher) RemoveSubscriber(id uuid.UUID) {
	delete(p.Subscribers, id)
}

func (p *Publisher) Broadcast(message string) {
	for _, s := range p.Subscribers {
		s.React(message)
	}
}

// <----- Publisher implementation

// Subscriber implementation ----->

type Subscriber struct {
	id uuid.UUID
}

func NewSubscriber() *Subscriber {
	return &Subscriber{
		id: uuid.Must(uuid.NewV4()),
	}
}

func (s *Subscriber) GetId() uuid.UUID {
	return s.id
}

func (s *Subscriber) React(message string) {
	log.Printf("Message recieved: %s ------ id: %v", message, s.id)
}

// <----- Subscriber implementation
