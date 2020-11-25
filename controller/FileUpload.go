package controller

import (
	"github.com/google/uuid"
	"github.com/kataras/iris/v12"
	"iris-antd-admin/libs"
	"path"
)

func FileUpload(ctx iris.Context) {
	file, info, err := ctx.FormFile("file")
	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		_, _ = ctx.JSON(ApiResource(false, nil, "1001", err.Error(), 2, ctx.GetID().(string)))
		return
	}
	defer file.Close()
	fname := info.Filename
	fileExt := path.Ext(fname)
	fileNewName := uuid.New().String()
	tempWebPath := libs.Config.UploadDir + "/" + libs.TimeNowToString() + "/" + fileNewName + fileExt
	rootFilePath := libs.CWD() + tempWebPath
	_, err = ctx.SaveFormFile(info, rootFilePath)
	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		_, _ = ctx.JSON(ApiResource(false, nil, "1001", "文件上传失败", 2, ctx.GetID().(string)))
		return
	}
	var WebPath string = ""
	if libs.Config.HTTPS {
		if libs.Config.Port == 443 {
			WebPath = "https://" + libs.Config.Host + tempWebPath
		} else {
			WebPath = "https://" + libs.Config.Host + ":" + libs.ParseString(libs.Config.Port) + tempWebPath
		}

	} else {
		if libs.Config.Port == 80 {
			WebPath = "http://" + libs.Config.Host + tempWebPath
		} else {
			WebPath = "http://" + libs.Config.Host + ":" + libs.ParseString(libs.Config.Port) + tempWebPath
		}

	}
	ctx.StatusCode(iris.StatusOK)
	_, _ = ctx.JSON(ApiResource(true, map[string]string{
		"name":     fileNewName,
		"status":   "done",
		"thumbUrl": WebPath,
		"url":      WebPath,
	}, "200", "处理成功", 0, ctx.GetID().(string)))
	return
}
