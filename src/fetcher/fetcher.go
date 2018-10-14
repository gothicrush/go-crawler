package fetcher

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
	"log"
	"net/http"
)

func Fetching(url string) []byte {

	var (
		rb     io.ReadCloser
		reader io.Reader
		err    error
		data   []byte
	)

	if rb, err = request(url); err != nil {
		log.Println(err)
		return nil
	}

	defer rb.Close()

	reader = transformReader(rb)

	if data, err = readData(reader); err != nil {
		log.Println(err)
		return nil
	}

	return data
}

func request(url string) (io.ReadCloser, error) {

	var (
		err    error
		resp   *http.Response
		req    *http.Request
		client *http.Client
	)

	client = &http.Client{}
	req, err = http.NewRequest("GET", url, nil)

	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/39.0.2171.71 Safari/537.36")

	if resp, err = client.Do(req); err != nil {
		err = errors.New("Http Get请求失败")
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		err = errors.New(fmt.Sprintf("Http Get获取非正常响应，URL：%s, 状态码：%s", url, resp.StatusCode))
		return nil, err
	}

	return resp.Body, nil
}

func transformReader(r io.Reader) io.Reader {

	var (
		bufr *bufio.Reader
		e    encoding.Encoding
	)

	bufr = bufio.NewReader(r)

	e = determineEncoding(bufr)

	return transform.NewReader(bufr, e.NewDecoder())
}

func determineEncoding(bufr *bufio.Reader) encoding.Encoding {

	var (
		data []byte
		err  error
		e    encoding.Encoding
	)

	if data, err = bufr.Peek(1024); err != nil {
		log.Println("获取Reader前1024个字节失败，直接返回UTF8编码")
		return unicode.UTF8
	}

	e, _, _ = charset.DetermineEncoding(data, "")

	return e
}

func readData(r io.Reader) ([]byte, error) {

	var (
		data []byte
		err  error
	)

	if data, err = ioutil.ReadAll(r); err != nil {
		err = errors.New("ioutil读取数据失败，返回空白数据")
		return nil, err
	}

	return data, nil
}
