package flake

import "time"

const (
	TimeStampLength = 42 // a bunch
	WorkerIDLength  = 12 // 4096
	ProcessIDLength = 4  // 16
	SequenceLength  = 6  // 64

	Epoch = 1514764800 // 01/01/2018 @ 12:00am (UTC)

	WorkerIDMask   = 0
	ProcessIDMask  = 0
	SequenceIDMask = 0x3F

	TimeStampOffset = 22
	WorkIDOffset    = 10
	ProcessIDOffset = 6
)

// Flake is a unique ID that takes the form of a UINT64. The Flake
// is composed of 4 parts. A timestamp, a workID, a processID, and
// a sequenceID. The upper 42 bits comprise the timestamp.
// The timestamp is generated by subtracting the Flake Epoch from
// the Unix timestamp. The next 12 bits identify a worker. In most
// cases this is a machine ID, but if multiple instances run on the
// same machine, they could be different. The next 4 bits represent
// the process. The process is essentially a thread ID. The next 6
// bits are a sequential integer that is atomically incremented.
// 64 TimeStamp                               22 WorkerID    10 PID 6 SeqID
// 012345678901234567890123456789012345678901 | 234567890123 | 4567 | 890123
type Flake struct {
	timestamp  time.Time
	workerID   uint16
	processID  uint8
	sequenceID uint8
}

// IFlake is an interface for the flake
type IFlake interface {
	TimeStamp() time.Time
	WorkerID() uint16
	ProcessID() uint8
	SequenceID() uint8
}

// TimeStamp returns the flake's timestamp
func (f Flake) TimeStamp() time.Time {
	return f.timestamp
}

// WorkerID returns the flake's workerid
func (f Flake) WorkerID() uint16 {
	return f.workerID
}

// ProcessID returns the flake's processid
func (f Flake) ProcessID() uint8 {
	return f.processID
}

// SequenceID returns the flake's sequenceid
func (f Flake) SequenceID() uint8 {
	return f.sequenceID
}

// Int64 returns the Flake as an Int64
func (f Flake) Int64() int64 {
	timestamp := f.timestamp.Unix() - Epoch
	// todo: actually make this math correct
	output := timestamp << f.workerID << f.processID << f.sequenceID
	return output
}

//
// this schema probably won't work on Javascript because it mishandles
// integers greater than   54 bits
type gflake int64

// GetTimeStamp takes a uint64  Flake ID
// and decodes the timestamp from it
func GetTimeStamp(id gflake) time.Time {
	timebits := int64((id >> TimeStampOffset) + Epoch)
	stamp := time.Unix(timebits, 0)
	return stamp
}

// GetWorkerID decodes the workerID from
// the flake and returns it as a uint16
func GetWorkerID(id gflake) uint16 {
	worker := (id & WorkerIDMask) >> WorkIDOffset
	return uint16(worker)
}

// GetProcessID decodes the ProcessID from
// the flake and returns it as a uint8
func GetProcessID(id gflake) uint8 {
	process := (id & ProcessIDMask) >> ProcessIDOffset
	return uint8(process)
}

// GetSequenceID decodes the SequenceID from
// the flake and returns it as a uint8
func GetSequenceID(id gflake) uint8 {
	sequence := id & SequenceIDMask
	return uint8(sequence)
}
