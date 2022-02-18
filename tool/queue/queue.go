package queue

type (
	//Queue 队列
	Queue struct {
		top    *node
		rear   *node
		length int
	}
	//双向链表节点
	node struct {
		pre   *node
		next  *node
		value interface{}
	}
)

// Create a new queue
func New() *Queue {
	return &Queue{nil, nil, 0}
}

//获取队列长度
func (q *Queue) Len() int {
	return q.length
}

//返回true队列不为空
func (q *Queue) IsEmpty() bool {
	return q.length > 0
}

//返回队列顶端元素
func (q *Queue) Peek() interface{} {
	if q.top == nil {
		return nil
	}
	return q.top.value
}

//入队操作
func (q *Queue) Push(v interface{}) {
	n := &node{nil, nil, v}
	if q.length == 0 {
		q.top = n
		q.rear = q.top
	} else {
		n.pre = q.rear
		q.rear.next = n
		q.rear = n
	}
	q.length++
}

//出队操作
func (q *Queue) Pop() interface{} {
	if q.length == 0 {
		return nil
	}
	n := q.top
	if q.top.next == nil { // 只有一个元素
		q.top = nil
	} else { // 跳到后一个节点来操作前一个节点到指向
		q.top = q.top.next
		q.top.pre.next = nil
		q.top.pre = nil
	}
	q.length--
	return n.value
}
