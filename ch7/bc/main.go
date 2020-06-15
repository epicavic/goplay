package main

import "fmt"

// Bc exported type (bytecounter)
type Bc int

func main() {
	var c Bc
	c.Write([]byte("hello")) // c+=5 | direct call of Write method
	fmt.Println(c)           // 5
	fmt.Fprintf(&c, "world") // c+=5 | implicit call of Write method
	fmt.Println(c)           // 10
}

func (c *Bc) Write(s []byte) (int, error) {
	*c += Bc(len(s))
	return len(s), nil // required by io.Writer interface type "want Write([]byte) (int, error)"
}
