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

// Validate checks the field values on Skill with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Skill) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Id

	if v, ok := interface{}(m.GetCreatedAt()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return SkillValidationError{
				field:  "CreatedAt",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetUpdatedAt()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return SkillValidationError{
				field:  "UpdatedAt",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetDeletedAt()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return SkillValidationError{
				field:  "DeletedAt",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for Disabled

	// no validation rules for LessonNumber

	// no validation rules for Description

	// no validation rules for Locked

	// no validation rules for Type

	// no validation rules for Title

	// no validation rules for UrlTitle

	// no validation rules for Category

	// no validation rules for Index

	// no validation rules for LanguageId

	if v, ok := interface{}(m.GetLanguage()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return SkillValidationError{
				field:  "Language",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	for idx, item := range m.GetWords() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return SkillValidationError{
					field:  fmt.Sprintf("Words[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

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

// Validate checks the field values on Word with the rules defined in the proto
// definition for this message. If any rules are violated, an error is returned.
func (m *Word) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Id

	if v, ok := interface{}(m.GetCreatedAt()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return WordValidationError{
				field:  "CreatedAt",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetUpdatedAt()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return WordValidationError{
				field:  "UpdatedAt",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetDeletedAt()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return WordValidationError{
				field:  "DeletedAt",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for Content

	// no validation rules for AudioSrc

	// no validation rules for LanguageId

	if v, ok := interface{}(m.GetLanguage()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return WordValidationError{
				field:  "Language",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

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

// Validate checks the field values on Language with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Language) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Id

	if v, ok := interface{}(m.GetCreatedAt()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return LanguageValidationError{
				field:  "CreatedAt",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetUpdatedAt()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return LanguageValidationError{
				field:  "UpdatedAt",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetDeletedAt()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return LanguageValidationError{
				field:  "DeletedAt",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for Name

	// no validation rules for Abbreviation

	// no validation rules for FlagSrc

	for idx, item := range m.GetWords() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return LanguageValidationError{
					field:  fmt.Sprintf("Words[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

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
