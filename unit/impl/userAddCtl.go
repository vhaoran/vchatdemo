package ctl

import (
	"fmt"

	"vchatdemo/unit/intf"
	"github.com/vhaoran/vchat/lib/ykit"
)

type UserAddCtl struct {
}

func (r *UserAddCtl) Add(in *intf.UserAddRequest) (*ykit.Result, error) {
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
