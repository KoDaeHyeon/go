package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	seed := time.Now().Unix()
	rand.Seed(seed)

	for i := 1; i < 6; i++ {
		dice := rand.Intn(6) + 1
		fmt.Println(dice)
	}
}
