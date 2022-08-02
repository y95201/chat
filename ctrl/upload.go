/*
 * @Description:
 * @Author: Y95201
 * @Date: 2022-08-01 10:09:37
 * @LastEditors: Y95201
 * @LastEditTime: 2022-08-01 17:16:29
 */
package ctrl

import (
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"os"
	"strings"
	"time"
	"weChat/util"
)

//创建文件
func init() {
	pathName := "./img"
	_, err := os.Stat(pathName)
	if err != nil {
		if os.IsNotExist(err) {
			os.Mkdir(pathName, 0777) //创建目录
		}
	}
	//os.MkdirAll("./img", os.ModePerm)
}

func Upload(w http.ResponseWriter, r *http.Request) {
	UploadLocal(w, r)
}

func UploadLocal(writer http.ResponseWriter, request *http.Request) {
	//获得上传的源文件s
	File, FileHeader, err := request.FormFile("file")
	if err != nil {
		util.RespFail(writer, err.Error())
	}
	//关闭文件，释放资源
	defer File.Close()
	//创建一个新文件
	suffix := ".png"
	//获取文件后缀
	fileName := FileHeader.Filename
	tmp := strings.Split(fileName, ".")
	if len(tmp) > 1 {
		suffix = "." + tmp[len(tmp)-1]
	}

	//生成文件名
	filename := fmt.Sprintf("%d%04d%s", time.Now().Unix(), rand.Int31(), suffix)
	dotfile, err := os.Create("./img/" + filename)
	if err != nil {
		util.RespFail(writer, err.Error())
		return
	}

	//将源文件内容copy到新文件
	_, err = io.Copy(dotfile, File)
	if err != nil {
		util.RespFail(writer, err.Error())
		return
	}
	defer dotfile.Close()
	//将新文件路径转换成url地址
	url := "/img/" + filename
	//响应到前端
	util.RespOk(writer, url, "")
}
