package testmod

import "fmt"

func SayHello(name, str string) string {
	return fmt.Sprintf("hello, %s, %s", name, str)
}
