package intf

//for snippet用于标准返回值的微服务接口
import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-kit/kit/endpoint"
	tran "github.com/go-kit/kit/transport/http"

	"github.com/weihaoranW/vchat/lib/ykit"
)

const (
	//外部定义的常量，每个微服务都不相同
	//MSTAG="api"
	P_UserAdd_HANDLER_PATH = "/UserAdd"
)

type (
	UserAddService interface {
		Add(in *UserAddRequest) (*ykit.Result, error)
	}
	//input data
	UserAddRequest struct {
		ID   string `json:"id"`
		Name string `json:"name,omitempty"`
		Age  int    `json:"age,omitempty"`
	}

	//output data
	//Result struct {
	//	Code int         `json:"code"`
	//	Msg  string      `json:"msg"`
	//	Data interface{} `json:"data"`
	//}

	// handler implements
	UserAddHandler struct {
		base ykit.RootTran
	}
)

//用作微服务的endPoint
func (r *UserAddHandler) MakeLocalEndpoint(svc UserAddService) endpoint.Endpoint {
	return func(_ context.Context, req interface{}) (interface{}, error) {
		//modify
		in := req.(*UserAddRequest)
		return svc.Add(in)
		//return ykit.Result{
		//	Code: 200,
		//	Msg:  "",
		//	Data: v,
		//}, err
	}
}

//个人实现,参数不能修改
func (r *UserAddHandler) DecodeRequest(ctx context.Context, req *http.Request) (interface{}, error) {
	return r.base.DecodeRequest(new(UserAddRequest), ctx, req)
}

//个人实现,参数不能修改
func (r *UserAddHandler) DecodeResponse(_ context.Context, res *http.Response) (interface{}, error) {
	var response ykit.Result
	if err := json.NewDecoder(res.Body).Decode(&response); err != nil {
		return nil, err
	}
	return response, nil
}

//handler for router，微服务本地接口，
func (r *UserAddHandler) HandlerLocal(service UserAddService) *tran.Server {
	ep := r.MakeLocalEndpoint(service)

	// 访问 request内容,丁当于Java中的拦截器
	before := tran.ServerBefore(func(ctx context.Context, req *http.Request) context.Context {
		fmt.Println("------------before host:", req.Host)
		return ctx
	})

	srv := tran.NewServer(
		ep,
		r.DecodeRequest,
		r.base.EncodeResponse,
		before)

	return srv
}

//sd,proxy实现,用于etcd自动服务发现时的handler
func (r *UserAddHandler) HandlerSD(mid ...endpoint.Middleware) *tran.Server {
	return r.base.HandlerSD(
		context.Background(),
		MSTAG,
		//外部定义的常量，每个微服务都不相同
		"POST",
		//具體的方法
		P_UserAdd_HANDLER_PATH,
		r.DecodeRequest,
		r.DecodeResponse,
		mid)
}

// for test
//测试proxy方式的实现,用於測試某一微服務的運行情況
func (r *UserAddHandler) HandlerProxyForTest() *tran.Server {
	ep := r.MakeProxyEndPointForTest(context.Background())
	return tran.NewServer(
		ep,
		r.DecodeRequest,
		r.base.EncodeResponse,
	)
}

// for test
//sd,proxy实现,调用 指定位置的endPoint
func (r *UserAddHandler) MakeProxyEndPointForTest(
	ctx context.Context) endpoint.Endpoint {
	//modify
	return r.base.MakeProxyEndPoint(
		//此为被调用的微服务的(host:port),
		"localhost:9001",
		"POST",
		P_UserAdd_HANDLER_PATH,
		r.DecodeResponse,
		ctx)
}
