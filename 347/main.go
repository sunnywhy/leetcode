package main

// Top K Frequent Elements

type Data struct {
	Val       int
	Frequency int
}

func topKFrequent(nums []int, k int) []int {
	m := make(map[int]int)
	for _, v := range nums {
		m[v] = m[v] + 1
	}
	var list []Data
	for k, v := range m {
		list = append(list, Data{Val: k, Frequency: v})
	}

	quickSelect(&list, 0, len(list)-1, k)

	var res []int
	for i := 0; i < k; i++ {
		res = append(res, list[i].Val)
	}
	return res
}

func quickSelect(list *[]Data, start, end, k int) {
	if start > end {
		return
	}
	pos := partition(list, start, end)
	if pos == k-1 {
		return
	}
	if pos > k-1 {
		quickSelect(list, start, pos-1, k)
	} else {
		quickSelect(list, pos+1, end, k)
	}
}
func partition(list *[]Data, start, end int) int {
	val := (*list)[end].Frequency
	pos := start
	for i := start; i <= end; i++ {
		if (*list)[i].Frequency > val {
			(*list)[pos], (*list)[i] = (*list)[i], (*list)[pos]
			pos++
		}
	}
	(*list)[pos], (*list)[end] = (*list)[end], (*list)[pos]
	return pos
}
