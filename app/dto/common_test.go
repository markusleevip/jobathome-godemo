package dto

import (
	"fmt"
	"testing"
)

func TestGePages(t *testing.T) {

	fmt.Println(GePages(11,10))
	fmt.Println(GePages(111,10))
	fmt.Println(GePages(222,100))

}
