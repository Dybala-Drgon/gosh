package token

type Opcode byte

const (
	OpConstant         Opcode = iota // Load constant
	OpAdd                            // Add
	OpSub                            // Sub
	OpMul                            // Multiply
	OpDiv                            // Divide
	OpRem                            // Remainder
	OpBAnd                           // bitwise AND
	OpBOr                            // bitwise OR
	OpBXor                           // bitwise XOR
	OpBShiftLeft                     // bitwise shift left
	OpBShiftRight                    // bitwise shift right
	OpBAndNot                        // bitwise AND NOT
	OpBComplement                    // bitwise complement
	OpPop                            // Pop
	OpTrue                           // Push true
	OpFalse                          // Push false
	OpEqual                          // Equal ==
	OpNotEqual                       // Not equal !=
	OpGreaterThan                    // Greater than >=
	OpGreaterThanEqual               // Greater than or equal to >=
	OpMinus                          // Minus -
	OpLNot                           // Logical not !
	OpJumpFalsy                      // Jump if falsy
	OpAndJump                        // Logical AND jump
	OpOrJump                         // Logical OR jump
	OpJump                           // Jump
	OpNull                           // Push null
	OpArray                          // Array object
	OpMap                            // Map object
	OpError                          // Error object
	OpImmutable                      // Immutable object
	OpIndex                          // Index operation
	OpSliceIndex                     // Slice operation
	OpCall                           // Call function
	OpReturn                         // Return
	OpReturnValue                    // Return value
	OpGetGlobal                      // Get global variable
	OpSetGlobal                      // Set global variable
	OpSetSelGlobal                   // Set global variable using selectors
	OpGetLocal                       // Get local variable
	OpSetLocal                       // Set local variable
	OpDefineLocal                    // Define local variable
	OpSetSelLocal                    // Set local variable using selectors
	OpGetFree                        // Get free variables
	OpSetFree                        // Set free variables
	OpSetSelFree                     // Set free variables using selectors
	OpGetBuiltin                     // Get builtin function
	OpClosure                        // Push closure
	OpIteratorInit                   // Iterator init
	OpIteratorNext                   // Iterator next
	OpIteratorKey                    // Iterator key
	OpIteratorValue                  // Iterator value
	OpModule                         // Module
	OpSetVar                         // assign value to var
	OpGetVar                         // get value from var
	OpLess
	OpLessEqual
	OPGreater
	OPGreaterEqual
	OpJumpSymTable
	OpExitSymTable
	OpIf
	OpForIfPre
	OpForIfPost
	OpBreak
)

var OpcodeOperands = [...][]int{
	OpConstant:         {2},
	OpPop:              {},
	OpTrue:             {},
	OpFalse:            {},
	OpAdd:              {},
	OpSub:              {},
	OpMul:              {},
	OpDiv:              {},
	OpRem:              {},
	OpBAnd:             {},
	OpBOr:              {},
	OpBXor:             {},
	OpBAndNot:          {},
	OpBShiftLeft:       {},
	OpBShiftRight:      {},
	OpBComplement:      {},
	OpEqual:            {},
	OpNotEqual:         {},
	OpGreaterThan:      {},
	OpGreaterThanEqual: {},
	OpMinus:            {},
	OpLNot:             {},
	OpJumpFalsy:        {2},
	OpAndJump:          {2},
	OpOrJump:           {2},
	OpJump:             {2},
	OpNull:             {},
	OpGetGlobal:        {2},
	OpSetGlobal:        {2},
	OpSetSelGlobal:     {2, 1},
	OpArray:            {2},
	OpMap:              {2},
	OpError:            {},
	OpImmutable:        {},
	OpIndex:            {},
	OpSliceIndex:       {},
	OpCall:             {2},
	OpReturn:           {},
	OpReturnValue:      {1},
	OpGetLocal:         {1},
	OpSetLocal:         {1},
	OpDefineLocal:      {1},
	OpSetSelLocal:      {1, 1},
	OpGetBuiltin:       {1},
	OpClosure:          {2, 1},
	OpGetFree:          {1},
	OpSetFree:          {1},
	OpSetSelFree:       {1, 1},
	OpIteratorInit:     {},
	OpIteratorNext:     {},
	OpIteratorKey:      {},
	OpIteratorValue:    {},
	OpModule:           {2},
	OpSetVar:           {2, 2},
	OpGetVar:           {2, 2},
	OpLessEqual:        {},
	OpLess:             {},
	OPGreaterEqual:     {},
	OPGreater:          {},
	OpJumpSymTable:     {2},
	OpExitSymTable:     {},
	OpIf:               {},
	OpForIfPre:         {},
	OpForIfPost:        {},
	OpBreak:            {},
}
