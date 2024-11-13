package util

import (
	"strings"
	"time"

	"github.com/araddon/dateparse"
)

type Time struct {
	time.Time
}

func (m *Time) UnmarshalJSON(data []byte) (err error) {
	strtt := strings.Trim(string(data), "\"")

	var tt time.Time
	tt, err = dateparse.ParseAny(strtt)

	if err != nil {
		return
	}

	*m = Time{tt}
	return
}
