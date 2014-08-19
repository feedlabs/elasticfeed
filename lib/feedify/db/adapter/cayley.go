package adapter

import (
	"log"
)

type Cayley struct {
	host	string
	port	string
}

func (m Cayley) Connect() {
	log.Printf("%T connected", m)
}

func NewCayley() *Cayley {
	return &Cayley{}
}
