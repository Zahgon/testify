package suite

import "time"

type SuiteInformation struct {
	Start, End time.Time
	TestStats  map[string]*TestInformation
}

type TestInformation struct {
	TestName   string
	Start, End time.Time
	Passed     bool
}

func newSuiteInformation() *SuiteInformation { _ = "STUB: not implemented"; return nil }

func (s *SuiteInformation) start(testName string) { _ = "STUB: not implemented"; return }

func (s *SuiteInformation) end(testName string, passed bool) { _ = "STUB: not implemented"; return }

func (s *SuiteInformation) Passed() bool { _ = "STUB: not implemented"; return false }
