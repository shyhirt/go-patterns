package iterator

import (
	"fmt"
	"testing"
)

func TestItemIterator(t *testing.T) {
	itemsCollection := &ItemsCollection{
		items: []*Item{
			{
				Id:   0,
				Data: "One",
			},
			{
				Id:   1,
				Data: "Two",
			},
		},
	}

	iterator := itemsCollection.newIterator()
	for iterator.hasNext() {
		item := iterator.getNext()
		fmt.Printf("Item is %+v\n", item)
	}
}
