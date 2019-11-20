package cli

import (
	"github.com/gogf/gf/net/ghttp"
)

type Controller struct {
	
}

func (c *Controller) Md5(r *ghttp.Request) {
	// https://goframe.org/cli/project/md5
	r.Response.Write("d5aec300905b18c9b128bfa6939449fd")
}

func (c *Controller) Zip(r *ghttp.Request) {
	// 这里直接返回zip包文件
	// 包文件下载地址：https://gf.cdn.johng.cn/cli/project/zip?d5aec300905b18c9b128bfa6939449fd
	r.Response.ServeFileDownload("public/zip/zip.zip")
}
