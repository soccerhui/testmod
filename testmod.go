package testmod

import "fmt"

func SayHello(name, str string) string {
	return fmt.Sprintf("你好哇, %s, %s", name, str)
}
