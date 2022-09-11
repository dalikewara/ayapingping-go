package rflgo

var ErrCantSet = "can't set a value into `%s`, it might be not addressable or obtained by the use of unexported struct fields"
var ErrDestCantAddr = "the `dest %s` must be an addressable value"
var ErrDestInvalid = "the `dest %s` must be a valid reflect.Value"
var ErrValueKindNotMatch = "the kind of type from both value doesn't match (dest %s, source %s) (%s, %s)"
var ErrValueKindNotPointer = "the kind of type from both value must be a pointer (dest %s, source %s) (%s, %s)"
var ErrValueKindNotStruct = "the kind of type from both value must be a struct (dest %s, source %s) (%s, %s)"
var ErrValueKindNotSlice = "the kind of type from both value must be a slice (dest %s, source %s) (%s, %s)"
