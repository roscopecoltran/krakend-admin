package main

import (
	// "bufio"
	"fmt"
	"io/ioutil"
	"log"
	// "os"

	"github.com/NYTimes/openapi2proto"
)

func openapi2Protobuf(swaggerFile string, prefixPath string, annotate bool) {
	api, err := openapi2proto.LoadDefinition(swaggerFile)
	if err != nil {
		log.Fatal("unable to load spec: ", err)
	}
	out, err := openapi2proto.GenerateProto(api, annotate)
	if err != nil {
		log.Fatal("unable to generate protobuf: ", err)
	}
	if err := ioutil.WriteFile(fmt.Sprintf("%s/swagger.proto", prefixPath), out, 0755); err != nil {
		// check(err)
		fmt.Println("error:", err)
	}
}
