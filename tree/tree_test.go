package tree

import (
	"fmt"
	"testing"
)

func TestNewTree(t *testing.T) {
	a := &Tree{nil, 10, nil}
	a = insert(a, 8)
	a = insert(a, 9)

	fmt.Println(a.String())

	fmt.Println("End")
}
