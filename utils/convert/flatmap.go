package convert

import (
	"encoding/json"
	"fmt"
	"sort"

	"github.com/dimiro1/flatmap"
)

func GetFlatMap(data []byte, debug bool) ([]string, error) {
	var mp map[string]interface{}
	var ks []string
	if err := json.Unmarshal(data, &mp); err != nil {
		return ks, err
	}
	fm, err := flatmap.Flatten(mp)
	if err != nil {
		return ks, err
	}
	for k := range fm {
		ks = append(ks, k)
	}
	sort.Strings(ks)
	if debug {
		for _, k := range ks {
			fmt.Println(k, ":", fm[k])
		}
	}
	return ks, nil

}
