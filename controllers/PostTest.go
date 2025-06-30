package controllers

import (
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin/binding"

	"github.com/gin-gonic/gin"
)

// TestPostForm 测试表单提交
// 请求示例: POST /api/v1/users (Content-Type: application/x-www-form-urlencoded)
func TestPostForm(ctx *gin.Context) {
	name := ctx.PostForm("name")
	age := ctx.DefaultPostForm("age", "18")
	ctx.JSON(http.StatusOK, gin.H{
		"method": ctx.Request.Method,
		"url":    ctx.Request.URL,
		"name":   name,
		"age":    age,
	})
}

// TestPostFormMap 测试表单Map数据
// 请求示例: POST /api/v1/users (Content-Type: application/x-www-form-urlencoded)
// 表单数据: user[name]=John&user[age]=30
func TestPostFormMap(ctx *gin.Context) {
	user := ctx.PostFormMap("user")
	ctx.JSON(http.StatusOK, gin.H{
		"method": ctx.Request.Method,
		"url":    ctx.Request.URL,
		"user":   user,
	})
}

// TestPostJSON 测试JSON数据绑定
// 请求示例: POST /api/v1/users (Content-Type: application/json)
// 请求体: {"name": "John", "age": 30}
func TestPostJSON(ctx *gin.Context) {
	type User struct {
		Name string `json:"name" binding:"required"`
		Age  int    `json:"age" binding:"required,min=1"`
	}
	var user User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"method": ctx.Request.Method,
		"url":    ctx.Request.URL,
		"user":   user,
	})
}

// TestSingleFileUpload 测试单文件上传
// 请求示例: POST /api/v1/upload (Content-Type: multipart/form-data)
func TestSingleFileUpload(ctx *gin.Context) {
	file, err := ctx.FormFile("file")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 确保上传目录存在
	if err := os.MkdirAll("uploads", os.ModePerm); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// 保存文件
	dst := filepath.Join("uploads", file.Filename)
	if err := ctx.SaveUploadedFile(file, dst); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"method":   ctx.Request.Method,
		"url":      ctx.Request.URL,
		"filename": file.Filename,
		"size":     file.Size,
		"path":     dst,
	})
}

// TestMultiFileUpload 测试多文件上传
// 请求示例: POST /api/v1/upload/multi (Content-Type: multipart/form-data)
func TestMultiFileUpload(ctx *gin.Context) {
	form, err := ctx.MultipartForm()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	files := form.File["files"]
	if len(files) == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "no files uploaded"})
		return
	}
	// 确保上传目录存在
	if err := os.MkdirAll("uploadsFiles", os.ModePerm); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	var uploadedFiles []gin.H
	for _, file := range files {
		dst := filepath.Join("uploads", filepath.Base(file.Filename))
		if err := ctx.SaveUploadedFile(file, dst); err != nil {
			log.Printf("Failed to save file %s: %v", file.Filename, err)
			continue
		}
		uploadedFiles = append(uploadedFiles, gin.H{
			"filename": file.Filename,
			"size":     file.Size,
			"path":     dst,
		})
	}
	ctx.JSON(http.StatusOK, gin.H{
		"count": len(uploadedFiles),
		"files": uploadedFiles,
	})
}

// TestRawBody 测试原始请求体
// 请求示例: POST /api/v1/raw-body
func TestRawBody(ctx *gin.Context) {
	body, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"rawBody": string(body),
	})
}

// TestXMLBinding 测试XML数据绑定
// 请求示例: POST /api/v1/xml (Content-Type: application/xml)
// 请求体: <user><name>John</name><age>30</age></user>
func TestXMLBinding(ctx *gin.Context) {
	type User struct {
		Name string `xml:"name"`
		Age  int    `xml:"age"`
	}

	var user User
	if err := ctx.ShouldBindXML(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"user": user,
	})
}

// TestYAMLBinding 测试YAML数据绑定
// 请求示例: POST /api/v1/yaml (Content-Type: application/x-yaml)
// 请求体: name: John\nage: 30
func TestYAMLBinding(ctx *gin.Context) {
	type User struct {
		Name string `yaml:"name"`
		Age  int    `yaml:"age"`
	}

	var user User
	if err := ctx.ShouldBindYAML(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"user": user,
	})
}

// TestFormStructBinding 测试表单数据绑定到结构体
// 请求示例: POST /api/v1/form-struct (Content-Type: application/x-www-form-urlencoded)
// 表单数据: name=John&age=30
func TestFormStructBinding(ctx *gin.Context) {
	type User struct {
		Name string `form:"name"`
		Age  int    `form:"age"`
	}
	var user User
	if err := ctx.ShouldBindWith(&user, binding.Form); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"user": user,
	})
}
