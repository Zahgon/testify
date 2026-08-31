//go:build testify_yaml_custom && !testify_yaml_fail && !testify_yaml_default

package yaml

var Unmarshal func(in []byte, out interface{}) error
