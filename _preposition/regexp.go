package main

import (
	"fmt"
	"regexp"
)

func main() {

	var (
		reg1 *regexp.Regexp // 用于接收正则表达式对象
		reg2 *regexp.Regexp
		//err  error // 用于接收error
	)

	/*
			1.关于写正则表达式
		    ``：反引号可以避免转义，安心写正则表达式
	*/
	const (
		regExpression1 string = `[0-9a-zA-Z]+@[0-9a-zA-Z]+\.[0-9a-zA-Z]+`
		regExpression2 string = `([0-9a-zA-Z]+)@([0-9a-zA-Z.]+)\.([0-9a-zA-Z]+)`
		//regExpression2 string = `([a-zA-Z0-9]+)@([a-zA-Z0-9.]+)\.([a-zA-Z0-9]+)`
		srcString string = `xxx abc@gmail.com yyy 123@163.com zzz`
		//srcString string = "My email is ccmouse@gmail.com.cn@dffsfs  aaa@163.com"
	)

	/*
		2.关于创建正则表达式对象
		有regexp.Compile和regexp.MustCompile两个方法
	*/
	reg1 = aboutCreateRegexpByCompile(regExpression1)
	reg2 = aboutCreateRegexpByMustCompile(regExpression2)

	/*
	   3.关于由正则表达式对象查找源字符串中符合条件的子串
	*/

	// 测试FindAll函数
	aboutFindAll(srcString, reg1)

	// 测试FindAllString函数
	aboutFindAllString(srcString, reg1)

	// 测试Find函数
	aboutFind(srcString, reg1)

	// 测试FindString函数
	aboutFindString(srcString, reg1)

	/*
	   4.关于由正则表达式对象提取匹配子串
	*/

	// 测试FindAllSubmatch函数
	aboutFindAllSubmatch(srcString, reg2)

	// 测试FindAllStringSubmatch函数
	aboutFindAllStringSubmatch(srcString, reg2)

	// 测试FindSubmatch函数
	aboutFindSubmatch(srcString, reg2)

	// 测试FindStringSubmatch函数
	aboutFindStringSubmatch(srcString, reg2)

}

func aboutCreateRegexpByCompile(regExpression string) *regexp.Regexp {

	// regexp.Compile，返回正则表达式对象和err，如果正则表达式有误，则err不为nil

	var (
		reg *regexp.Regexp
		err error
	)

	if reg, err = regexp.Compile(regExpression); err != nil {
		panic("regExpression1错误：")
	}

	return reg
}

func aboutCreateRegexpByMustCompile(regExpression string) *regexp.Regexp {

	// regexp.MustCompile，返回正则表达式对象，如果正则表达式有误，则会panic

	return regexp.MustCompile(regExpression)
}

func aboutFindAll(srcString string, reg *regexp.Regexp) {

	// FindAll返回所有的匹配项
	// 接收源字符串的[]byte形式，返回[][]byte
	// 第二个参数n表示要拿多少个匹配项
	// n=-1时拿全部；n>匹配总个数时拿全部

	fmt.Println("\n-----------FindAll----------")

	fmt.Println("FindAll, -1")

	ret1 := reg.FindAll([]byte(srcString), -1)
	for _, item := range ret1 {
		fmt.Printf("%s ", string(item))
	}

	fmt.Println("\nFindAll, 1")

	ret2 := reg.FindAll([]byte(srcString), 1)
	for _, item := range ret2 {
		fmt.Printf("%s ", string(item))
	}

	fmt.Println("\nFindAll, 666")

	ret3 := reg.FindAll([]byte(srcString), 666)
	for _, item := range ret3 {
		fmt.Printf("%s ", string(item))
	}
}

