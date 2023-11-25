package vm

type FrameType byte

const (
	TGBlk FrameType = iota // Load constant
	TGIf                   // Add
	TGFor                  // Sub
	TGFunc
)
