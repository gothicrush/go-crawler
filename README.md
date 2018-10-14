## 开发流程

### $ 单机单线程版爬虫

#### Step 1：学习前置知识

* [使用HTTP库发出GET请求](https://github.com/gothicrush/go-crawler/blob/master/_preposition/httpget.go) 
* [判断网页编码并改变读取编码](https://github.com/gothicrush/go-crawler/blob/master/_preposition/readerencode.go) 
* [正则表达式库的使用](https://github.com/gothicrush/go-crawler/blob/master/_preposition/regexp.go) 

#### Step 2：了解爬取流程

* 流程图

  ![爬取流程图.PNG](https://github.com/gothicrush/go-crawler/blob/master/_images/%E7%88%AC%E5%8F%96%E6%B5%81%E7%A8%8B%E5%9B%BE.PNG) 

  1. 通过【城市列表解析器】爬取【城市列表】页面获取【城市链接】
  2. 通过【城市解析器】爬取【城市链接】对应页面获取【用户链接】
  3. 通过【用户解析器】爬取【用户链接】对应页面获取【用户信息】

#### Step 3：了解解析器 Parser

* 输入：二进制UTF8编码的文本

* 输出：ParseRequest数组，每个ParseRequest包括要解析的URL和解析函数

  ```go
  type ParseRequest struct {
  	URL       string
  	ParseFunc func([]byte) []ParseRequest
  }
  ```


#### Step 4：了解单机单线程版爬虫系统架构

![单机版架构图.PNG](https://github.com/gothicrush/go-crawler/blob/master/_images/%E5%8D%95%E6%9C%BA%E7%89%88%E6%9E%B6%E6%9E%84%E5%9B%BE.PNG)

#### Step 5： 实现单机单线程版爬虫

* [engine模块](https://github.com/gothicrush/go-crawler/tree/master/src/engine) 
* [parser模块](https://github.com/gothicrush/go-crawler/tree/master/src/parser) 
* [fetcher模块](https://github.com/gothicrush/go-crawler/tree/master/src/fetcher) 
* [main函数](https://github.com/gothicrush/go-crawler/blob/master/src/main.go) 

