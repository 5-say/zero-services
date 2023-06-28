# zero-services

与业务无关的 go-zero rpc 服务集

> 封装独立于业务的 rpc 服务

## 开发流程

在主项目中执行以下命令

```sh
cd work
git clone git@github.com:5-say/zero-services.git
cd ..
go work init ./work/zero-services
code ./work/zero-services
```

> 以 jwtx rpc 服务 为例

- 数据结构设计
  - 在 `define/jwtx.sql` 内分别定义各表数据结构
  - 利用 vscode 插件创建数据库、导入数据结构
  - 在根目录下执行 `make jd` 生成 dao 文件
- api 设计
  - 在 `define/jwtx.proto` 下定义 rpc 结构
  - 在根目录下执行 `make jr` 生成 rpc 基础代码
  - 在根目录下执行 `make j`  运行 rpc 服务
- 编写业务逻辑
  - 配置文件路径 `private/jwtx/rpc/etc`
  - 逻辑文件路径 `private/jwtx/rpc/internal/logic`
- 业务相关的通用代码
  - 目录 `private/jwtx/rpc/common`
  - 仅供内部调用
