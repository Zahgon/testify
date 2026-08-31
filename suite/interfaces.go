package suite

import "testing"

type TestingSuite interface {
	T() *testing.T
	SetT(*testing.T)
	SetS(suite TestingSuite)
}

type SetupAllSuite interface {
	SetupSuite()
}

type SetupTestSuite interface {
	SetupTest()
}

type TearDownAllSuite interface {
	TearDownSuite()
}

type TearDownTestSuite interface {
	TearDownTest()
}

type BeforeTest interface {
	BeforeTest(suiteName, testName string)
}

type AfterTest interface {
	AfterTest(suiteName, testName string)
}

type WithStats interface {
	HandleStats(suiteName string, stats *SuiteInformation)
}

type SetupSubTest interface {
	SetupSubTest()
}

type TearDownSubTest interface {
	TearDownSubTest()
}
