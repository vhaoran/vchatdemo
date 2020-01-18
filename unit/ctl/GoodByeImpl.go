package ctl

import (
	"vchatdemo/unit/intf"
)

type GoodByeImpl struct {
}

func (r *GoodByeImpl) Exec(in *intf.GoodByeRequest) (string, error) {
	s := in.S
	return s + " byte bye.....", nil
}
