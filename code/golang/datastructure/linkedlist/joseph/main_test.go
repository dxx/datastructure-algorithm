package joseph

import "testing"

func TestJoseph(t *testing.T) {
    personLinkedList := NewBoyLinkedList(5)
    personLinkedList.ShowPersons()

    personLinkedList.Count(1, 3)
}
