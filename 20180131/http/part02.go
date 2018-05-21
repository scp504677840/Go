package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	//fmt.Println(math.Trunc())
	//r := rand.New(rand.NewSource(time.Now().UnixNano()))
	//fmt.Println(math.Trunc(r.Float64() * 1000000))
	//fmt.Println(math.Trunc(r.Float64() * 1000000))
	//fmt.Println(math.Trunc(r.Float64() * 1000000))
	//fmt.Println(math.Trunc(r.Float64() * 1000000))
	//fmt.Println(rand.ExpFloat64())

	stu := Student{"Tom",20}

	//names := map[string]int{}
	//names["Tom"] = 20

	bytes, err := json.Marshal(&stu)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(bytes))

}

type Student struct {
	name string
	age  int
}
