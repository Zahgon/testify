//go:build testify_yaml_fail && !testify_yaml_custom && !testify_yaml_default

package yaml

import "errors"

var errNotImplemented = errors.New("YAML functions are not available (see https://pkg.go.dev/github.com/stretchr/testify/assert/yaml)")

func Unmarshal([]byte, interface{}) error { _ = "STUB: not implemented"; return nil }
