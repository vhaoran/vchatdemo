package intf

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/davecgh/go-spew/spew"
	"github.com/go-kit/kit/endpoint"
	tran "github.com/go-kit/kit/transport/http"

	"github.com/weihaoranW/vchat/lib/ykit"
)

const (
	//todo
	HelloWorld_HANDLER_PATH = "/HelloWorld"
)

type (
	HelloWorldService interface {
		//todo
		Hello(in *HelloWorldRequest) (string, error)
	}

	//input data
	//todo
	HelloWorldRequest struct {
		S string `json:"s"`
	}

	//output data
	//Result struct {
	//	Code int         `json:"code"`
	//	Msg  string      `json:"msg"`
	//	Data interface{} `json:"data"`
	//}

	// handler implements
	HelloWorldHandler struct {
		base ykit.RootTran
	}
)

func (r *HelloWorldHandler) MakeLocalEndpoint(svc HelloWorldService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		fmt.Println("#############  hello ###########")
		spew.Dump(ctx)

		//todo
		in := request.(*HelloWorldRequest)
		ret, err := svc.Hello(in)
		if err != nil {
			return &ykit.Result{
				Code: 500,
				Msg:  "服务器内部错误",
				Data: nil,
			}, err
		}
		//todo
		return &ykit.Result{
			Code: 200,
			Msg:  "ok",
			Data: ret,
		}, nil
	}
}

//个人实现,参数不能修改
func (r *HelloWorldHandler) DecodeRequest(ctx context.Context, req *http.Request) (interface{}, error) {
	return r.base.DecodeRequest(new(HelloWorldRequest), ctx, req)
}

//个人实现,参数不能修改
func (r *HelloWorldHandler) DecodeResponse(_ context.Context, res *http.Response) (interface{}, error) {
	var response ykit.Result
	if err := json.NewDecoder(res.Body).Decode(&response); err != nil {
		return nil, err
	}
	return response, nil
}

//handler for router，微服务本地接口，
func (r *HelloWorldHandler) HandlerLocal(service HelloWorldService) *tran.Server {
	endpoint := r.MakeLocalEndpoint(service)
	handler := tran.NewServer(
		endpoint,
		r.DecodeRequest,
		r.base.EncodeResponse,
	)
	//handler = loggingMiddleware()
	return handler
}

//sd,proxy实现,用于etcd自动服务发现时的handler
func (r *HelloWorldHandler) HandlerSD(mid ...endpoint.Middleware) *tran.Server {
	return r.base.HandlerSD(
		context.Background(),
		MSTAG,
		//todo
		"POST",
		HelloWorld_HANDLER_PATH,
		r.DecodeRequest,
		r.DecodeResponse,
		mid)
}

// for test
//测试proxy方式的实现,用於測試某一微服務的運行情況
func (r *HelloWorldHandler) HandlerProxyForTest() *tran.Server {
	e := r.MakeProxyEndPointForTest(context.Background())
	handler := tran.NewServer(
		e,
		r.DecodeRequest,
		r.base.EncodeResponse,
	)

	return handler
}

// for test
//sd,proxy实现,调用 指定位置的endPoint
func (r *HelloWorldHandler) MakeProxyEndPointForTest(
	ctx context.Context) endpoint.Endpoint {
	//modify
	return r.base.MakeProxyEndPoint(
		//此为被调用的微服务的(host:port),
		"localhost:9001",
		"POST",
		HelloWorld_HANDLER_PATH,
		r.DecodeResponse,
		ctx)
}
