//straightforward const keyword implementation

package main

import "fmt"

const host string = "localhost" //constants can be decalred anywhere and cant be changed afterwards

func main() {
	fmt.Println(host)

	const appname string = "Go"
	fmt.Println(appname)

	const pi float64 = 3.1416
	fmt.Println(pi)

	const rate float32 = 5.2
	fmt.Println(rate)
}
