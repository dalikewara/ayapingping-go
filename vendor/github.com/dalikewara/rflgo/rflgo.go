package rflgo

import (
	"errors"
	"fmt"
	"reflect"
)

// Compose composes new value into `dest` based on the `source` value.
// Both must have the same kind of type. For example, if `dest int` then
// the source must be `source int`.
func Compose(dest, source interface{}) error {
	destVal := reflect.Indirect(reflect.ValueOf(dest))
	sourceVal := reflect.ValueOf(source)
	if err := Set(destVal, sourceVal); err != nil {
		return err
	}
	return nil
}

// Set sets new value into `dest` based on the `source` value.
// Both must have the same kind of type. For example, if `dest int` then
// the source must be `source int`. The `dest` must be an addressable value and valid reflect.Value.
func Set(dest, source reflect.Value) error {
	if skipSource(source) {
		return nil
	}
	if err := checkDest(dest); err != nil {
		return err
	}
	if dest.Kind() != source.Kind() {
		return errors.New(fmt.Sprintf(ErrValueKindNotMatch, dest.Type().Kind(), source.Type().Kind(), dest.Type(), source.Type()))
	}
	switch dest.Kind() {
	case reflect.Pointer:
		if err := SetPointer(dest, source); err != nil {
			return err
		}
	case reflect.Struct:
		if err := SetStruct(dest, source); err != nil {
			return err
		}
	case reflect.Slice:
		if err := SetSlice(dest, source); err != nil {
			return err
		}
	default:
		dest.Set(source)
	}
	return nil
}

// SetSlice sets new slice value into `dest` based on the `source` value.
// Both kind of type must be a slice and the `dest` must be an addressable value and valid reflect.Value.
func SetSlice(dest, source reflect.Value) error {
	if skipSource(source) {
		return nil
	}
	if err := checkDest(dest); err != nil {
		return err
	}
	if dest.Kind() != reflect.Slice || source.Kind() != reflect.Slice {
		return errors.New(fmt.Sprintf(ErrValueKindNotSlice, dest.Kind(), source.Kind(), dest.Type(), source.Type()))
	}
	tmp := reflect.MakeSlice(reflect.SliceOf(dest.Type().Elem()), source.Len(), source.Cap())
	for i := 0; i < tmp.Len(); i++ {
		if err := Set(tmp.Index(i), source.Index(i)); err != nil {
			return err
		}
	}
	dest.Set(tmp)
	return nil
}

// SetStruct sets new struct value into `dest` based on the `source` value.
// Both kind of type must be a struct and the `dest` must be an addressable value and valid reflect.Value.
func SetStruct(dest, source reflect.Value) error {
	if skipSource(source) {
		return nil
	}
	if err := checkDest(dest); err != nil {
		return err
	}
	if dest.Kind() != reflect.Struct || source.Kind() != reflect.Struct {
		return errors.New(fmt.Sprintf(ErrValueKindNotStruct, dest.Kind(), source.Kind(), dest.Type(), source.Type()))
	}
	for i := 0; i < dest.NumField(); i++ {
		fieldName := dest.Type().Field(i).Name
		if err := Set(dest.FieldByName(fieldName), source.FieldByName(fieldName)); err != nil {
			return err
		}
	}
	return nil
}

// SetPointer sets new pointer value into `dest` based on the `source` value.
// Both kind of type must be a pointer and the `dest` must be an addressable value and valid reflect.Value.
func SetPointer(dest, source reflect.Value) error {
	if skipSource(source) {
		return nil
	}
	if err := checkDest(dest); err != nil {
		return err
	}
	if dest.Kind() != reflect.Pointer || source.Kind() != reflect.Pointer {
		return errors.New(fmt.Sprintf(ErrValueKindNotPointer, dest.Kind(), source.Kind(), dest.Type(), source.Type()))
	}
	dest.Set(reflect.New(dest.Type().Elem()))
	if err := Set(dest.Elem(), source.Elem()); err != nil {
		return err
	}
	return nil
}

// checkDest checks `dest` value. It must be an addressable value and valid reflect.Value.
func checkDest(dest reflect.Value) error {
	if !dest.IsValid() {
		return errors.New(fmt.Sprintf(ErrDestInvalid, dest))
	}
	if !dest.CanAddr() {
		return errors.New(fmt.Sprintf(ErrDestCantAddr, dest.Type()))
	}
	if !dest.CanSet() {
		return errors.New(fmt.Sprintf(ErrCantSet, dest.Type().Name()))
	}
	return nil
}

// skipSource returns true if the `source` value is nil, empty, or not valid reflect.Value.
func skipSource(source reflect.Value) bool {
	if source.Kind() == reflect.Pointer {
		if !source.Elem().IsValid() || source.Elem().IsZero() {
			return true
		}
	}
	if !source.IsValid() || source.IsZero() {
		return true
	}
	return false
}
