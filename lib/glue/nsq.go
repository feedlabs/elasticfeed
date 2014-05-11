package glue

import (
	"log"
	"github.com/bitly/go-nsq"
)

type Nsq struct {
	host	string
	port	string
}

func (m Nsq) Connect() {
	log.Printf("%T connected", m)
}

func NewNsq() *Nsq {
	return &Nsq{}
}

func NewNsqWriter() *nsq.Writer {
	return nsq.NewWriter("test")
}
