package service

import (
	"fmt"

	"github.com/weihaoranW/vchat/demo/unit/intf"
	"github.com/weihaoranW/vchat/lib/ykit"
)

type UserAddImpl struct {
}

func (r *UserAddImpl) Add(in *intf.UserAddRequest) (*ykit.Result, error) {
	// do some thing,add userInfo to db
	//
	fmt.Println("------", "input params", "-----------")
	fmt.Println(*in)
	fmt.Println("------", "end", "-----------")

	return &ykit.Result{
		Code: 200,
		Msg:  "操作成功",
		Data: nil,
	}, nil
}
