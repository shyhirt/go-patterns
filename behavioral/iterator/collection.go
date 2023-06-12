package iterator

type Collection interface {
	newIterator() Iterator
}
