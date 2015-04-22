package pipeline

import (
	"time"
	"math/rand"
)

func RandomAnimator(data interface {}) interface {} {

	// PIPE DELAY SIMULATION

	amt := time.Duration(rand.Intn(200))
	time.Sleep(amt * time.Millisecond)

	return data
}

func init() {}
