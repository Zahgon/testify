//go:build !testify_yaml_fail && !testify_yaml_custom

package yaml

func Unmarshal(in []byte, out interface{}) error { _ = "STUB: not implemented"; return nil }
