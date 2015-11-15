package fib

func FibFunc() func() uint64 {
    var a, b uint64 = 0, 1
    
    return func() uint64 {
        a, b = b, a + b
        return a
    }
}

