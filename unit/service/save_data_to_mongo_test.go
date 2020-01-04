package service

import (
	"fmt"
	"log"
	"testing"

	"github.com/weihaoranW/vchat/lib"
	"github.com/weihaoranW/vchat/lib/ymongo"
)

type MongoHello struct {
	ID    int
	CName string
	Age   int
}

func Test_insert_one(t *testing.T) {
	// load config
	opt := &lib.LoadOption{
		LoadMicroService: false,
		LoadEtcd:         false,
		LoadPg:           false,
		LoadRedis:        false,

		//-----------attention here------------
		LoadMongo: true,
		//-----------attention here------------

		LoadMq:  false,
		LoadJwt: false,
	}
	if _, err := lib.InitModulesOfOptions(opt); err != nil {
		panic(err.Error())
	}

	//
	bean := &MongoHello{
		ID:    1,
		CName: "hello world",
		Age:   88,
	}

	r, err := ymongo.XMongo.DoInsertOne("test", "myTest", bean)
	log.Println(r)
	fmt.Println("------", "", "-----------")
	log.Println(err)
}
