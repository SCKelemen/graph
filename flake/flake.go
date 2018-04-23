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

// 64 TimeStamp                               22 WorkerID    10 PID 6 SeqID
// 012345678901234567890123456789012345678901 | 234567890123 | 4567 | 890123

type Flake struct {
	TimeStamp  time.Time
	WorkerID   uint16
	ProcessID  uint8
	SequenceID uint8
}

func (f Flake) Int64() int64 {
	timestamp := f.TimeStamp.Unix() - Epoch
	output := timestamp << f.WorkerID << f.ProcessID << f.SequenceID
	return output.(int64)
}

type gflake int

func GetTimeStamp(id gflake) time.Time {
	timebits := (id >> TimeStampOffset) + Epoch
	stamp := time.Unix(timebits, 0)
	return stamp
}

func GetWorkerId(id gflake) int {
	worker := (id & WorkerIDMask) >> WorkIDOffset
	return worker
}

func GetProcessId(id gflake) int {
	processId := (id & ProcessIDMask) >> ProcessIDOffset
	return processId
}

func GetSequenceId(id gflake) int {
	sequence := id & SequenceIDMask
	return sequence
}
