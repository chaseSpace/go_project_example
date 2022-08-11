--- 
这个工程演示了一个单体go工程的项目布局，在这个工程内可能包含多个app；
所有的目录都是小写，多语义请加下划线，目录都是可选的，根据需要创建

#### 如果想要启动一个新的单体go工程，可直接clone本仓库，然后基于它开发你的应用。

#### 编译运行

```bash
cd applicaiton/app_apple
go run .
```

#### 结构

```
|-- application
|   |-- app_apple    // 单个app
|   |   |-- api      // api层（只能）调用service层接口聚合数据，提供给外部
|   |   |   |-- api.go
|   |   |   `-- route.go
|   |   |-- main.go  // app启动入口
|   |   |-- proto    // req&rsp 协议
|   |   |   |-- base.go
|   |   |   `-- proto.go
|   |   |-- repo     // repo包含model和repo，不再独立model层（降低复杂度），通过规范文件命名区分即可
|   |   |   |-- model_user.go
|   |   |   |-- model_user_other.go
|   |   |   `-- repo_user.go
|   |   `-- service  // 一个app下可能有多个service，根据业务分类划分
|   |       `-- service_user.go
|   `-- app_banana   // 另一个app
|-- config           // 统一的config目录
|   `-- config.go
|-- global           // 所有app共享代码，可以是全局变量，也可只是公共代码，根据个人习惯
|   `-- infra        // infrastructure（基础架构）
|       `-- db.go
|-- go.mod
|-- go.sum
|-- readme.md
`-- utils            // 工具包
    |-- build.go
    `-- proto_util.go
```