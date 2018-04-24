package flake

import (
	"testing"
)

func TestFlake(t *testing.T) {
	generator := newGenerator(151, 2)
	id := generator.Generate()
	if id.workerID != 151 {
		t.Errorf("WorkerID was incorrect.\tExpected: %d\tActual: %d.", 151, id.workerID)
	}
	if id.processID != 2 {
		t.Errorf("ProcessID was incorrect.\tExpected: %d\tActual: %d.", 2, id.processID)
	}

	t.Errorf("Working correctly")
}
