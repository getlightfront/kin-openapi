package jsoninfo

import (
	"encoding/json"
)

var Flatten bool

func MarshalRef(value string, otherwise interface{}) ([]byte, error) {
	if !Flatten && len(value) > 0 {
		return json.Marshal(&refProps{
			Ref: value,
		})
	}
	return json.Marshal(otherwise)
}

func UnmarshalRef(data []byte, destRef *string, destOtherwise interface{}) error {
	refProps := &refProps{}
	if err := json.Unmarshal(data, refProps); err == nil {
		ref := refProps.Ref
		if len(ref) > 0 {
			*destRef = ref
			return nil
		}
	}
	return json.Unmarshal(data, destOtherwise)
}

type refProps struct {
	Ref string `json:"$ref,omitempty"`
}
