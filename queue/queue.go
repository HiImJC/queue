package queue

import (
	"container/list"
	"fmt"
)

type Queue struct {
	*list.List
	existingEntrys map[interface{}]*list.Element
}

func New() Queue {
	return Queue{
		list.New(),
		make(map[interface{}]*list.Element, 0),
		}
}

func (q Queue) Push(v interface{}) {
		elem, ok := q.existingEntrys[v]
		if !ok {
			q.existingEntrys[v] = q.PushBack(v)

			return
		}

		q.MoveToBack(elem)
}

func (q Queue) Pop() interface{} {
	e := q.Front()
	if e != nil {
		v := q.Remove(e)
		delete(q.existingEntrys, v)

		return v
	}

	return nil
}

func (q Queue) Print() {
	for e := q.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}