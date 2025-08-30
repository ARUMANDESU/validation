// Copyright 2016 Qiang Xue. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package validation

import (
	"database/sql"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLength(t *testing.T) {
	var v *string
	tests := []struct {
		tag      string
		min, max int
		value    interface{}
		err      string
	}{
		{"t1", 2, 4, "abc", ""},
		{"t2", 2, 4, "", ""},
		{"t3", 2, 4, "abcdf", "the length must be between 2 and 4"},
		{"t4", 0, 4, "ab", ""},
		{"t5", 0, 4, "abcde", "the length must be no more than 4"},
		{"t6", 2, 0, "ab", ""},
		{"t7", 2, 0, "a", "the length must be no less than 2"},
		{"t8", 2, 0, v, ""},
		{"t9", 2, 0, 123, "cannot get the length of int"},
		{"t10", 2, 4, sql.NullString{String: "abc", Valid: true}, ""},
		{"t11", 2, 4, sql.NullString{String: "", Valid: true}, ""},
		{"t12", 2, 4, &sql.NullString{String: "abc", Valid: true}, ""},
		{"t13", 2, 2, "abcdf", "the length must be exactly 2"},
		{"t14", 2, 2, "ab", ""},
		{"t15", 0, 0, "", ""},
		{"t16", 0, 0, "ab", "the value must be empty"},
		{"t17", 2, 4, []int{1, 2, 3}, ""},
		{"t18", 2, 4, []int{1}, "the length must be between 2 and 4"},
		{"t19", 2, 4, []int{1, 2, 3, 4, 5}, "the length must be between 2 and 4"},
		{"t20", 2, 4, map[string]int{"a": 1, "b": 2}, ""},
		{"t21", 2, 4, map[string]int{"a": 1}, "the length must be between 2 and 4"},
		{"t22", 2, 4, map[string]int{"a": 1, "b": 2, "c": 3, "d": 4, "e": 5}, "the length must be between 2 and 4"},
		{"t23", 2, 4, [3]int{1, 2, 3}, ""},
		{"t24", 2, 4, [1]int{1}, "the length must be between 2 and 4"},
		{"t25", 2, 4, [5]int{1, 2, 3, 4, 5}, "the length must be between 2 and 4"},
	}

	for _, test := range tests {
		r := Length(test.min, test.max)
		err := r.Validate(test.value)
		assertError(t, test.err, err, test.tag)
	}
}

func TestCount(t *testing.T) {
	var v *string
	tests := []struct {
		tag      string
		min, max int
		value    interface{}
		err      string
	}{
		{"t1", 2, 4, "abc", ""},
		{"t2", 2, 4, "", ""},
		{"t3", 2, 4, "abcdf", "the count must be between 2 and 4"},
		{"t4", 0, 4, "ab", ""},
		{"t5", 0, 4, "abcde", "the count must be no more than 4"},
		{"t6", 2, 0, "ab", ""},
		{"t7", 2, 0, "a", "the count must be no less than 2"},
		{"t8", 2, 0, v, ""},
		{"t10", 2, 4, sql.NullString{String: "abc", Valid: true}, ""},
		{"t11", 2, 4, sql.NullString{String: "", Valid: true}, ""},
		{"t12", 2, 4, &sql.NullString{String: "abc", Valid: true}, ""},
		{"t13", 2, 2, "abcdf", "the count must be exactly 2"},
		{"t14", 2, 2, "ab", ""},
		{"t15", 0, 0, "", ""},
		{"t16", 0, 0, "ab", "the value must be empty"},
		{"t17", 2, 4, []string{"a", "b", "c"}, ""},
		{"t18", 2, 4, []string{"a"}, "the count must be between 2 and 4"},
		{"t19", 2, 4, map[string]int{"a": 1, "b": 2}, ""},
		{"t20", 2, 4, map[string]int{"a": 1}, "the count must be between 2 and 4"},
		{"t21", 2, 4, [3]int{1, 2, 3}, ""},
		{"t22", 2, 4, [1]int{1}, "the count must be between 2 and 4"},
		{"t23", 2, 4, []int{}, ""},
		{"t24", 0, 4, []int{}, ""},
		{"t25", 2, 0, []int{}, ""},
	}

	for _, test := range tests {
		r := Count(test.min, test.max)
		err := r.Validate(test.value)
		assertError(t, test.err, err, test.tag)
	}
}

func TestRuneLength(t *testing.T) {
	var v *string
	tests := []struct {
		tag      string
		min, max int
		value    interface{}
		err      string
	}{
		{"t1", 2, 4, "abc", ""},
		{"t1.1", 2, 3, "💥💥", ""},
		{"t1.2", 2, 3, "💥💥💥", ""},
		{"t1.3", 2, 3, "💥", "the length must be between 2 and 3"},
		{"t1.4", 2, 3, "💥💥💥💥", "the length must be between 2 and 3"},
		{"t2", 2, 4, "", ""},
		{"t3", 2, 4, "abcdf", "the length must be between 2 and 4"},
		{"t4", 0, 4, "ab", ""},
		{"t5", 0, 4, "abcde", "the length must be no more than 4"},
		{"t6", 2, 0, "ab", ""},
		{"t7", 2, 0, "a", "the length must be no less than 2"},
		{"t8", 2, 0, v, ""},
		{"t9", 2, 0, 123, "cannot get the length of int"},
		{"t10", 2, 4, sql.NullString{String: "abc", Valid: true}, ""},
		{"t11", 2, 4, sql.NullString{String: "", Valid: true}, ""},
		{"t12", 2, 4, &sql.NullString{String: "abc", Valid: true}, ""},
		{"t13", 2, 3, &sql.NullString{String: "💥💥", Valid: true}, ""},
		{"t14", 2, 3, &sql.NullString{String: "💥", Valid: true}, "the length must be between 2 and 3"},
	}

	for _, test := range tests {
		r := RuneLength(test.min, test.max)
		err := r.Validate(test.value)
		assertError(t, test.err, err, test.tag)
	}
}

func Test_LengthRule_Error(t *testing.T) {
	r := Length(10, 20)
	assert.Equal(t, "the length must be between 10 and 20", r.Validate("abc").Error())

	r = Length(0, 20)
	assert.Equal(t, "the length must be no more than 20", r.Validate(make([]string, 21)).Error())

	r = Length(10, 0)
	assert.Equal(t, "the length must be no less than 10", r.Validate([9]string{}).Error())

	r = Length(0, 0)
	assert.Equal(t, "validation_length_empty_required", r.err.Code())

	r = r.Error("123")
	assert.Equal(t, "123", r.err.Message())
}

func TestLengthRule_ErrorObject(t *testing.T) {
	r := Length(10, 20)
	err := NewError("code", "abc")
	r = r.ErrorObject(err)

	assert.Equal(t, err, r.err)
	assert.Equal(t, err.Code(), r.err.Code())
	assert.Equal(t, err.Message(), r.err.Message())
}
