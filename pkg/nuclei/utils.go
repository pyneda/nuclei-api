package nuclei

import (
	"fmt"
	"github.com/projectdiscovery/goflags"
	"github.com/projectdiscovery/nuclei/v2/pkg/output"
	"reflect"
)

func printFields(resultEvent output.ResultEvent) {
	v := reflect.ValueOf(resultEvent)

	for i := 0; i < v.NumField(); i++ {
		fmt.Printf("%v: %v\n", v.Type().Field(i).Name, v.Field(i).Interface())
	}
}

func assignIfNotEmpty(dst *goflags.StringSlice, src *[]string) {
	if len(*src) > 0 {
		*dst = *src
	}
}
