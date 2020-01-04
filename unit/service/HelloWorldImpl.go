package service

import (
	"fmt"
	"time"

	"github.com/weihaoranW/vchat/demo/unit/intf"
)

type (
	HelloWorldImpl struct {
	}
)

func (h *HelloWorldImpl) Hello(in *intf.HelloWorldRequest) (string, error) {
	return fmt.Sprint("hello world,now is ", time.Now()), nil
}
