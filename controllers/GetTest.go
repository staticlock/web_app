package controllers

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

// GET + 路径参数  请求示例: GET /api/v1/path/123
func TestGetPathParam(ctx *gin.Context) {
	id := ctx.Param("id")
	ctx.JSON(http.StatusOK, gin.H{
		"method": ctx.Request.Method,
		"url":    ctx.Request.URL,
		"id":     id,
	})
}

// GET + 查询参数  请求示例: GET /api/v1/query?page=1&size=20
func TestGetQuery(ctx *gin.Context) {
	page := ctx.DefaultQuery("page", "1")
	size := ctx.DefaultQuery("size", "10")
	ctx.JSON(http.StatusOK, gin.H{
		"method": ctx.Request.Method,
		"url":    ctx.Request.URL,
		"page":   page,
		"size":   size,
	})
}

// GET + 查询参数绑定到结构体 请求示例: GET /api/v1/query/struct?id=123&name=John
func TestGetQueryStruct(ctx *gin.Context) {
	type Query struct {
		ID   string `form:"id"`
		Name string `form:"name"`
	}
	var query Query
	if err := ctx.ShouldBindQuery(&query); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"method": ctx.Request.Method,
		"url":    ctx.Request.URL,
		"query":  query,
	})
}

// TestHeaderParams 测试请求头参数
// 请求示例: GET /api/v1/header
func TestHeaderParams(ctx *gin.Context) {
	token := ctx.GetHeader("Authorization")
	contentType := ctx.GetHeader("Content-Type")
	ctx.JSON(http.StatusOK, gin.H{
		"token":       token,
		"contentType": contentType,
	})
}

// TestCookieParams 测试Cookie参数
// 请求示例: GET /api/v1/cookie
func TestCookieParams(ctx *gin.Context) {
	token, err := ctx.Cookie("token")
	if err != nil {
		token = "not set"
	}
	ctx.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}

// TestSetCookie 测试设置Cookie
// 请求示例: GET /api/v1/set-cookie
func TestSetCookie(ctx *gin.Context) {
	ctx.SetCookie("token", "abc123", 3600, "/", "localhost", false, true)
	ctx.JSON(http.StatusOK, gin.H{
		"message": "cookie set",
	})
}

// TestDownloadFile 测试文件下载
// 请求示例: GET /api/v1/download/:filename
func TestDownloadFile(ctx *gin.Context) {

	filename := ctx.Param("filename")
	filePath := filepath.Join("uploads", filename)

	// 检查文件是否存在
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "file not found"})
		return
	}
	//// 设置响应头
	//ctx.Header("Content-Disposition", fmt.Sprintf("attachment; filename=%s", filename))
	//ctx.Header("Content-Type", "application/octet-stream")
	//// 发送文件
	//ctx.File(filePath)
	// 返回对应文件
	ctx.FileAttachment("./uploads/"+filename, filename)
}

// TestQueryStructBinding 测试查询参数绑定到结构体
// 请求示例: GET /api/v1/query-struct?name=John&age=30
func TestQueryStructBinding(ctx *gin.Context) {
	type Query struct {
		Name string `form:"name"`
		Age  int    `form:"age"`
	}
	var query Query
	if err := ctx.ShouldBindQuery(&query); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"query": query,
	})
}

// TestURIStructBinding 测试URI参数绑定到结构体
//
//	GET /api/v1/uri-struct/:name/:age
//
// 请求示例: GET /api/v1/uri-struct/John/30
func TestURIStructBinding(ctx *gin.Context) {
	type User struct {
		Name string `uri:"name"`
		Age  int    `uri:"age"`
	}

	var user User
	if err := ctx.ShouldBindUri(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"user": user,
	})
}

// TestStreamResponse 测试流式响应
// 请求示例: GET /api/v1/stream
func TestStreamResponse(ctx *gin.Context) {
	// 设置响应头
	ctx.Writer.Header().Set("Content-Type", "text/plain")
	ctx.Writer.Header().Set("Transfer-Encoding", "chunked")
	// 流式写入响应
	for i := 0; i < 10; i++ {
		ctx.Writer.Write([]byte(fmt.Sprintf("Chunk %d\n", i)))
		ctx.Writer.Flush()
	}
}

// TestSSE 测试服务器发送事件(Server-Sent Events)
// 请求示例: GET /api/v1/sse
func TestSSE(ctx *gin.Context) {
	// 设置SSE相关响应头
	ctx.Writer.Header().Set("Content-Type", "text/event-stream")
	ctx.Writer.Header().Set("Cache-Control", "no-cache")
	ctx.Writer.Header().Set("Connection", "keep-alive")
	// 发送事件
	for i := 0; i < 5; i++ {
		event := fmt.Sprintf("data: Message %d\n\n", i)
		ctx.Writer.Write([]byte(event))
		ctx.Writer.Flush()
	}
}
