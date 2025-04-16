// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"fmt"
	"io"
	"strconv"
)

type PerfRemindableStatus struct {
	Status  RemindableStatus `json:"status"`
	Message string           `json:"message"`
}

type RemindableStatus string

const (
	RemindableStatusCanRemind RemindableStatus = "CAN_REMIND"
	RemindableStatusTooSoon   RemindableStatus = "TOO_SOON"
)

var AllRemindableStatus = []RemindableStatus{
	RemindableStatusCanRemind,
	RemindableStatusTooSoon,
}

func (e RemindableStatus) IsValid() bool {
	switch e {
	case RemindableStatusCanRemind, RemindableStatusTooSoon:
		return true
	}
	return false
}

func (e RemindableStatus) String() string {
	return string(e)
}

func (e *RemindableStatus) UnmarshalGQL(v any) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = RemindableStatus(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid RemindableStatus", str)
	}
	return nil
}

func (e RemindableStatus) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
