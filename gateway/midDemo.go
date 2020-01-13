package main

import (
	"context"
	"fmt"
	golog "log"

	"github.com/go-kit/kit/endpoint"
)

func Middleware1(next endpoint.Endpoint) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		fmt.Println("--------middle ware 1-----------------------")
		//spew.Dump(request)
		golog.Println("----------", "", "------------")
		//spew.Dump(ctx)
		fmt.Println("middle 1 is call ---pre")

		defer fmt.Println("middle 1 is call ---post")
		return next(ctx, request)
	}
}
