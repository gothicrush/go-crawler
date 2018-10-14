## 开发流程

### $ 单机单线程版爬虫

#### Step 1：学习前置知识

* [http库的Get方法使用]()
* [判断网页编码并改变读取编码]()
* [正则表达式库的查找和获取分组的使用]()

#### Step 2：了解爬取流程

* 流程图

  ![](C:\Users\narli\Desktop\go-crawler\_images\爬取流程图.PNG)

* 通过【城市列表解析器】爬取【城市列表】页面获取【城市链接】

* 通过【城市解析器】爬取【城市链接】对应页面获取【用户链接】

* 通过【用户解析器】爬取【用户链接】对应页面获取【用户信息】

#### Step 3：了解解析器 Parser

* 输入：UTF8编码的文本

* 输出：ParseResult集合，每个ParseResult代表一个子URL及其解析器、对应信息

  ```
  ParseResult
  |-- Info：页面下该子URL对应的信息
  |-- Request
     |-- URL：子URL
     |-- Parser：该子URL对应的解析器
  ```

* 例子

  ```bash
  ParseResult
  |-- Info：百度首页
  |-- Request
     |-- URL：https://www.baidu.com
     |-- Parser：parseFunc1
     
  ParseResult
  |-- Info：谷歌首页
  |-- Request
     |-- URL：https://www.google.com
     |-- Parser：parseFunc2
  ```

#### Step 4：了解单机单线程版爬虫系统架构

![](C:\Users\narli\Desktop\go-crawler\_images\单机版架构图.PNG)

#### Step 5： 实现单机单线程版爬虫



### $ 单机并发版爬虫

#### Step 6：单机单线程版爬虫性能分析

