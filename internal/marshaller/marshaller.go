package marshaller

import (
	"errors"
	"fmt"
)

const zeroByte = byte(0)

var hexTranslator = map[rune]byte{
	'F': 15,
	'E': 14,
	'D': 13,
	'C': 12,
	'B': 11,
	'A': 10,
	'9': 9,
	'8': 8,
	'7': 7,
	'6': 6,
	'5': 5,
	'4': 4,
	'3': 3,
	'2': 2,
	'1': 1,
	'0': 0,
}

func MarshallHexByte(value string) error {
	res := make([]byte, ((len(value)-2)/2)+1)
	append := func(idx int, val byte) {
		res[idx] = (res[idx] << 4) + val
	}

	for i, r := range value[2:] {
		val := hexTranslator[r]
		if r != '0' && val == zeroByte {
			return errors.New(fmt.Sprintf("Unknown Hex Char %v", r))
		}
		append(i/2, val)
	}

	return nil
}