func aboutFindAllString(srcString string, reg *regexp.Regexp) {
	fmt.Println("\n----------FindAllString---------")

	// FindAllString，与FindAll基本一样
	// 接收源字符串的string形式，返回值是string数组

	fmt.Println("FindAllString, -1")

	ret4 := reg.FindAllString(srcString, -1)
	for _, item := range ret4 {
		fmt.Printf("%s ", item)
	}

	fmt.Println("\nFindAllString, 1")

	ret5 := reg.FindAllString(srcString, 1)
	for _, item := range ret5 {
		fmt.Printf("%s ", item)
	}

	fmt.Println("\nFindAllString, 666")

	ret6 := reg.FindAllString(srcString, 666)
	for _, item := range ret6 {
		fmt.Printf("%s ", item)
	}
}

func aboutFind(srcString string, reg *regexp.Regexp) {

	// Find返回第一个匹配项
	// 接收源字符串的[]byte形式，返回[]byte

	fmt.Println("\n-----------Find----------")

	ret := reg.Find([]byte(srcString))

	fmt.Printf("%s ", string(ret))
}

func aboutFindString(srcString string, reg *regexp.Regexp) {

	// 与Find基本一样，返回第一个匹配项
	// 接收源字符串的string形式，返回string

	fmt.Println("\n-----------FindString----------")

	ret := reg.FindString(srcString)

	fmt.Printf("%s ", ret)
}

func aboutFindAllSubmatch(srcString string, reg *regexp.Regexp) {

	// FindAllSubmatch返回所有的匹配项和其分组项
	// 接收源字符串的[]byte形式，返回[][][]byte
	// 第二个参数n表示要拿多少个匹配项
	// n=-1时拿全部；n>匹配总个数时拿全部

	fmt.Println("\n-----------FindAllSubmatch----------")

	var (
		ret [][][]byte
	)

	ret = reg.FindAllSubmatch([]byte(srcString), -1)

	for _, out := range ret { // 所有【匹配项以及其分组项组成的数组】的数组
		fmt.Printf("匹配项：%s 分组1：%s 分组2：%s 分组3：%s\n",
			string(out[0]), string(out[1]), string(out[2]), string(out[3]))
	}

}

func aboutFindAllStringSubmatch(srcString string, reg *regexp.Regexp) {

	// FindAllStringSubmatch与FindAllSubmatch基本相同，返回所有的匹配项和其分组项
	// 接收源字符串的[]byte形式，返回[][]string
	// 第二个参数n表示要拿多少个匹配项
	// n=-1时拿全部；n>匹配总个数时拿全部

	fmt.Println("\n-----------FindAllStringSubmatch----------")

	var (
		ret [][]string
	)

	ret = reg.FindAllStringSubmatch(srcString, -1)

	for _, out := range ret { // 所有【匹配项以及其分组项组成的数组】的数组
		fmt.Printf("匹配项：%s 分组1：%s 分组2：%s 分组3：%s\n",
			out[0], out[1], out[2], out[3])
	}

}

func aboutFindSubmatch(srcString string, reg *regexp.Regexp) {

	// FindSubmatch返回第一个匹配项及其分组项
	// 接收源字符串的[]byte形式，返回[][]byte

	fmt.Println("\n-----------FindSubmatch----------")

	var (
		ret [][]byte
	)

	ret = reg.FindSubmatch([]byte(srcString))

	fmt.Printf("第一个匹配项：%s，分组1：%s，分组2：%s，分组3：%s",
		string(ret[0]), string(ret[1]), string(ret[2]), string(ret[3]))
}

func aboutFindStringSubmatch(srcString string, reg *regexp.Regexp) {

	// FindStringSubmatch与FindSubmatch相似，返回第一个匹配项及其分组项
	// 接收源字符串的string形式，返回[]string

	fmt.Println("\n-----------FindStringSubmatch----------")

	var (
		ret []string
	)

	ret = reg.FindStringSubmatch(srcString)

	fmt.Printf("第一个匹配项：%s，分组1：%s，分组2：%s，分组3：%s",
		ret[0], ret[1], ret[2], ret[3])
}
