// Package dt a toolkit for the standard time package
package dt

import "time"

const (
	// FormatDefault  default output format in dt package
	FormatDefault = "2006-01-02 15:04:05"
)

// formatMap supported types of format
// TODO: how to support and use more convenient
var formatMap = []string{
	FormatDefault,
	time.ANSIC,
	time.UnixDate,
	time.RubyDate,
	time.RFC822,
	time.RFC822Z,
	time.RFC850,
	time.RFC1123,
	time.RFC1123Z,
	time.RFC3339,
	time.RFC3339Nano,
	time.Kitchen,
	time.Stamp,
	time.StampMilli,
	time.StampMicro,
	time.StampNano,
}

var (
	// LocationUTC UTC setting in string from the time package
	LocationUTC = time.UTC.String()
	// LocationLocal Local setting in string from the time package
	LocationLocal = time.Local.String()
)

// DateTime struct include time.Time and conversion related configuration
// TimeLocation is not the time.Time's location; time.Time has a loc which is the real location setting from time package
// TimeFormats is the target format if you call Format()
type DateTime struct {
	time.Time
	TimeLocation *time.Location
	TimeFormats  string
}

// NewDateTime returns the TimeP struct
func NewDateTime(t time.Time) *DateTime {
	return &DateTime{
		t,
		t.Location(),
		FormatDefault,
	}
}

// Now get TimeP struct for now time
func Now() *DateTime {
	return NewDateTime(time.Now())
}

// Date get DateTime struct for date time by parameter
func Date(year int, month time.Month, day int, hour int, min int, sec int, nsec int, locString string) *DateTime {
	loc, err := time.LoadLocation(locString)
	if err != nil {
		loc = time.UTC
	}
	return NewDateTime(time.Date(year, month, day, hour, min, sec, nsec, loc))
}

// Location set the target location you want to convert
func (t *DateTime) Location(s string) *DateTime {
	loc, err := time.LoadLocation(s)
	if err == nil {
		t.TimeLocation = loc
	}
	return t
}

// Layout set the format layout you want to get
func (t *DateTime) Layout(s string) *DateTime {
	t.TimeFormats = s
	return t
}

// Format Get the string in the target format (TimeP.TimeFormats) and location (TimeP.TimeLocation)
func (t *DateTime) Format() string {
	return t.In(t.TimeLocation).Format(t.TimeFormats)
}
