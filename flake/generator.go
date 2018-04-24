package flake

import (
	"sync"
	"time"
)

func NewGenerator(worker uint16, process uint8) Generator {
	return Generator{Worker: worker, Process: process, Sequence: 0}
}

type Generator struct {
	mutex    sync.Mutex
	Sequence uint8
	Worker   uint16
	Process  uint8
}

func (g Generator) Generate() Flake {
	now := time.Now()
	seq := g.increment()
	return Flake{
		timestamp:  now,
		workerID:   g.Worker,
		processID:  g.Process,
		sequenceID: seq}
}

func (g Generator) increment() uint8 {
	g.mutex.Lock()
	defer g.mutex.Unlock()
	// todo: make Sequence a private member, so this can't be modified outside of the interlock.
	g.Sequence++
	return g.Sequence
}
