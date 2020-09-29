package main

import (
	"encoding/json"
	"fmt"
	"goweb/cache"
	_ "os"
)

type ColorGroup struct {
	Id     int
	Name   string
	Colors []string
}

func main() {
	cache.Init()
	groups := []ColorGroup{
		{1, "php", []string{"red", "blue"}},
		{2, "redis", []string{"red", "blue"}},
	}

	//json 序列化
	b, _ := json.Marshal(groups)
	str := string(b)
	fmt.Println(str)

	//反序列化
	var ungroups []ColorGroup
	err := json.Unmarshal([]byte(str), &ungroups)
	if err != nil {
		panic(err)
	}
	fmt.Println(ungroups)

    //遍历
	for k, v := range ungroups {
		str1 := fmt.Sprintf("key = %d, v = %+v, id = %d, name = %s, color = %v", k, v, v.Id, v.Name, v.Colors)
		fmt.Println(str1)
	}

	//闭包
	next := getSequence()
	fmt.Println(next())
	fmt.Println(next())
}

func getSequence() func() int{
	var num int = 0
	return func() int {
		num = num + 1
		return num
	}
}