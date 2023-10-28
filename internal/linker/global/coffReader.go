package linker

import (
	"errors"
	"fmt"
	"io"
	"log"
)

type BinaryFileResultCOFF struct {
	coffFile BinaryFile
}

func (bfrPE BinaryFileResultCOFF) Result() *BinaryFile {
	return &bfrPE.coffFile
}

func ReadIntoFSM(reader io.Reader, ctxt interface{}, fsm func(ByteQueue, interface{}) error) error {
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

const finalState uint = 100

type COFFReadingState struct {
	bFile       BinaryFile
	switchState uint
}

func coffFsmImpl(bQueue ByteQueue, fsmStateUncasted interface{}) error {
	var fsmState *COFFReadingState = (fsmStateUncasted).(*COFFReadingState)

	for fsmState.switchState < 20 {

		switch fsmState.switchState {
		case 0: // ver id
			segSize := 2
			if !bQueue.CanPop(segSize) {
				break // Read another line
			}
			fsmState.bFile.VersionNo = (BytesToVarBigEndian(bQueue.Pop(segSize))).(uint16)
			fsmState.switchState = 1
			break
		case finalState:
			return nil
		default:
			return errors.New(fmt.Sprintf("Unknown State: %v", fsmState.switchState))
		}
	}

	return nil
}

func GetCOFFReader() BinaryFileReader {
	return (func(reader io.Reader) (BinaryFileResult, error) {
		fsmState := COFFReadingState{}

		err := ReadIntoFSM(reader, (&fsmState), coffFsmImpl)
		if err != nil {
			return nil, err
		}

		if fsmState.switchState != 20 {
			return nil, errors.New(fmt.Sprintf("Not Enough Bytes To Fulfill COFF File! State: %v", fsmState.switchState))
		}

		result := BinaryFileResultCOFF{
			coffFile: fsmState.bFile,
		}

		return result, nil
	})
}
