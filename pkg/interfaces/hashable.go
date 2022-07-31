package interfaces

type Hashable interface {
	Sortable
	Hash() uint64
}
