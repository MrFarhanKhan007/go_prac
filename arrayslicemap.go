package main

import "fmt"

func main2() {
    // Array (fixed size)
    arr := [3]int{1, 2, 3}
    
    // Slice (resizable, very common in Go)
    nums := []int{10, 20, 30}
    nums = append(nums, 40)

    // Map (like HashMap in Java)
    scores := map[string]int{"Alice": 95, "Bob": 88}

    fmt.Println(arr)
    fmt.Println(nums)
    fmt.Println(scores["Alice"])
}
