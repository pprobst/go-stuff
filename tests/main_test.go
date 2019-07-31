package main

import (
    "testing"
)

func TestMySum(t *testing.T) {

    type test struct {
        data []int
        answer int
    }

    tests := []test{
        test{[]int{21, 21}, 42},
        test{[]int{16, 0}, 16},
        test{[]int{-5, -50}, -55},
    }

    for _, v := range tests {
        x := mySum(v.data...)
        if x != v.answer {
            t.Error("Expected", v.answer, "Got", x) 
        }
    }
}
