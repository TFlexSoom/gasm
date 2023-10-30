package linker

import (
	"io"
	"log"
)

type FsmState interface {
	NumBytes() int
	Transform([]byte) (interface{}, error)
	Consume(interface{})
	NextState()
	IsEndState() bool
}

func ReadIntoFSM(reader io.Reader, ctxt FsmState, fsm func(ByteQueue, FsmState) error) error {
	bytesReadAtOnce := 1024
	readingBuffer := make([]byte, 0, bytesReadAtOnce)
	byteQueue := NewByteQueueCap(bytesReadAtOnce)

	for {
		n, err := reader.Read(readingBuffer)
		if err == io.EOF {
			// There is no more data to read
			break
		}
		if err != nil {
			return err
		}

		if n > 0 {
			byteQueue.Append(readingBuffer)
			err = fsm(byteQueue, ctxt)
			if err != nil {
				return err
			}
		} else {
			log.Default().Println("Zero Read!")
		}
	}

	return nil
}

func Fsm(bQueue ByteQueue, fsmState FsmState) error {
	for ; fsmState.IsEndState(); fsmState.NextState() {
		numBytes := fsmState.NumBytes()
		if !bQueue.CanPop(numBytes) {
			break // Load More Bytes
		}

		data, err := fsmState.Transform(bQueue.Pop(numBytes))
		if err != nil {
			return err
		}

		fsmState.Consume(data)
	}

	return nil
}
