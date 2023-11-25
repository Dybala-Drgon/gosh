package vm

type FrameType byte

const (
	Blk FrameType = iota // Load constant
	If                   // Add
	For                  // Sub
	Func
)
