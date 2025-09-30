package main

import (
	"fmt"
	/*
	   for including the local module we have to first do
	*/)

// to download this module we have to run go mod tidy

func swap(a *int ,b *int){
	var temp int = *a
	*a = *b
	*b = temp
}

func main() {
	// fmt.Println("Hello, World!")
    // fmt.Println(quote.Go())

    // fmt.Println(greeting.Greet("John"))

// 	fruits := [3]string{"apple", "orange", "banana"}
//   for _, val := range fruits {
//      fmt.Printf("%v\n", val)
//   }
    var a int = -20
	var b int = 90
	// var temp int = a
    // a = b
	// b= temp
    swap(&a,&b)
    fmt.Println(a,b)

	var mp = map[string]int{"apple":5,"banana":3}
    mp["apple"]++;
	fmt.Println(mp["apple"]);
}