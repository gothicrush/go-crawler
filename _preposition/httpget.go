package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {

	var (
		resp *http.Response
		err error
		data []byte
	)

	/*
	http.Get(url)，用http的get方法向url发送请求
	返回一个响应结构体指针和一个error
	 */
	if resp,err = http.Get("http://www.zhenai.com/zhenghun"); err != nil {
		fmt.Println("发送get请求失败")
		return
	}

	/*
	一旦http.Get方法正确执行，就会生成http.Response对象
	该Response对象有一个字段 Body，该字段拥有响应数据主体的信息，需要手动Close掉
	 */
	defer resp.Body.Close()

	/*
	resp包含有响应状态码，resp.StatusCode
	 */
	if resp.StatusCode != http.StatusOK {
		fmt.Println("非正常的响应，状态码：", resp.StatusCode)
		return
	}

	/*
	http.Response.Body字段拥有数据主体的信息，实现了 io.Reader 接口
	用ioutil的ReadAll方法，读取响应的数据主体的信息，返回字节数组，可转为string查看
	 */
	 if data, err = ioutil.ReadAll(resp.Body); err != nil {
	 	fmt.Println("读取Body信息失败")
	 	return
	 }

     fmt.Println(string(data))
}