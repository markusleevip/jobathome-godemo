package dto

import (
	"fmt"
	"testing"
)

func TestGePages(t *testing.T) {

	fmt.Println(GetPages(11, 10))
	fmt.Println(GetPages(111, 10))
	fmt.Println(GetPages(222, 100))

}
