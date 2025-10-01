package main

import "fmt"
import "container/ring"

func main () {
	var data *ring.Ring = ring.New(6)

	data.Value = "Eko Kurniawan"

    data = data.Next()
	data.Value = "value 2"

	data = data.Next()
	data.Value = "value 3"

	data = data.Next()
	data.Value = "value 4"

	data = data.Next()
	data.Value = "value 5"

	data = data.Next()
	data.Value = "value 6"

	data.Do(func(p interface{}) {
		fmt.Println(p.(string))
		fmt.Println(data.Value)
	})

	for i := 0; i < data.Len(); i++ {
		fmt.Println(data.Value)
		data = data.Next()
	}
}