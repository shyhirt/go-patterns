package iterator

type Iterator interface {
	hasNext() bool
	getNext() any
}
