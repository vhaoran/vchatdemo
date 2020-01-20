package ctl

import (
	"github.com/vhaoran/vchat/lib/ykit"
	"vchatdemo/unit/intf"
)

type CtxTestImpl struct {
}

func (r *CtxTestImpl) Exec(in *intf.CtxTestIn) (*ykit.Result, error) {
	return ykit.ROK(in.S + "hello"), nil
}
