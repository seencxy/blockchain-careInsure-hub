# Go Server Directory Structure

## `/api`

这个目录用于存放API的定义，包括handler，middleware和router。

- `handler`: API的处理函数，负责处理HTTP请求和返回响应。
- `middleware`: 用于处理API请求的中间件，比如认证，日志，错误处理等。
- `router`: 定义API的路由规则。

## `/app`

这个目录用于存放应用程序的主要代码。

- `core`: 存放全局常量，错误定义，和变量。
- `utils`: 存放在应用程序中使用的实用函数。
- `models`: 存放应用程序的结构体存储。

## `/bootstrap`

这个目录用于存放应用程序的初始化代码。

## `/cmd`

这个目录用于存放应用程序的入口代码。

- `app`: 应用程序的入口代码。

## `/internal`

这个目录用于存放应用程序的内部代码。

- `database`: 存放处理数据库连接和操作的代码。
- `model`: 存放数据模型的定义。
- `pkg`: 存放应用程序使用的内部包。
- `logs`: 存放应用程序的日志文件。
- `service`: 存放应用程序的业务逻辑。

## `/go.mod` 和 `/go.sum`

这两个文件是Go的模块系统生成的，用于管理项目的依赖。

- `go.mod`: 定义了项目的模块路径和依赖项。
- `go.sum`: 包含了每个依赖项的预期加密哈希。

