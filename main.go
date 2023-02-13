//demonstration of wait groups
// Filename: main.go
package main

import (
	"fmt"
	"sync"
)

func one(WG *sync.WaitGroup) {
	defer WG.Done()
	fmt.Println("guten tag")
}
func two(WG *sync.WaitGroup) {
	defer WG.Done()
	fmt.Println("ni hao")
}
func three(WG *sync.WaitGroup) {
	defer WG.Done()
	fmt.Println("hola")
}

func main() {
	var WG sync.WaitGroup
	WG.Add(3)

	go one(&WG)
	go two(&WG)
	go three(&WG)

	//LETS DELAY
	WG.Wait()

}
