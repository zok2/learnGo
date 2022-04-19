package tool

type LinkNode struct {
	val  int
	Next *LinkNode
}
func getLength( head *LinkNode) int  {
	lenght := 0
	for ;head !=nil;head = head.Next {
		lenght++
	}
	return lenght
}
func addLinkNode(head *LinkNode,val int)  {

}