package pointer

// StringPtr returns a pointer to the passed string.
func StringPtr(s string) *string {
	return &s
}

// Int32Ptr returns a pointer to an int32
func Int32Ptr(i int32) *int32 {
	return &i
}
