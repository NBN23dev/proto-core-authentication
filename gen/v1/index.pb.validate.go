// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: v1/index.proto

package _go

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"sort"
	"strings"
	"time"
	"unicode/utf8"

	"google.golang.org/protobuf/types/known/anypb"
)

// ensure the imports are used
var (
	_ = bytes.MinRead
	_ = errors.New("")
	_ = fmt.Print
	_ = utf8.UTFMax
	_ = (*regexp.Regexp)(nil)
	_ = (*strings.Reader)(nil)
	_ = net.IPv4len
	_ = time.Duration(0)
	_ = (*url.URL)(nil)
	_ = (*mail.Address)(nil)
	_ = anypb.Any{}
	_ = sort.Sort
)

// Validate checks the field values on JWK with the rules defined in the proto
// definition for this message. If any rules are violated, the first error
// encountered is returned, or nil if there are no violations.
func (m *JWK) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on JWK with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in JWKMultiError, or nil if none found.
func (m *JWK) ValidateAll() error {
	return m.validate(true)
}

func (m *JWK) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Kid

	// no validation rules for Alg

	// no validation rules for Use

	// no validation rules for Kty

	// no validation rules for E

	// no validation rules for N

	if len(errors) > 0 {
		return JWKMultiError(errors)
	}

	return nil
}

// JWKMultiError is an error wrapping multiple validation errors returned by
// JWK.ValidateAll() if the designated constraints aren't met.
type JWKMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m JWKMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m JWKMultiError) AllErrors() []error { return m }

// JWKValidationError is the validation error returned by JWK.Validate if the
// designated constraints aren't met.
type JWKValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e JWKValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e JWKValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e JWKValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e JWKValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e JWKValidationError) ErrorName() string { return "JWKValidationError" }

// Error satisfies the builtin error interface
func (e JWKValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sJWK.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = JWKValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = JWKValidationError{}

// Validate checks the field values on PublicKeysRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *PublicKeysRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on PublicKeysRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// PublicKeysRequestMultiError, or nil if none found.
func (m *PublicKeysRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *PublicKeysRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return PublicKeysRequestMultiError(errors)
	}

	return nil
}

// PublicKeysRequestMultiError is an error wrapping multiple validation errors
// returned by PublicKeysRequest.ValidateAll() if the designated constraints
// aren't met.
type PublicKeysRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m PublicKeysRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m PublicKeysRequestMultiError) AllErrors() []error { return m }

// PublicKeysRequestValidationError is the validation error returned by
// PublicKeysRequest.Validate if the designated constraints aren't met.
type PublicKeysRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PublicKeysRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PublicKeysRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PublicKeysRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PublicKeysRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PublicKeysRequestValidationError) ErrorName() string {
	return "PublicKeysRequestValidationError"
}

// Error satisfies the builtin error interface
func (e PublicKeysRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPublicKeysRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PublicKeysRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PublicKeysRequestValidationError{}

// Validate checks the field values on PublicKeysResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *PublicKeysResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on PublicKeysResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// PublicKeysResponseMultiError, or nil if none found.
func (m *PublicKeysResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *PublicKeysResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetKeys() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, PublicKeysResponseValidationError{
						field:  fmt.Sprintf("Keys[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, PublicKeysResponseValidationError{
						field:  fmt.Sprintf("Keys[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return PublicKeysResponseValidationError{
					field:  fmt.Sprintf("Keys[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return PublicKeysResponseMultiError(errors)
	}

	return nil
}

// PublicKeysResponseMultiError is an error wrapping multiple validation errors
// returned by PublicKeysResponse.ValidateAll() if the designated constraints
// aren't met.
type PublicKeysResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m PublicKeysResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m PublicKeysResponseMultiError) AllErrors() []error { return m }

// PublicKeysResponseValidationError is the validation error returned by
// PublicKeysResponse.Validate if the designated constraints aren't met.
type PublicKeysResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PublicKeysResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PublicKeysResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PublicKeysResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PublicKeysResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PublicKeysResponseValidationError) ErrorName() string {
	return "PublicKeysResponseValidationError"
}

// Error satisfies the builtin error interface
func (e PublicKeysResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPublicKeysResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PublicKeysResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PublicKeysResponseValidationError{}
