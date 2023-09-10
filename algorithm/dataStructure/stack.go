package main

type Item interface{

}
type ItemStack struct {
	items []Item
}

// InitStack 1.栈的初始化
func (s *ItemStack) InitStack() *ItemStack{
	s.items = []Item{}
	return s
}

func (s *ItemStack) Push (node Item)  {
	s.items = append(s.items,node)
}

func (s *ItemStack) Pop() *Item {
	node := s.items[len(s.items)-1]
	s.items = s.items[0:len(s.items)-1]
	return &node
}

func (s *ItemStack) IsEmpty() bool {
	if len(s.items) == 0 {
		return true
	}else {
		return false
	}
}

/*
栈（stack）一种特殊的的串行形式的抽象数据类型，通常一链表形式实现，栈的性质是先进先出。
*/