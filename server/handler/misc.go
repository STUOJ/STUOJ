package handler

import (
	"STUOJ/external/yuki"
	"STUOJ/internal/model"
	"STUOJ/utils"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

func UploadImage(c *gin.Context) {
	role, err := strconv.Atoi(c.Query("role"))
	if err != nil || model.GetAlbumName(uint8(role)) == "unknown" {
		c.JSON(http.StatusBadRequest, model.RespError("参数错误", nil))
		return
	}
	file, err := c.FormFile("file")
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, model.RespError("文件上传失败", nil))
		return
	}

	// 保存文件
	dst := fmt.Sprintf("tmp/%s", utils.GetRandKey())
	if err := c.SaveUploadedFile(file, dst); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, model.RespError("文件上传失败", nil))
		return
	}
	defer os.Remove(dst)

	image, err := yuki.UploadImage(dst, uint8(role))
	if err != nil {
		c.JSON(http.StatusBadRequest, model.RespError("上传失败", nil))
		return
	}
	c.JSON(http.StatusOK, model.RespOk("上传成功", image.Url))
}

type ReqEmail struct {
	Email string `json:"email" binding:"required"`
}

func SendVerificationCode(c *gin.Context) {
	var req ReqEmail
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, model.RespError("参数错误", nil))
		return
	}
	if err := utils.SendVerificationCode(req.Email); err != nil {
		c.JSON(http.StatusBadRequest, model.RespError(err.Error(), nil))
		return
	}
	c.JSON(http.StatusOK, model.RespOk("OK", nil))
}

type ReqVerify struct {
	Email string `json:"email" binding:"required"`
	Code  string `json:"code" binding:"required"`
}
