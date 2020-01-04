package main

import (
	"fmt"
	golog "log"
	"net/http"

	"github.com/weihaoranW/vchat/common/g"
	"github.com/weihaoranW/vchat/demo/unit/intf"
	"github.com/weihaoranW/vchat/demo/unit/service"
	"github.com/weihaoranW/vchat/lib"
	"github.com/weihaoranW/vchat/lib/yetcd"
	"github.com/weihaoranW/vchat/lib/ylog"
)

var (
	// 每个微服务都不同，这里需要更改
	// todo
	msTag   = "api"
	host    = "0.0.0.0"
	port    = 9000
	regHost = "127.0.0.1"
	regPort = 9000
)

func init() {
	//------------ prepare modules----------
	//本步骤主要是装入系统必备的模块
	cfg, err := lib.InitModulesOfOptions(&lib.LoadOption{
		LoadEtcd:  true,
		LoadPg:    false,
		LoadRedis: false,
		LoadMongo: false,
		LoadMq:    false,
		LoadJwt:   true,
	})
	if err != nil {
		panic(err.Error())
	}

	//-------###############-----------------------------------
	//装入微服务配置,对于微服务开发人员，中需要写跌幅，写接口及实现
	ms := cfg.MicroService
	//assert yconfig.config
	if g.IsEmptyOr(ms.Host, ms.Tag, ms.RegHost) {
		panic("配置文件错误，必须有microService.host/regHost/tag")
	}

	if g.IsZeroOr(ms.Port, ms.RegPort) {
		panic("配置文件错误，必须有microService.port/regPort")
	}

	//微服務tag,用於註冊時標識/監聽端口/監聽主機
	msTag, port, regPort = ms.Tag, ms.Port, ms.RegPort
	//注冊用監聽/注册用主机
	host, regHost = "0.0.0.0", ms.RegHost
}

func main() {
	ylog.Info("微服务tag：", msTag)

	mux := http.NewServeMux()
	//--------handlers-----------------------------
	// 每一步：配置路由
	// HelloWorld handler
	mux.Handle("/HelloWorld", new(intf.HelloWorldHandler).HandlerLocal(new(service.HelloWorldImpl)))
	// userAdd,注意每个路由的来源
	mux.Handle("/UserAdd", new(intf.UserAddHandler).HandlerLocal(new(service.UserAddImpl)))
	// 每一个微服务都需要实现的方法，用于测试服务是否运行
	mux.Handle("/ping", http.HandlerFunc(new(Ping).handler))

	//-------register micro-service-----------------
	// 每二步:註冊微服務到etcd
	ylog.Info("正在向etcd注册微服务......")
	if err := yetcd.RegisterService(msTag, regHost, fmt.Sprint(regPort)); err != nil {
		ylog.Error("error", err)
		return
	}
	ylog.Info("注册微服务", msTag, " 成功")

	//--------start server -------------------------
	//用于显示服务器状态，用于测试
	fmt.Println(fmt.Sprint("监听:", host, ":", port))
	testStr := fmt.Sprintf(
		`测试：curl -X POST  -H 'Content-Type:application/json'  -d '{"S":"hello,weihaoran"}' %s:%d/HelloWorld`,
		host, port)
	fmt.Println(testStr)
	ylog.Info(fmt.Sprint("监听:", host, ":", port))

	addr := fmt.Sprint(host, ":", port)

	//-------------------------------------
	//  第三步：启动微服务
	golog.Fatal(http.ListenAndServe(addr, mux))
}

type Ping struct{}

func (r *Ping) handler(out http.ResponseWriter, in *http.Request) {
	_, _ = out.Write([]byte("pong........................"))
}
