package db

import (
	"github.com/coreos/etcd/clientv3"
	"github.com/soyking/e3ch"
)

/*
	github.com/soyking/e3ch
*/

func e3chExample() {

	// initial etcd v3 client
	e3Clt, err := clientv3.New(clientv3.Config{Endpoints: []string{"127.0.0.1:2379"}})
	if err != nil {
		panic(err)
	}

	// new e3ch client with namespace(rootKey)
	clt, err := client.New(e3Clt, "my_space")
	if err != nil {
		panic(err)
	}

	// set the rootKey as directory
	err = clt.FormatRootKey()
	if err != nil {
		panic(err)
	}

	clt.CreateDir("/dir1")
	clt.Create("/dir1/key1", "")
	clt.Create("/dir", "")
	clt.Put("/dir1/key1", "value1")
	clt.Get("/dir1/key1")
	clt.List("/dir1")
	clt.Delete("/dir")

}
