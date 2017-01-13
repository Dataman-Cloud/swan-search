# 多集群搜索的设计和实现
多集群搜索即在多集群之上建立一层搜索引擎，方便用户进行资源搜索（目前只支持应用的搜索)。

开发工程师可以访问多集群管理前端页面做文本关键词查询，输入可以为任意文本关键词，关键词可以命中集群名、用户名、应用名。

源码地址: [https://github.com/Dataman-Cloud/swan-search](https://github.com/Dataman-Cloud/swan-search)

##架构实现
多集群搜索的实现架构如下图所示:

<img src="./images/search-engine.jpg" width="500" />

架构实现说明:

首先将集群信息(集群名,ip列表)配置到多集群搜索服务的配置文件中,搜索层通过读取配置文件获取到每个集群的服务地址，其实也就是调度器的地址列表。  
首先通过调用每个集群的api获取应用的信息，进行解析，聚合，初始化索引。  
同时搜索引擎会启动多个并发进程，与每个集群建立长连接，来监听每个集群的事件消息，当应用／实例更新时，调度器就会发送事件消息，搜索引擎监听到消息，进行解析，聚合，更新索引。
搜索引擎对外提供api,当用户发起搜索请求时，引擎会通过关键字，在索引中进行模糊查询，返回查询结果。

实现约束:
各个集群的调度器服务都必须被多集群服务访问到。

## 详细设计:
+ API定义:
```
GET /search/v1/luckysearch?keyword=xxx   xxx为关键字，搜索方法为模糊搜索
```
前端展示: 通过api获得搜索结果后，前端会根据结果的数据结构和规约来显示主要信息和拼接链接。

索引数据结构规约:

```
Index: []string //存储对象数据的id或name或唯一标识
Document{
    ID: string,
    Name: string,
    Type: string,
    ClusterId: string,
    Param: map[string]string{
    },
    ...
}
```

+ 示例:
```
curl http://192.168.1.155:9888/search/v1/luckysearch?keyword=zliu
```

搜索结果:

```
{"code":"0","data":[{"ID":"0092-zliu-datamanmesos","Name":"0092","Type":"app","Param":{"AppId":"0092-zliu-datamanmesos","ClusterId":"","RunAs":"zliu"},"ClusterId":""},{"ID":"0-0092-zliu-datamanmesos","Name":"0-0092-zliu-datamanmesos","Type":"task","Param":{"AppId":"0092-zliu-datamanmesos","ClusterId":"","RunAs":"zliu","TaskIndex":"0"},"ClusterId":""}]}
```
+ 搜索算法:
目前可采用的搜索算法为Levenshtein distance ， 其相应的实现有fuzzysearch: https://github.com/renstrom/fuzzysearch
