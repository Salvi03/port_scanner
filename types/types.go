package types

type Head struct {
	port int
	isOpened bool
	hasBeenSorted bool

	next *Head
}

type ListPorts struct {
	head *Head
	length int
}
