package main

import "fmt"

/*
[go fmt] -- Used to format code via command line
[go fmt ./..] -- To format all the '.go' files in the directory
[go run file.go] -- To run the go file without creating binary file(executable)
[go build] -- To create an executable file
*/
func main() {
	fmt.Println("Usual convention to have this program while learning new lang")
	fmt.Println("Hello World :-) ")
	// fmt.Println returns to values (integer and error)
	n, err := fmt.Println("Hello GoLang")
	fmt.Println(n)
	fmt.Println(err)

	// to void(ignore) the error value put "_" in the second place
	m, _ := fmt.Println("Hello GoLang")
	fmt.Println(m)
}
