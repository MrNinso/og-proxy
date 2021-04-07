package types

import (
	j "github.com/json-iterator/go"
)

var json j.API = nil

func init() {
	if json == nil {
		json = j.Config{
			MarshalFloatWith6Digits: true,
			SortMapKeys:             false,
			UseNumber:               false,
			OnlyTaggedField:         true,
			CaseSensitive:           true,
		}.Froze()
	}
}


func GetJson() j.API {
	return json
} 