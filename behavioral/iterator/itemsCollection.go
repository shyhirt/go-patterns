package iterator

type ItemsCollection struct {
	items []*Item
}

func (i *ItemsCollection) newIterator() Iterator {
	return &ItemIterator{
		items: i.items,
	}
}
