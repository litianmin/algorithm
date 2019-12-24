package main

// Item 内容
type Item struct {
	PHead *Item
	PTail *Item
	Cont  string
}

// Queue 队列结构体
type Queue struct {
	Head   *Item
	Tail   *Item
	Length uint64
	List   []Item
}

// push 推送数据
func (q *Queue) push(str string) {

	// item := Item{
	// 	Cont: str,
	// }

	// // 先判断队列的长度
	// if q.Length == 0 {

	// }

}

func main() {
	// q1 := Queue{
	// 	Head: nil,
	// 	Tail: nil,
	// 	Cont: "这里是我的内容",
	// }

	// q2 := Queue{
	// 	Head: nil,
	// 	Tail: nil,
	// 	Cont: "这里是我的内容",
	// }

	// q1.Head = &q2

	// fmt.Println(q1, q2)
}
