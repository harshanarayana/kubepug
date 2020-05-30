package schema

import "k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"

// ListObjects is a generic function that receives unstructured Kubernetes objects and convert to deprecatedItem
func ListObjects(items []unstructured.Unstructured) (deprecatedItems []DeprecatedItem) {
	for _, d := range items {
		name := d.GetName()
		namespace := d.GetNamespace()
		if namespace != "" {
			deprecatedItems = append(deprecatedItems, DeprecatedItem{Scope: "OBJECT", ObjectName: name, Namespace: namespace})
		} else {
			deprecatedItems = append(deprecatedItems, DeprecatedItem{Scope: "GLOBAL", ObjectName: name})
		}
	}
	return deprecatedItems
}
