package boot

import "github.com/gogf/gf/frame/g"

func init() {
    g.Server().SetServerRoot("public")
    g.Server().SetPort(443)

    // 开启https
    g.Server().EnableHTTPS("cert/server.crt", "cert/server.key")
}

