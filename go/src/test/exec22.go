package main

import(
	"fmt"
)

type Student struct {
	Name string
	Age int
	scor int
}

func (stu *Student) testing() {
	fmt.Println("考试正在进行中...")
}

func (stu *Student) Setscor(s int) {
	if s > 0 && s < 101 {
		stu.scor = s
	} else {
		fmt.Println("输入的成绩不正确")
	}
}

func (stu *Student) sum(n1 int,n2 int) int{
	return n1 + n2 
}

func (stu *Student) ShowInfo(){
	fmt.Printf("学生的名字=%v 年龄=%v 成绩=%v\n",stu.Name,stu.Age,stu.scor)
}

type Min struct {
	Student
}

func (min *Min) Saychina() {
	fmt.Println("小学生读语文",min.Name)
}

type Max struct {
	Student
}

func (max *Max) SayEinglish() {
	fmt.Println("大学生读英语",max.Name)
}



func main() {

	var x Min
	x.Name = "tom"
	x.Age = 8
	x.testing()
	x.Setscor(80)
	fmt.Println(x.sum(1,2))
	x.ShowInfo()
	x.Saychina()

	fmt.Println()

	var d Max
	d.Name = "skr"
	d.Age = 20
	d.testing()
	d.Setscor(100)
	fmt.Println(d.sum(10,20))
	d.SayEinglish()
	d.ShowInfo()
}