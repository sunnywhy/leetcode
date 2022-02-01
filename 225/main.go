package main

// Implement Stack using Queues
type MyStack struct {
    Data []int
}


func Constructor() MyStack {
    return MyStack{Data:[]int{}}
}


func (this *MyStack) Push(x int)  {
    this.Data = append(this.Data, x)
}


func (this *MyStack) Pop() int {
    l := len(this.Data)
    val := this.Data[l-1]
    this.Data = this.Data[:l-1]
    return val
}


func (this *MyStack) Top() int {
    l := len(this.Data)
    return this.Data[l-1]
}


func (this *MyStack) Empty() bool {
    return len(this.Data) == 0
}
