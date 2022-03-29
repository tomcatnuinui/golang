package model

import(
	"fmt"
)

type persion struct {
	Name string
	age int
	sal float64
}

func Nerpersion(name string) *persion {
	return &persion {
		Name : name,
	}
}

func (p *persion) SetAge(age int) {
	if age > 150 || age <= 0 {
		fmt.Println("你输入的年龄不正确...")
	} else {
		p.age = age
	}
}

func (p *persion) GetAge() {
	fmt.Println(p.age)
}