package interfaces

type Sortable interface {
	Equtable
	LessThan(b Sortable) bool
}
