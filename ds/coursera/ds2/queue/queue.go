package queue

import (
	"errors"
)

var (
	ErrEmpty = errors.New("empty")
)

type Queue struct {
	data []interface{}
}

func (q *Queue) Push(x interface{}) {
	q.data = append(q.data, x)
}

func (q *Queue) Pop() (interface{}, error) {
	if len(q.data) == 0 {
		return nil, ErrEmpty
	}

	d := q.data[0]
	q.data = append(q.data[1:])
	return d, nil
}

func (q *Queue) Len() int {
	return len(q.data)
}

func (q *Queue) IsEmpty() bool {
	return q.Len() == 0
}
