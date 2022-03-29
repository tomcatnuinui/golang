package main
import(
	"fmt"
	"model"
)

func main() {
	a := model.NewAccount(1)
	a.SetBalance(1)
	a.SetPasswd(1)
	fmt.Printf("账号=%v 密码=%v 余额=%v\n",a.GetAccount(),a.GetPasswd(),a.GetBalance())
}