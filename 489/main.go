package main

// Robot Room Cleaner
type Robot struct {
}

// Returns true if the cell in front is open and robot moves into the cell.
// Returns false if the cell in front is blocked and robot stays in the current cell.
func (robot *Robot) Move() bool { return true }

// Robot will stay in the same cell after calling TurnLeft/TurnRight.
// Each turn will be 90 degrees.
func (robot *Robot) TurnLeft()  {}
func (robot *Robot) TurnRight() {}

// Clean the current cell.
func (robot *Robot) Clean() {}

var dirs = [][]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}} // 0 - up, 1 - right, 2 - down, 3 - left
func cleanRoom(robot *Robot) {
	visited := make(map[[2]int]bool)
	backtracking(robot, 0, 0, 0, visited)
}

func backtracking(robot *Robot, row, col, dir int, visited map[[2]int]bool) {
	visited[[2]int{row, col}] = true
	robot.Clean()
	for i := 0; i < 4; i++ {
		newDir := (dir + i) % 4
		newRow := row + dirs[newDir][0]
		newCol := col + dirs[newDir][1]
		if !visited[[2]int{newRow, newCol}] && robot.Move() {
			backtracking(robot, newRow, newCol, newDir, visited)
			goback(robot)
		}
		robot.TurnRight()
	}
}

func goback(robot *Robot) {
	robot.TurnLeft()
	robot.TurnLeft()
	robot.Move()
	robot.TurnRight()
	robot.TurnRight()
}
