package types

import "fmt"

func NewPortList() *ListPorts {
	return &ListPorts{nil, 0}
}


func increment(list *ListPorts) {
	list.length++
}


func (list *ListPorts) AddElement(port int, isOpened bool) {
	defer increment(list)

	if list.head == nil {
		list.head = &Head{port: port, isOpened: isOpened, next: nil, hasBeenSorted: false}
		return
	}

	head := &list.head
	for (*head).next != nil {
		head = &(*head).next
	}

	(*head).next = &Head{port: port, isOpened: isOpened, next: nil, hasBeenSorted: false}
}


func (list *ListPorts) PrintResult() {
	if list.head == nil {
		return
	}

	head := list.head
	i := 0

	for (*head).next != nil {
		if head.isOpened {
			fmt.Printf("La porta %d è aperta\n", head.port)
		}

		head = (*head).next
		i++
	}
}


func (list *ListPorts) Sort() {
	if list.head == nil {
		return
	}

	head := list.head

	for head != nil {
		head1 := list.head
		for head1 != nil {

		}
	}
}
