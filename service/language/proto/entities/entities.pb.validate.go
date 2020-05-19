// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: service/language/proto/entities/entities.proto

package entities

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/golang/protobuf/ptypes"
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
	_ = ptypes.DynamicAny{}
)

// define the regex for a UUID once up-front
var _entities_uuidPattern = regexp.MustCompile("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$")

// Validate is disabled for Skill. This method will always return nil.
func (m *Skill) Validate() error {
	return nil
}

// SkillValidationError is the validation error returned by Skill.Validate if
// the designated constraints aren't met.
type SkillValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SkillValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SkillValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SkillValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SkillValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SkillValidationError) ErrorName() string { return "SkillValidationError" }

// Error satisfies the builtin error interface
func (e SkillValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSkill.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SkillValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SkillValidationError{}

// Validate is disabled for Word. This method will always return nil.
func (m *Word) Validate() error {
	return nil
}

// WordValidationError is the validation error returned by Word.Validate if the
// designated constraints aren't met.
type WordValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e WordValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e WordValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e WordValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e WordValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e WordValidationError) ErrorName() string { return "WordValidationError" }

// Error satisfies the builtin error interface
func (e WordValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sWord.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = WordValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = WordValidationError{}

// Validate is disabled for Language. This method will always return nil.
func (m *Language) Validate() error {
	return nil
}

// LanguageValidationError is the validation error returned by
// Language.Validate if the designated constraints aren't met.
type LanguageValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e LanguageValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e LanguageValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e LanguageValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e LanguageValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e LanguageValidationError) ErrorName() string { return "LanguageValidationError" }

// Error satisfies the builtin error interface
func (e LanguageValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sLanguage.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = LanguageValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = LanguageValidationError{}
