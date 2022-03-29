package main
import(
	"fmt"
	"strconv"
	"strings"
)

func main(){
	str :="hello北京"
// 	//r :=[]rune(str)
// 	for i:=0;i<len(str);i++{
// 		fmt.Printf("%c\n",str[i])
// 	}
	r := []rune(str)
	for i,_ := range r {
		fmt.Printf("%c\n",r[i])
	}
	//字符串转整数
	n,err :=strconv.Atoi("123")
	if err != nil {
		fmt.Println("转换失败",err)
	} else {
		fmt.Println("转换成功",n)
	}
	//整数转字符串
	str2 := strconv.Itoa(123456)
	fmt.Printf("str2=%v,str2=%T\n",str2,str2)
	//字符串转换为byte
	var bytes = []byte("hello")
	fmt.Printf("%v\n",bytes)
	//byte转换为字符串
	str = string([]byte{97,98,99})
	fmt.Printf("%v\n",str)
	//十进制转换为二进制
	a := strconv.FormatInt(123,2)
	fmt.Printf("%v\n",a)
	//十进制转换为八进制
	a = strconv.FormatInt(123,8)
	fmt.Printf("%v\n",a)
	//十进制转换为十六进制
	a = strconv.FormatInt(123,16)
	fmt.Printf("%v\n",a)
	//查找子串是否有指定的字符
	b := strings.Contains("seafood","foo")
	fmt.Println(b)
	//统计一个字符串有几个指定的子串
	num := strings.Count("cheese","e")
	fmt.Printf("num=%v\n",num)
	//不区分大小写字符串比较
	b = strings.EqualFold("abc","ABc")
	fmt.Printf("b=%v\n",b)

	fmt.Println("abc"=="abc")
	
}