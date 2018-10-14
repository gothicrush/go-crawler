package main

import (
	"bufio"
	"errors"
	"fmt"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
	"io"
	"io/ioutil"
	"net/http"
)

func main() {

    var (
    	oldReader io.Reader
    	newReader *bufio.Reader
    	err error
    	e encoding.Encoding
    	data []byte
	)

    /*
    获取IOReader
     */
    if oldReader, err = getIOReader("http://www.zhenai.com/zhenghun"); err != nil {
    	fmt.Println("获取IOReadr失败：", err)
    	return
	}

    /*
    生成一个新的IOReader，因为在转编码过程中，需要读取了一些数据
    而IOReader一旦读取后，已读取内容不能再次读取
     */
    newReader = bufio.NewReader(oldReader)

    /*
    利用 golang.org/x/net/html的DetermineEncoding方法 获取页面的编码
     */
     e = getEncoding(newReader)

     /*
     根据页面编码，将Reader转为UTF8编码格式的Reader
     利用 golang.org/x/text/transform中NewReader方法进行转换
      */

      oldReader = transform.NewReader(oldReader,e.NewDecoder())

      /*
      读取Reader中的数据
       */
       if data,err = ioutil.ReadAll(oldReader); err != nil {
       	    fmt.Println("IOReader读取失败", err)
       	    return
	   }

       fmt.Println(string(data))
}

/*
获取IOReader
 */
func getIOReader(url string) (io.Reader,error) {

	var (
		resp *http.Response
		err error
	)

	if resp,err = http.Get(url); err != nil {
		fmt.Println("Get请求失败")
		return nil, err
	}

    // defer resp.Body.Close() //不能关闭，因为要进行返回

	if resp.StatusCode != http.StatusOK {
		fmt.Println("非正常响应，状态码：",resp.StatusCode)
		return nil, errors.New(fmt.Sprintf("非正常响应，状态码：%s",resp.StatusCode))
	}

	return resp.Body, nil

}

/*
使用 golang.org/x/net/html/charset的DetermineEncoding
获取页面编码
 */
 func getEncoding(r *bufio.Reader) encoding.Encoding {

    var (
    	data []byte
    	err error
    	e encoding.Encoding
	)

    /*
    DeterminEncoding方法需要数据源的一些字节数据来进行判断
     */
    if data, err = r.Peek(1024); err != nil {
    	fmt.Println("读取前1024字节数据失败，返回默认utf8编码")
    	return unicode.UTF8
	}

    e, _, _ = charset.DetermineEncoding(data,"")

    return e
 }