package main

// Moving Average from Data Stream

// method 1: use a list,  *list.List
// method 2: circular array

type MovingAverage struct {
	Data  []int
	Len   int
	Total int
	Pos   int
}

func Constructor(size int) MovingAverage {
	data := make([]int, size)
	return MovingAverage{Data: data}
}

func (this *MovingAverage) Next(val int) float64 {
	this.Total -= this.Data[this.Pos]
	this.Total += val
	this.Data[this.Pos] = val
	this.Pos = (this.Pos + 1) % len(this.Data)
	if this.Len < len(this.Data) {
		this.Len++
	}
	return float64(this.Total) / float64(this.Len)
}
