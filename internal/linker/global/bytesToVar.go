package linker

type bytesToVarApplication struct {
	cast       func(byte) interface{}
	cummulator func(interface{}, interface{}) interface{}
	start      interface{}
}

var byte4ToVar = bytesToVarApplication{
	cast: func(b byte) interface{} { return uint32(b) },
	cummulator: func(a interface{}, b interface{}) interface{} {
		return (((a).(uint32)) << 1) + (b).(uint32)
	},
	start: uint32(0),
}

var byte8ToVar = bytesToVarApplication{
	cast: func(b byte) interface{} { return uint64(b) },
	cummulator: func(a interface{}, b interface{}) interface{} {
		return (((a).(uint64)) << 1) + (b).(uint64)
	},
	start: uint64(0),
}

var endianMap = map[int]bytesToVarApplication{
	0: {
		cast: func(b byte) interface{} { return nil },
		cummulator: func(a interface{}, b interface{}) interface{} {
			return nil
		},
		start: nil,
	},
	1: {
		cast: func(b byte) interface{} { return uint8(b) },
		cummulator: func(a interface{}, b interface{}) interface{} {
			return b
		},
		start: uint8(0),
	},
	2: {
		cast: func(b byte) interface{} { return uint16(b) },
		cummulator: func(a interface{}, b interface{}) interface{} {
			return (((a).(uint16)) << 1) + (b).(uint16)
		},
		start: uint16(0),
	},
	3: byte4ToVar,
	4: byte4ToVar,
	5: byte8ToVar,
	6: byte8ToVar,
	7: byte8ToVar,
	8: byte8ToVar,
}

func BytesToVarBigEndian(slice []byte) interface{} {
	length := len(slice)
	application := endianMap[length]
	result := application.start
	cast := application.cast
	cummulator := application.cummulator

	for i := 0; i < length; i++ {
		result = cummulator(result, cast(slice[i]))
	}

	return result
}

func BytesToVarLittleEndian(slice []byte) interface{} {
	length := len(slice)
	application := endianMap[length]
	result := application.start
	cast := application.cast
	cummulator := application.cummulator

	for i := length - 1; i >= 0; i-- {
		result = cummulator(result, cast(slice[i]))
	}

	return result
}

func VariableToFlagValues[N uint8 | uint16 | uint32 | uint64](value N, onValid func(N)) {
	var i N = 1
	for ; i != 0; i = (i << 1) {
		if value&i != 0 {
			onValid(i)
		}
	}
}
