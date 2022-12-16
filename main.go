package main

import (
	"fmt"
	"github.com/vinit-chauhan/mapstr/mapstr"
)

var (
	test1 = "some.random.string"
	test2 = "some.otherKey"
)

func main() {
	m := map[string]interface{}{
		"some": map[string]interface{}{
			"random": map[string]interface{}{
				"string": "Hey there!!!",
			},
			"otherKey": "Not the one",
		},
	}

	mStr := mapstr.New()
	mStr.Put(m)

	key := test2

	val, present, err := mStr.Get(key)
	if err != nil {
		fmt.Errorf("%w", err)
	}
	if present {
		fmt.Println(val)
	} else {
		fmt.Println("does not exist")
	}
}
