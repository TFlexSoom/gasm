package linker

import (
	"testing"
)

func take(val byte, n int) []byte {
	result := make([]byte, 0, n)

	for i := 0; i < n; i++ {
		result = append(result, val)
	}

	return result
}

func sum(slice []byte) int {
	result := 0
	length := len(slice)
	for i := 0; i < length; i++ {
		result += int(slice[i])
	}

	return result
}

func TestByteQueueAppend(t *testing.T) {
	windowAlloc := 5
	ones := take(0x01, windowAlloc)

	bQueue := NewByteQueueCap(windowAlloc)
	bQueue.Append(ones)

	if bQueue.Length() != len(ones) {
		t.Error("Length of ByteQueue is incorrect!")
	}
}

func TestByteQueueSeek(t *testing.T) {
	windowAlloc := 5
	ones := take(0x01, windowAlloc)

	bQueue := NewByteQueueCap(windowAlloc)
	bQueue.Append(ones)

	for i := 0; i < bQueue.Length(); i++ {
		if bQueue.SeekAt(i, 1)[0] != 0x01 {
			t.Error("Seek grabbed bad data!")
		}
	}
}

func TestByteQueuePop(t *testing.T) {
	windowAlloc := 5
	ones := take(0x01, windowAlloc)

	bQueue := NewByteQueueCap(windowAlloc)
	bQueue.Append(ones)

	for i := 0; i < len(ones); i++ {
		if bQueue.Pop(1)[0] != 0x01 {
			t.Error("bQueue Popped bad data!")
		}
	}

	if bQueue.Length() != 0 {
		t.Error("Length of ByteQueue is incorrect!")
	}
}

func TestByteQueueCapacity(t *testing.T) {
	windowAlloc := 5
	ones := take(0x01, windowAlloc)

	bQueue := NewByteQueueCap(windowAlloc)

	if (bQueue).(*byteQueueImpl).capacity != 0 {
		t.Errorf("Capacity invalid. Capacity = %v | Should be %v",
			(bQueue).(*byteQueueImpl).capacity,
			0,
		)
	}

	bQueue.Append(ones)

	if (bQueue).(*byteQueueImpl).capacity != len(ones) {
		t.Errorf("Capacity invalid. Capacity = %v | Should be %v",
			(bQueue).(*byteQueueImpl).capacity,
			windowAlloc,
		)
	}

	bQueue.Pop(len(ones))

	if (bQueue).(*byteQueueImpl).capacity != 0 {
		t.Errorf("Capacity invalid. Capacity = %v | Should be %v",
			(bQueue).(*byteQueueImpl).capacity,
			0,
		)
	}

}

func TestByteQueueLessThanPageAppend(t *testing.T) {
	windowAlloc := 5
	ones := take(0x01, windowAlloc)
	twos := take(0x02, windowAlloc)

	bQueue := NewByteQueueCap(windowAlloc)
	bQueue.Append(ones)
	bQueue.Append(twos)
	shouldBeLength := len(twos) + len(ones)
	shouldBeSum := len(twos)*2 + len(ones)
	for i := 0; i < windowAlloc; i++ {
		seekSum := sum(bQueue.SeekNext(shouldBeLength))
		if seekSum != shouldBeSum {
			t.Errorf("Incorrect Sum! Should Be: %v Is: %v", shouldBeSum, seekSum)
		}

		elem := bQueue.Pop(1)
		if elem[0] != 0x01 || len(elem) > 1 {
			t.Errorf("Incorrect elem(s) popped! %v", elem)
		}

		seekSum = sum(bQueue.SeekNext(shouldBeLength - 1))
		if seekSum != (shouldBeSum - 1) {
			t.Errorf("Incorrect Sum! Should Be: %v Is: %v", shouldBeSum, seekSum)
		}

		bQueue.Append([]byte{0x01})
	}
}
