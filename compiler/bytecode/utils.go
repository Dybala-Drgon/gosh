package bytecode

import "gosh/compiler/token"

func separateBytes(number uint16) (highByte, lowByte uint8) {
	highByte = uint8(number >> 8)  // 右移8位获取高8位
	lowByte = uint8(number & 0xFF) // 使用位与运算获取低8位
	return highByte, lowByte
}

func combineBytes(highByte, lowByte uint8) uint16 {
	combined := (uint16(highByte) << 8) | uint16(lowByte)
	return combined
}

func MakeInstruction(opcode token.Opcode, operands ...int) []byte {
	numOperands := token.OpcodeOperands[opcode]

	totalLen := 1
	for _, w := range numOperands {
		totalLen += w
	}

	instruction := make([]byte, totalLen, totalLen)
	instruction[0] = byte(opcode)
	offset := 1
	for i, o := range operands {
		width := numOperands[i]
		switch width {
		case 1:
			instruction[offset] = byte(o)
		case 2:
			n := uint16(o)
			instruction[offset] = byte(n >> 8)
			instruction[offset+1] = byte(n)
		}
		offset += width
	}

	return instruction
}
