package linear

type QueueNode struct {
	Value interface{}
	Next  *QueueNode
}

type Queue struct {
	Head *QueueNode
}
