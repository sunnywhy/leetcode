package main

// Min Stack
type MinStack struct {
    Origin []int
    Min []int
}


func Constructor() MinStack {
    return MinStack{
        Origin: []int{},
        Min: []int{},
    }
}


func (this *MinStack) Push(val int)  {
    this.Origin = append(this.Origin, val)
    l := len(this.Min)
    if l == 0 || this.Min[l-1] >= val {
        this.Min = append(this.Min, val)
    }
}


func (this *MinStack) Pop()  {
    l := len(this.Origin) 
    if l == 0 {
        return
    }
    val := this.Origin[l-1]
    minLen := len(this.Min)
    if val == this.Min[minLen-1] {
        this.Min = this.Min[:minLen-1]
    }
    this.Origin = this.Origin[:l-1]
}


func (this *MinStack) Top() int {
    l := len(this.Origin) 
    return this.Origin[l-1]
}


func (this *MinStack) GetMin() int {
    minLen := len(this.Min)
    return this.Min[minLen-1]
}
