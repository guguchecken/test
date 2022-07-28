package main

import (
	"fmt"
)

func main() {

	// for i := 0; i < len(os.Args); i++ {
	// 	fmt.Printf("os.Args[%d]: %v\n", i, os.Args[i])
	// }
	// if _, ok := os.LookupEnv(`GOPATH`); ok {
	// 	fmt.Printf("os.Getenv(`GOPATH`): %v\n", os.Getenv(`GOPATH`))
	// } else {
	// 	fmt.Println(`error`)
	// }
	// p := git.NewPR(`matrixorigin`, `matrixone`, `4174`, `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`)
	// p.GetPaths()
	i := 10
	for ; i >= 0; i-- {
		fmt.Printf("i: %v\n", i)
	}
	fmt.Printf("i: %v\n", i)
}















