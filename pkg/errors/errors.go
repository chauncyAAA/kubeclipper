package errors

import (
	"fmt"
	"strings"
)

var _ error = &StatusError{}

type StatusError struct {
	Message string         `json:"message,omitempty"`
	Reason  StatusReason   `json:"reason,omitempty"`
	Details *StatusDetails `json:"details,omitempty"`
	// Status number, 0 if not set.
	Code int32 `json:"code,omitempty"`
}

type StatusDetails struct {
	AgentID   string        `json:"agent_id,omitempty"`
	Subject   string        `json:"subject,omitempty"`
	Operation int32         `json:"operation,omitempty"`
	Causes    []StatusCause `json:"causes,omitempty"`
}

type StatusCause struct {
	Type    CauseType `json:"reason,omitempty"`
	Message string    `json:"message,omitempty"`
}

func (s *StatusError) Error() string {
	return fmt.Sprintf("%s due to reason %s", s.Message, s.Reason)
}

func StatusErrorCause(err error, name CauseType) (StatusCause, bool) {
	s, ok := err.(*StatusError)
	if !ok || s == nil || s.Details == nil {
		return StatusCause{}, false
	}
	for _, cause := range s.Details.Causes {
		if cause.Type == name {
			return cause, true
		}
	}
	return StatusCause{}, false
}

// IsConflict TODO: make 409 enum or const
func IsConflict(err error) bool {
	return CodeForError(err) == 409
}

func IsNotFound(err error) bool {
	return CodeForError(err) == 404
}

func IsInternalError(err error) bool {
	return CodeForError(err) == 500
}

func IsTooManyRequests(err error) bool {
	return CodeForError(err) == 429
}

func CodeForError(err error) int32 {
	if err == nil {
		return -1
	}
	s, ok := err.(*StatusError)
	if !ok {
		return -1
	}
	if strings.Contains(strings.ToLower(fmt.Sprintf("%v", s.Reason)), "not found") {
		return 404
	}
	return s.Code
}

type HTTPError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Reason  string `json:"reason"`
}
