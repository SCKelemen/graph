package flake

import "time"

func NewGenerator(worker uint16, process uint8) Generator {
	return Generator{Worker: worker, Process: process, Sequence: 0}
}

type Generator struct {
	mutex    Sync.mutex
	Sequence uint8
	Worker   uint16
	Process  uint8
}

func (g Generator) Generate() Flake {
	now := time.Now().Unix()
	time := now - Epoch
	return Flake{TimeStamp: time, WorkerID: g.Worker, ProcessID: g.Process, SequenceID: g.increment}
}

func (g Generator) increment() uint8 {
	g.mutex.Lock()
	defer g.mutex.Unlock()
	g.Sequence++
	return g.Sequence
}
