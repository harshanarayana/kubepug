package kubepug

import (
	"github.com/rikatz/kubepug/pkg/schema"
)

// GetDeprecations returns the results of the comparision between the Input and the APIs
func GetDeprecations(d Deprecator) (result schema.Result) {
	result.DeprecatedAPIs = d.ListDeprecated()
	result.DeletedAPIs = d.ListDeleted()
	return result
}
