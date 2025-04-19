# GraphQL 待办事项应用项目介绍

这是一个使用 Go 语言和 GraphQL 技术栈构建的简单待办事项 (Todo) 应用程序。项目利用 gqlgen 框架自动生成大部分 GraphQL 相关代码，让开发者可以专注于业务逻辑的实现。

## 项目结构

项目主要包含以下组件：

1. **GraphQL Schema**：位于 `graph/schema.graphqls`定义了 API 的类型、查询和变更操作
2. **生成的代码**：
   - `graph/generated.go`：自动生成的 GraphQL 执行引擎
   - `graph/model/models_gen.go`：根据 schema 生成的数据模型
3. **业务逻辑**：
   - `graph/resolver.go`：依赖注入和数据存储
   - `graph/schema.resolvers.go`：查询和变更操作的实现
4. **服务器入口**：`server.go`负责设置和启动 HTTP 服务器
5. 

## 数据模型

项目定义了两个主要的数据模型：


1. **Todo**：包含 ID、文本内容、完成状态和关联用户
2. **User**：包含 ID 和名称

## 数据存储

这个项目使用内存存储数据，没有连接外部数据库：

- 所有的待办事项都存储在 `Resolver` 结构体的 `todos` 切片中
- 数据只在应用程序运行期间保存，服务重启后数据会丢失
- 这种设计适合快速原型开发和学习，但在实际生产环境中应替换为持久化存储解决方案

## 功能特性

该应用程序提供以下 GraphQL 操作：

1. **查询 (Query)**：
   - `todos`：获取所有待办事项的列表

2. **变更 (Mutation)**：
   - `createTodo`：创建新的待办事项

## 技术栈

- **Go**：编程语言 (版本 1.24.2)
- **gqlgen**：Go 的 GraphQL 服务器库 (版本 0.17.70)
- **GraphQL Playground**：提供一个交互式界面来测试查询

## 如何使用

1. **启动服务器**：
   ```
   go run server.go
   ```

2. **访问 GraphQL Playground**：
   在浏览器中打开 http://localhost:8080/

3. **示例查询**：
   ```graphql
   query {
     todos {
       id
       text
       done
       user {
         id
         name
       }
     }
   }
   ```

4. **创建新待办事项**：
   ```graphql
   mutation {
     createTodo(input: {text: "学习 GraphQL", userId: "1"}) {
       id
       text
       done
       user {
         id
         name
       }
     }
   }
   ```

这个项目是学习 GraphQL 和 Go 集成的很好的起点，提供了基本的 CRUD 操作示例，并展示了如何使用 gqlgen 框架高效地构建 GraphQL API。 
