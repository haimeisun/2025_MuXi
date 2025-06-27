package main

import "fmt"

func Prime(n int) []int {
    if n < 2 {
        return []int{}
    }

    primes := []int{2}
    for i := 3; i <= n; i += 2 {
        isPrime := true
        for _, p := range primes {
            if p*p > i {
                break
            }
            if i%p == 0 {
                isPrime = false
                break
            }
        }
        if isPrime {
            primes = append(primes, i)
        }
    }
    return primes
}

func main() {
    fmt.Println("===== 质数计算程序 =====")
    fmt.Println("请输入一个整数n：")
    var n int
    fmt.Scanln(&n)
    
    result := Prime(n)
    fmt.Printf("%d以内的质数有：%v\n", n, result)
}