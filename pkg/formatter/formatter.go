package formatter

import (
	"github.com/rikatz/kubepug/pkg/schema"
)

//Formatter defines the behavior for a Formatter
type Formatter interface {
	Output(results schema.Result) ([]byte, error)
}

//NewFormatter returns a new instance of formatter
func NewFormatter(t string) Formatter {
	switch t {
	case "stdout":
		return newSTDOUTFormatter()
	case "plain":
		return newPlainFormatter()
	case "json":
		return newJSONFormatter()
	case "yaml":
		return newYamlFormatter()
	default:
		return newSTDOUTFormatter()
	}
}
