package greeting

import "fmt"

func Greet(s string)string{

	message := fmt.Sprintf("Hi, %v. Welcome!", s)
    return message
	//In Go, the := operator is a shortcut for declaring and initializing a variable in one line 
	//(Go uses the value on the right to determine the variable's type). Taking the long way, you might have written this as:
    // var message string
    // message = fmt.Sprintf("Hi, %v. Welcome!", s)
}