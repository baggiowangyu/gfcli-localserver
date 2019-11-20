package router

import (
    "gf-app/app/api/cli"
    "gf-app/app/api/hello"
    "github.com/gogf/gf/frame/g"
)

// 统一路由注册.
func init() {
    g.Server().BindHandler("/", hello.Handler)

    cli_controller := new(cli.Controller)
    g.Server().BindObject("GET:/cli/project", cli_controller, "Md5")
    g.Server().BindObject("GET:/cli/project", cli_controller, "Zip")
}
