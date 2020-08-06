package main

import (
    "math/big"
)

// func main() {
//     fmt.Println("the Sum of First 100 even Fibonacci number is ::",getEvenFibonacciSum(100).String())
// }

func getEvenFibonacciSum(length int) *big.Int {
    //even fibonacci numbers can be obtained by F(n) = 4*f(n-1)+f(n-2), starting from seed 0,2
    ef1 := big.NewInt(0)
    if length == 0 {
        return ef1
    }
    ef2 := big.NewInt(2)
    multiplier := big.NewInt(4)
    sum := big.NewInt(2)
    for i := 1; i < length; i++ {
        ef3 := big.NewInt(0).Add((big.NewInt(0).Mul(multiplier,ef2)),ef1)
        sum = big.NewInt(0).Add(sum,ef3)
        ef1 = ef2
        ef2=ef3

    }
    
    return sum
}