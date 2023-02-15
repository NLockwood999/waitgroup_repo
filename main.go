//demonstration of wait groups
// Filename: main.go
package main

import (
	"flag"
	"fmt"
)

/*
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

	//application of waitgroups
	WG.Wait()

}
*/
func main() {
	msg := flag.String("msg", "Howdy, Stranger", "the message to display")
	num := flag.Int("num", 1, "how many time to print the message")
	caps := flag.Bool("caps", false, "should the string be all caps")
	flag.Parse()
	for i := 0; i < *num; i++ {
		fmt.Println(*msg)
	}
	//check if user set the flag
	fmt.Println(*msg)
	fmt.Println(*caps)
}

//gggmn
