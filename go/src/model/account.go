package model

import (
	"fmt"
	"strconv"
)

type account struct {
	AccountNum int
	balance    float32
	passwd     int
}

func NewAccount(num int) *account {
	if len(strconv.Itoa(num)) < 6 || len(strconv.Itoa(num)) > 10 {
		fmt.Println("你设置的账号长度不正确，账号长度在6-10之间。")
		return &account{
			AccountNum: 000000,
		}
	} else {
		return &account{
			AccountNum: num,
		}
	}
}

func (a *account) SetBalance(bal float32) {
	if bal < 20 {
		fmt.Println("你设置的金额不正确！金额应该大于20元。")
	} else {
		a.balance = bal
		fmt.Println(a.balance)
	}
}

func (a *account) SetPasswd(passwd int) {
	if len(strconv.Itoa(passwd)) != 6 {
		fmt.Println("你设置的密码不正确，密码应该只有6位")
	} else {
		a.passwd = passwd
	}
}

func (a *account) GetPasswd() int {
	return a.passwd
}
func (a *account) GetAccount() int {
	return a.AccountNum
}
func (a *account) GetBalance() float32 {
	return a.balance
}
