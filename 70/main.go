package main

// Climbing Stairs
var memo = make(map[int]int)
func climbStairs(n int) (r int) {
    defer func() {
        memo[n] = r
    }() 
    if n == 0 || n == 1 {
        return 1
    }
    if val, ok := memo[n]; ok {
        return val
    }
    return climbStairs(n-1) + climbStairs(n-2)
}