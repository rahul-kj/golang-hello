package hello

// Function that computes the sum of two variables
func Addition(a, b int) int {
    return a + b
}

// Multiple returns from a single function
func Compute(a, b int) (int, int, int, int) {
    return a+b, a-b, a*b, a/b
}

// Multiple returns from a single function
func EnhancedCompute(a, b int) (add int, sub int, multi int, div int) {
    add = a + b
    sub = a - b
    multi = a * b
    div = a/b
    return
}
