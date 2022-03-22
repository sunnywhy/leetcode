package main

import "math"

// Design Tic-Tac-Toe

type TicTacToe struct {
	Rows         []int
	Cols         []int
	Diagonal     int
	AntiDiagonal int
	Size         int
}

func Constructor(n int) TicTacToe {
	rows := make([]int, n)
	cols := make([]int, n)
	return TicTacToe{rows, cols, 0, 0, n}
}

func (this *TicTacToe) Move(row int, col int, player int) int {
	val := 1
	if player != 1 {
		val = -1
	}
	this.Rows[row] += val
	this.Cols[col] += val
	if row == col {
		this.Diagonal += val
	}
	if row+col == this.Size-1 {
		this.AntiDiagonal += val
	}
	if math.Abs(float64(this.Rows[row])) == float64(this.Size) ||
		math.Abs(float64(this.Cols[col])) == float64(this.Size) ||
		math.Abs(float64(this.Diagonal)) == float64(this.Size) ||
		math.Abs(float64(this.AntiDiagonal)) == float64(this.Size) {
		return player
	}
	return 0
}
