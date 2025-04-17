package main

import (
	"gqlgen-example/graph"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/lru"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/vektah/gqlparser/v2/ast"
)

const defaultPort = "8080"

func main() {
	// 从环境变量获取端口，如果未设置则使用默认端口
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	// 创建一个新的GraphQL服务器实例
	// 使用graph包中的NewExecutableSchema函数创建可执行的GraphQL schema
	// 并传入一个空的Resolver实例作为解析器
	srv := handler.New(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}))

	// 添加支持的HTTP传输方式
	srv.AddTransport(transport.Options{}) // 支持OPTIONS请求（用于CORS预检）
	srv.AddTransport(transport.GET{})     // 支持GET请求
	srv.AddTransport(transport.POST{})    // 支持POST请求

	// 设置查询缓存，可以缓存1000个已解析的查询文档
	// 这有助于提高性能，避免重复解析相同的查询
	srv.SetQueryCache(lru.New[*ast.QueryDocument](1000))

	// 启用GraphQL内省功能，允许客户端查询schema信息
	srv.Use(extension.Introspection{})

	// 启用自动持久化查询功能
	// 这允许客户端发送查询的哈希值而不是完整查询，减少网络传输
	srv.Use(extension.AutomaticPersistedQuery{
		Cache: lru.New[string](100), // 缓存100个持久化查询
	})

	// 设置HTTP路由处理
	// 根路径提供GraphQL playground界面，方便开发和测试
	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	// /query路径处理实际的GraphQL请求
	http.Handle("/query", srv)

	// 打印服务器启动信息
	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	// 启动HTTP服务器，如果发生错误则记录并终止程序
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
