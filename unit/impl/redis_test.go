package ctl

import (
	"fmt"
	"log"
	"testing"
	"time"

	"github.com/vhaoran/vchat/lib"
	"github.com/vhaoran/vchat/lib/yredis"
)

func Test_redis_set(t *testing.T) {
	// load config
	opt := &lib.LoadOption{
		LoadMicroService: false,
		LoadEtcd:         false,
		LoadPg:           false,
		//-----------attention here------------
		LoadRedis: true,
		//-----------attention here------------

		LoadMongo: false,

		LoadMq:  false,
		LoadJwt: false,
	}

	_, err := lib.InitModulesOfOptions(opt)
	if err != nil {
		log.Println(err)
		return
	}

	key := "hello/key"
	ret, err := yredis.X.Set(key, "hello_value", time.Second*1000).Result()
	fmt.Println("---ret---", ret, "-----------")
	fmt.Println("---err---", err, "-----------")

	fmt.Println("------", "demo get", "-----------")
	str, err := yredis.X.Get(key).Result()
	fmt.Println("----err--", err, "-----------")
	log.Println("key value:", str)
}
