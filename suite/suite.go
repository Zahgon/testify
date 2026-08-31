package suite

import (
	"flag"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var matchMethod = flag.String("testify.m", "", "regular expression to select tests of the testify suite to run")

type Suite struct {
	*assert.Assertions

	mu      sync.RWMutex
	require *require.Assertions
	t       *testing.T

	s TestingSuite
}

func (suite *Suite) T() *testing.T { _ = "STUB: not implemented"; return nil }

func (suite *Suite) SetT(t *testing.T) { _ = "STUB: not implemented"; return }

func (suite *Suite) SetS(s TestingSuite) { _ = "STUB: not implemented"; return }

func (suite *Suite) Require() *require.Assertions { _ = "STUB: not implemented"; return nil }

func (suite *Suite) Assert() *assert.Assertions { _ = "STUB: not implemented"; return nil }

func recoverAndFailOnPanic(t *testing.T) { _ = "STUB: not implemented"; return }

func failOnPanic(t *testing.T, r interface{}) { _ = "STUB: not implemented"; return }

func (suite *Suite) Run(name string, subtest func()) bool { _ = "STUB: not implemented"; return false }

type test = struct {
	name string
	run  func(t *testing.T)
}

func Run(t *testing.T, suite TestingSuite) { _ = "STUB: not implemented"; return }

func runTests(t *testing.T, tests []test) { _ = "STUB: not implemented"; return }
