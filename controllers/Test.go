package controllers

import (
	"fmt"
	"io"
	"mime"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func TestFunc1(ctx *gin.Context) {
	id := ctx.Param("id")
	ctx.JSON(http.StatusOK, gin.H{
		":id": id,
	})
}
func TestFunc2(ctx *gin.Context) {
	page := ctx.Query("page")
	size := ctx.Query("size")
	ctx.JSON(http.StatusOK, gin.H{
		"page": page,
		"size": size,
	})
}
func Download(ctx *gin.Context) {
	// 获取文件名
	filename := ctx.Param("filename")
	// 返回对应文件
	ctx.FileAttachment("./uploads/"+filename, "文件")
}
func TestFunc3(ctx *gin.Context) {
	type user struct {
		Name string `json:"name"` // 添加json标签以便自定义字段名
		Age  int    `json:"age"`
	}
	var user1 user
	ctx.ShouldBindJSON(&user1)
	ctx.JSON(http.StatusOK, gin.H{
		"user": user1,
	})
}

func TestFunc4(ctx *gin.Context) {
	// info[name] 徐迪
	// info[age] 20
	res := ctx.PostFormMap("info")
	fmt.Println(res)
	ctx.JSON(http.StatusOK, gin.H{
		"msg": res,
	})
}

func TestFunc5(ctx *gin.Context) {
	file, err := ctx.FormFile("file")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// 指定文件保存路径
	dst := filepath.Join("uploads", file.Filename)
	// 保存文件到本地
	if err := ctx.SaveUploadedFile(file, dst); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error":   "文件保存失败",
			"details": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message":  "文件上传成功",
		"filename": file.Filename,
		"size":     file.Size,
		"header":   file.Header,
	})
}
func UploadFiles(ctx *gin.Context) {
	// 获取gin解析好的multipart表单
	form, _ := ctx.MultipartForm()
	// 根据键值取得对应的文件列表
	files := form.File["files"]
	// 遍历文件列表，保存到本地
	for _, file := range files {
		// 指定文件保存路径
		dst := filepath.Join("uploadsFiles", file.Filename)
		// 保存文件到本地
		err := ctx.SaveUploadedFile(file, dst)
		if err != nil {
			ctx.String(http.StatusBadRequest, "upload failed")
			return
		}
	}
	// 返回结果
	ctx.String(http.StatusOK, "upload %d files successfully!", len(files))
}
func TestFunc6(ctx *gin.Context) {
	// 要有请求头  Content-Disposition ： attachment; filename=example.bin
	// 解析Content-Disposition头获取文件名
	contentDisposition := ctx.GetHeader("Content-Disposition")
	var filename string
	if contentDisposition != "" {
		// 解析形如: "attachment; filename=\"example.bin\""的内容
		_, params, err := mime.ParseMediaType(contentDisposition)
		if err == nil {
			filename = filepath.Base(params["filename"]) // 使用filepath.Base防止路径遍历
		}
	}
	// 从请求体中读取二进制数据
	file, err := os.Create("uploads/" + filename)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer file.Close()
	// 将请求体中的数据写入文件
	if _, err := io.Copy(file, ctx.Request.Body); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	// 返回响应
	ctx.JSON(http.StatusOK, gin.H{
		"message":  "success",
		"filename": "uploaded_file.bin",
	})
}

func TestFunc7(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "success...",
	})
}

func TestFunc8(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "success...",
	})
}

func TestFunc9(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "success...",
	})
}

func TestFunc10(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "success...",
	})
}

func TestFunc11(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "success...",
	})
}

func TestFunc12(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "success...",
	})
}

func TestFunc13(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "success...",
	})
}
