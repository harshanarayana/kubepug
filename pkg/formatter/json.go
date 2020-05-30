package formatter

import (
	jsonencoding "encoding/json"

	"github.com/rikatz/kubepug/pkg/schema"
)

type json struct{}

func newJSONFormatter() Formatter {
	return &json{}
}

func (f *json) Output(results schema.Result) ([]byte, error) {
	j, err := jsonencoding.Marshal(results)
	if err != nil {
		return nil, err
	}
	return j, nil
}
