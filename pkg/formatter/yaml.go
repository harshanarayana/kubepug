package formatter

import (
	"github.com/rikatz/kubepug/pkg/schema"
	yamlencoder "gopkg.in/yaml.v3"
)

type yaml struct{}

func newYamlFormatter() Formatter {
	return &yaml{}
}

func (f *yaml) Output(results schema.Result) ([]byte, error) {
	y, err := yamlencoder.Marshal(results)
	if err != nil {
		return nil, err
	}
	return y, nil
}
