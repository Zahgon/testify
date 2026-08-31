package assert

import (
	"reflect"
	"time"
)

type CompareType = compareResult

type compareResult int

const (
	compareLess compareResult = iota - 1
	compareEqual
	compareGreater
)

var (
	intType   = reflect.TypeOf(int(1))
	int8Type  = reflect.TypeOf(int8(1))
	int16Type = reflect.TypeOf(int16(1))
	int32Type = reflect.TypeOf(int32(1))
	int64Type = reflect.TypeOf(int64(1))

	uintType   = reflect.TypeOf(uint(1))
	uint8Type  = reflect.TypeOf(uint8(1))
	uint16Type = reflect.TypeOf(uint16(1))
	uint32Type = reflect.TypeOf(uint32(1))
	uint64Type = reflect.TypeOf(uint64(1))

	uintptrType = reflect.TypeOf(uintptr(1))

	float32Type = reflect.TypeOf(float32(1))
	float64Type = reflect.TypeOf(float64(1))

	stringType = reflect.TypeOf("")

	timeType  = reflect.TypeOf(time.Time{})
	bytesType = reflect.TypeOf([]byte{})
)

func compare(obj1, obj2 interface{}, kind reflect.Kind) (compareResult, bool) {
	_ = "STUB: not implemented"
	return *new(compareResult), false
}

func Greater(t TestingT, e1 interface{}, e2 interface{}, msgAndArgs ...interface{}) bool {
	_ = "STUB: not implemented"
	return false
}

func GreaterOrEqual(t TestingT, e1 interface{}, e2 interface{}, msgAndArgs ...interface{}) bool {
	_ = "STUB: not implemented"
	return false
}

func Less(t TestingT, e1 interface{}, e2 interface{}, msgAndArgs ...interface{}) bool {
	_ = "STUB: not implemented"
	return false
}

func LessOrEqual(t TestingT, e1 interface{}, e2 interface{}, msgAndArgs ...interface{}) bool {
	_ = "STUB: not implemented"
	return false
}

func Positive(t TestingT, e interface{}, msgAndArgs ...interface{}) bool {
	_ = "STUB: not implemented"
	return false
}

func Negative(t TestingT, e interface{}, msgAndArgs ...interface{}) bool {
	_ = "STUB: not implemented"
	return false
}

func compareTwoValues(t TestingT, e1 interface{}, e2 interface{}, allowedComparesResults []compareResult, failMessage string, msgAndArgs ...interface{}) bool {
	_ = "STUB: not implemented"
	return false
}

func containsValue(values []compareResult, value compareResult) bool {
	_ = "STUB: not implemented"
	return false
}
