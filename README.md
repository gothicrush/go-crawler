## 开发流程

### $ 单机单线程版爬虫

#### Step 1：学习前置知识

* [http库的Get方法使用]()
* [判断网页编码并改变读取编码]()
* [正则表达式库的查找和获取分组的使用]()

#### Step 2：了解爬取流程

* 流程图

  ![](C:\Users\narli\Desktop\go-crawler\_images\爬取流程图.PNG)

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

![](C:\Users\narli\Desktop\go-crawler\_images\单机版架构图.PNG)

#### Step 5： 实现单机单线程版爬虫

* engine模块
* parser模块
* fetch模块

