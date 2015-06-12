package main

import "os"
import "fmt"
import "mpris2"

func main() {
	if len(os.Args) != 2 {
		fmt.Printf(os.Args[0] + " [play | pause | playpause]")
	} else {
		switch os.Args[1] {
		case "playpause":

		default:
			fmt.Printf("%s is currently unsupported\n", os.Args[1])
		}
	}
}
