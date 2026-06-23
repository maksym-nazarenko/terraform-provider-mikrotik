package main

import (
	"fmt"
	"strconv"
	"testing"
)

func TestDebug(t *testing.T) {

	items := []string{"one", "two", "three"}

	for i := 0; ; i++ {
		if i < 4 {
			items = append(items, strconv.Itoa(i))
		}
		item := items[0]
		fmt.Printf("processing %s\n", item)
		items = items[1:]
	}

	t.Fail()
}
