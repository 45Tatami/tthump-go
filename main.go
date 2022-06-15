package main

import "fmt"

func main() {
	fmt.Println("vim-go")

	creq, cres := Create(4)

	creq <- &TthRequest{"/tmp/example.mkv"}
	creq <- nil
	_ = <-cres
}
