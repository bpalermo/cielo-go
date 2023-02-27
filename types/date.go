package types

import (
	"fmt"
	"strings"
	"time"
)

type DateTime time.Time

const ctLayout = "2006-01-02 15:04:05"

// UnmarshalJSON Parses the json string in the custom format
func (d *DateTime) UnmarshalJSON(b []byte) (err error) {
	s := strings.Trim(string(b), `"`)
	nt, err := time.Parse(ctLayout, s)
	*d = DateTime(nt)
	return
}

// MarshalJSON writes a quoted string in the custom format
func (d *DateTime) MarshalJSON() ([]byte, error) {
	return []byte(d.String()), nil
}

// String returns the time in the custom format
func (d *DateTime) String() string {
	t := time.Time(*d)
	return fmt.Sprintf("%q", t.Format(ctLayout))
}
