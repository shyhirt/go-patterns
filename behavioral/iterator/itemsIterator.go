package iterator

type ItemIterator struct {
	index int
	items []*Item
}

func (u *ItemIterator) hasNext() bool {
	if u.index < len(u.items) {
		return true
	}
	return false

}
func (u *ItemIterator) getNext() any {
	if u.hasNext() {
		user := u.items[u.index]
		u.index++
		return user
	}
	return nil
}
