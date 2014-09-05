package adapter

import (
	"log"
)

type CayleyAdapter struct {
	host string
	port string
}

func (m CayleyAdapter) Connect() {
	log.Printf("%T connected", m)
}

func NewCayleyAdapter() *CayleyAdapter {
	return &CayleyAdapter{}
}
