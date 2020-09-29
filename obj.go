package main

import "fmt"

type Person struct {
	Name string
	Age  int
	Sex  string
}

func (p Person) Show() string{
	return fmt.Sprintf("my name is %s, Im %d years old", p.Name, p.Age)
}

func main(){
	student := Person{"sunyang", 19, "M"}
	str := student.Show()
	fmt.Println(str)

	teacher := Person{"zhangsan", 30, "M"}
	str2 := teacher.Show()
	fmt.Println(str2)
}