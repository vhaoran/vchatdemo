package ctl

import (
	"fmt"
	"time"

	"vchatdemo/unit/intf"
)

type (
	HelloWorldCtl struct {
	}
)

func (h *HelloWorldCtl) Hello(in *intf.HelloWorldRequest) (string, error) {
	return fmt.Sprint("hello world,now is ", time.Now()), nil
}
