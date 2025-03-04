package testhelper

import "time"

// StringPointer is a helper to get a *string from a string. It is in the tests
// package because we normally don't want to deal with pointers to basic types,
// but it's useful in some tests.
func StringPointer(s string) *string {
	return &s
}

// IntPointer is a helper to get a *int from a int. It is in the tests package
// because we normally don't want to deal with pointers to basic types, but it's
// useful in some tests.
func IntPointer(i int) *int {
	return &i
}

// FloatPointer is a helper to get a *float64 from a float64. It is in the tests
// package because we normally don't want to deal with pointers to basic types,
// but it's useful in some tests.
func FloatPointer(f float64) *float64 {
	return &f
}

// BoolPointer is a helper to get a *bool from a bool. It is in the tests package
// because we normally don't want to deal with pointers to basic types, but it's
// useful in some tests.
func BoolPointer(b bool) *bool {
	return &b
}

// TimePointer is a helper to get a *time.Time from a time.Time. It is in the tests
// package because we normally don't want to deal with pointers to basic types,
// but it's useful in some tests.
func TimePointer(t time.Time) *time.Time {
	return &t
}
