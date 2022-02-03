package main

// Pow(x, n)
func myPow(x float64, n int) float64 {
    if n == 0 {
        return 1
    }
    if n == 1 {
        return x
    }
    if n < 0 {
        return 1/myPow(x, -n)
    }
    if n % 2 == 0 {
        val := myPow(x, n/2)
        return val * val
    }
    return x * myPow(x, n-1)
}