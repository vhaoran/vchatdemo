package intf

//for snippet用于标准返回值的微服务接口

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"github.com/go-kit/kit/auth/jwt"
	"github.com/go-kit/kit/endpoint"
	tran "github.com/go-kit/kit/transport/http"
	"github.com/vhaoran/vchat/lib/ykit"
	"github.com/vhaoran/vchat/lib/ylog"
	"log"
	"net/http"
)

const (
	CtxTest_HANDLER_PATH = "/CtxTest"
)

type (
	CtxTestService interface {
		Exec(in *CtxTestIn) (*ykit.Result, error)
	}

	//input data
	CtxTestIn struct {
		S string `json:"s"`
	}

	//output data
	//Result struct {
	//	Code int         `json:"code"`
	//	Msg  string      `json:"msg"`
	//	Data interface{} `json:"data"`
	//}

	// handler implements
	CtxTestHandler struct {
		base ykit.RootTran
	}
)

func (r *CtxTestHandler) MakeLocalEndpoint(svc CtxTestService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		fmt.Println("#######  Local Context ######  CtxTest ###########")
		ylog.Debug("jwt:", ctx.Value("Jwt"))
		spew.Dump(ctx)

		in := request.(*CtxTestIn)
		return svc.Exec(in)
	}
}

//个人实现,参数不能修改
func (r *CtxTestHandler) DecodeRequest(ctx context.Context, req *http.Request) (interface{}, error) {
	v, err := r.base.DecodeRequest(new(CtxTestIn), ctx, req)
	if err == nil {
		jwt := req.Header.Get("Jwt")
		ylog.Debug("header jwt: ", jwt)

		ylog.DebugDump("v:", v)
		s := ctx.Value("jwt")
		log.Println("----------", "ctx jwt :", s, "------------")

		log.Println("----------", "req.GetContext:", "------------")
		ylog.DebugDump("req.Context()", req.Context())

		//v.(*CtxTestIn).S = ":" + "abc"
	}
	return v, err
}

//个人实现,参数不能修改
func (r *CtxTestHandler) DecodeResponse(_ context.Context, res *http.Response) (interface{}, error) {
	var response ykit.Result
	if err := json.NewDecoder(res.Body).Decode(&response); err != nil {
		return nil, err
	}
	return response, nil
}

//handler for router，微服务本地接口，
func (r *CtxTestHandler) HandlerLocal(service CtxTestService,
	mid []endpoint.Middleware,
	options ...tran.ServerOption) *tran.Server {

	before := tran.ServerBefore(func(ctx context.Context, req *http.Request) context.Context {
		fmt.Println("-------HandlerLocal-----before ", req.Host)
		for k, v := range req.Header {
			ylog.Debug(k, ":", v)
		}
		//
		ylog.DebugDump("HandlerLocal ctx", ctx)

		return ctx
	})

	ep := r.MakeLocalEndpoint(service)
	for _, f := range mid {
		ep = f(ep)
	}

	handler := tran.NewServer(
		ep,
		r.DecodeRequest,
		r.base.EncodeResponse,
		before)
	//handler = loggingMiddleware()
	return handler
}

//sd,proxy实现,用于etcd自动服务发现时的handler
func (r *CtxTestHandler) HandlerSD(mid []endpoint.Middleware,
	options ...tran.ServerOption) *tran.Server {

	// 访问 request内容,丁当于Java中的拦截器
	before := tran.ServerBefore(func(ctx context.Context, req *http.Request) context.Context {
		jwt := req.Header.Get("Authorization")
		ylog.Debug("jwt: ", jwt)

		for k, v := range req.Header {
			ylog.Debug("header: ", k, ":", v)
		}

		fmt.Println("-------HandlerSD-----before host:", req.Host)
		return context.WithValue(ctx, "Jwt", jwt)
	})

	opts := append(options, before)
	return r.base.HandlerSD(
		context.Background(),
		MSTAG,
		"POST",
		CtxTest_HANDLER_PATH,
		r.DecodeRequest,
		r.DecodeResponse,
		mid,
		append(opts, tran.ServerBefore(jwt.HTTPToContext()))...)

}

func (r *CtxTestHandler) ProxySD() endpoint.Endpoint {
	return r.base.ProxyEndpointSD(
		context.Background(),
		MSTAG,
		"POST",
		CtxTest_HANDLER_PATH,
		r.DecodeRequest,
		r.DecodeResponse)
}
