# server 后端服务器

## 设计拆分思想

从 0 到 1 功能拆分方法和思想, 参考 [Go语言实战流媒体网站-bilibili](https://www.bilibili.com/video/BV1zy4y1J7LS)

1. 拆分对象: 拆分对应的操作对象/资源， 抽象成对应的结构体， 生成对应表结构
    + **SDK Driver** 表 : 注册 SDK 信息
    + **DNS解析商账号** 表 : 存储DNS商的账号信息， 与 SDK_Driver 有绑定关系
    + **域名** 表 : 域名列表， 与 DNS解析上账户 有绑定关系
    + **解析记录** 表: 解析记录， 与 **域名** 表是从属关系。
    + **解析记录的变更历史** 表: 解析记录的 ChangeLog， 与 **解析记录** 是从属关系。
2. 寻找关系: 寻找各个对象/资源之间的从属关系， 与设计表关联关系。
3. 设计 RESTful API: 依照上述的资源关系， 根据 RESTful 范式， 设计 API 对应结构。

```go
// BasePath 及 版本信息
/dnsx/:version // /dnsx/v0 

// 注册 Driver
POST /dnsx/v0/driver/:name

// dnsx解析商账户
    /dnsx/v0/dnsaccount/:name

// 域名
    /dnsx/v0/domain/:domain

// 解析记录
GET /dnsx/v0/domain/:domain/record // 列出所有解析记录

POST /dnsx/v0/domain/:domain/record/:record // 创建一条解析记录
GET  /dnsx/v0/domain/:domain/record/:record // 获取解析记录详情

// 解析记录 ChangeLog
POST /dnsx/v0/domain/:domain/record/:record/changelog  // 创建一条变更记录
GET  /dnsx/v0/domain/:domain/record/:record/changelog  // 获取所有变更记录

```

4. 构建代码: 根据设计的 API ， 数据 handler 代码， 处理数据。
