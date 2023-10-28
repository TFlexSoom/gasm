package linker

type ByteQueue interface {
	Length() int
	Append([]byte) ByteQueue
	SeekNext(int) []byte
	SeekAt(int, int) []byte
	Pop(int) []byte
	CanPop(int) bool
}

type bytePageLink struct {
	this []byte
	next *bytePageLink
}

type bytePageLinkedList struct {
	head *bytePageLink
	tail *bytePageLink
}

type byteQueueImpl struct {
	pages         *bytePageLinkedList
	headIndexPage *bytePageLink
	tailIndexPage *bytePageLink
	headIndex     int
	tailIndex     int
	length        int
	capacity      int
	alloc         int
}

const DefaultQueueAllocCap = 1024

func NewByteQueue() ByteQueue {
	return NewByteQueueCap(DefaultQueueAllocCap)
}

func NewByteQueueCap(alloc int) ByteQueue {
	return &byteQueueImpl{
		pages:         nil,
		headIndexPage: nil,
		tailIndexPage: nil,
		headIndex:     0,
		tailIndex:     0,
		length:        0,
		capacity:      0,
		alloc:         alloc,
	}
}

func (bQueue *byteQueueImpl) Length() int {
	return bQueue.length
}

func (bQueue *byteQueueImpl) ensureCapacity(newCap int) {
	if newCap < bQueue.capacity {
		return
	}

	allocAmt := bQueue.alloc
	extraPage := 1
	if newCap%allocAmt == 0 {
		extraPage = 0
	}

	newAllocs := ((newCap - bQueue.capacity) / allocAmt) + extraPage
	head := &bytePageLink{
		this: make([]byte, allocAmt),
		next: nil,
	}

	tail := head

	for i := 1; i < newAllocs; i++ {
		tail.next = &bytePageLink{
			this: make([]byte, allocAmt),
			next: nil,
		}

		tail = tail.next
	}

	if bQueue.pages != nil {
		bQueue.pages.tail.next = head
		bQueue.pages.tail = tail
	} else {
		bQueue.pages = &bytePageLinkedList{
			head: head,
			tail: tail,
		}
		bQueue.headIndexPage = head
		bQueue.tailIndexPage = head
	}

	if bQueue.tailIndexPage == nil {
		bQueue.tailIndexPage = head
	}

	bQueue.capacity += allocAmt * newAllocs
}

func sliceCopy(fromIndex int, fromLimit int, fromSlice []byte, intoIndex int, intoLimit int, intoSlice *[]byte) {
	for i := 0; i < fromLimit && i < intoLimit; i++ {
		(*intoSlice)[i+intoIndex] = fromSlice[fromIndex+i]
	}
}

func (bQueue *byteQueueImpl) Append(bytes []byte) ByteQueue {
	allocAmt := bQueue.alloc
	bQLength := bQueue.length
	bQHead := bQueue.headIndex
	tailIndex := bQueue.tailIndex
	bytesLen := len(bytes)
	bytesCopied := 0

	bQueue.ensureCapacity(bQHead + bQLength + bytesLen)
	lastPage := bQueue.tailIndexPage

	if tailIndex+bytesLen < allocAmt {
		bytesCopied = bytesLen
		sliceCopy(0, bytesLen, bytes, tailIndex, allocAmt-tailIndex, &lastPage.this)
	} else if tailIndex > 0 {
		bytesCopied = allocAmt - tailIndex
		sliceCopy(0, bytesLen, bytes, tailIndex, bytesCopied, &lastPage.this)
		lastPage = lastPage.next
	}

	for ; bytesCopied < bytesLen; bytesCopied += allocAmt {
		sliceCopy(bytesCopied, bytesLen-bytesCopied, bytes, tailIndex, allocAmt, &lastPage.this)
		lastPage = lastPage.next
	}

	bQueue.length += bytesLen
	bQueue.tailIndexPage = lastPage
	bQueue.tailIndex = (tailIndex + bytesLen) % allocAmt
	return bQueue
}

func (bQueue *byteQueueImpl) seekNextPageTuple(n int) ([]byte, *bytePageLink, int) {
	bQHead := bQueue.headIndex
	headIndexPage := bQueue.headIndexPage
	allocAmt := bQueue.alloc

	if bQHead+n < allocAmt {
		return headIndexPage.this[bQHead : bQHead+n], headIndexPage, bQHead + n
	}

	firstPageOffset := allocAmt - bQHead

	buffer := make([]byte, 0, n)
	buffer = append(buffer, headIndexPage.this[bQHead:]...)

	for i := firstPageOffset; i < n-allocAmt; i += allocAmt {
		headIndexPage = headIndexPage.next
		buffer = append(buffer, headIndexPage.this...)
	}

	lastBitLen := n - len(buffer)
	headIndexPage = headIndexPage.next
	if lastBitLen > 0 {
		buffer = append(buffer, headIndexPage.this[:lastBitLen]...)
	}

	return buffer, headIndexPage, lastBitLen
}

func (bQueue *byteQueueImpl) SeekNext(n int) []byte {
	bytes, _, _ := bQueue.seekNextPageTuple(n)
	return bytes
}

func (bQueue *byteQueueImpl) SeekAt(index int, n int) []byte {
	bytes, _, _ := bQueue.seekNextPageTuple(index + n)
	return bytes[index:]
}

func (bQueue *byteQueueImpl) Pop(n int) []byte {
	allocAmt := bQueue.alloc
	bytes, newCurPage, newHeadIndex := bQueue.seekNextPageTuple(n)

	for bQueue.pages.head != newCurPage {
		bQueue.pages.head = bQueue.pages.head.next
		bQueue.capacity -= allocAmt
	}

	bQueue.length -= n
	bQueue.headIndexPage = newCurPage
	bQueue.headIndex = newHeadIndex

	return bytes
}

func (bQueue *byteQueueImpl) CanPop(n int) bool {
	return bQueue.length >= n
}
