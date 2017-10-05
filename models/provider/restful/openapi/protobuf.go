package openapi

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/NYTimes/openapi2proto"
)

// improve, update ... ^^
func ConvertToProtobuf(swaggerFile string, prefixPath string, annotate bool) {
	_, err := os.Stat(swaggerFile)
	if err != nil {
		fmt.Printf("ConvertToProtobuf(...): File doesn't exist, skipping\n", swaggerFile)
		return
	}

	api, err := openapi2proto.LoadDefinition(swaggerFile)
	if err != nil {
		fmt.Println("unable to load spec: ", err)
		return
	}
	out, err := openapi2proto.GenerateProto(api, annotate)
	if err != nil {
		fmt.Println("unable to generate protobuf: ", err)
		return
	}
	if err := ioutil.WriteFile(fmt.Sprintf("%s/swagger.proto", prefixPath), out, 0755); err != nil {
		// check(err)
		fmt.Println("error:", err)
		return
	}
}
