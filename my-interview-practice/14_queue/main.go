package main

import "errors"

type Queue interface {
	// 出队
	Dequeue() (i int, err error)
	// 入队
	EnQueue(i int)
}

type queue struct {
	slice []int
}

func (q *queue) EnQueue(i int) {
	if q.slice == nil {
		return
	}

	q.slice = append(q.slice, i)
}

func (q *queue) Dequeue() (i int, err error) {
	if len(q.slice) == 0 {
		return 0, errors.New("nil queue")
	}

	// 获取队首
	tmpQ := q.slice[0]
	// 删除队首
	q.slice = q.slice[1:]

	return tmpQ, nil
}
