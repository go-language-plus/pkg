package timep

import (
	"testing"
	"time"
)

func TestTimeP_Format(t1 *testing.T) {
	timeNow := Now()
	timeDate := Date(2009, time.November, 10, 23, 0, 0, 0, LocationUTC)
	timeDateUnknown := Date(2009, time.November, 10, 23, 0, 0, 0, "UnknownLocation")
	timeDateLocation := Date(2009, time.November, 10, 23, 0, 0, 0, LocationUTC).Location("Asia/Shanghai")
	timeDateLayout := Date(2009, time.November, 10, 23, 0, 0, 0, LocationUTC).Layout(time.RFC1123)

	tests := []struct {
		name  string
		timeP *TimeP
		want  string
	}{
		{
			"test_now",
			timeNow,
			timeNow.Format(),
		},
		{
			"test_date",
			timeDate,
			"2009-11-10 23:00:00",
		},
		{
			"test_date",
			timeDateUnknown,
			"2009-11-10 23:00:00",
		},
		{
			"test_date_location",
			timeDateLocation,
			"2009-11-11 07:00:00",
		},
		{
			"test_date_layout",
			timeDateLayout,
			"Tue, 10 Nov 2009 23:00:00 UTC",
		},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			if got := tt.timeP.Format(); got != tt.want {
				t1.Errorf("Format() = %v, want %v", got, tt.want)
			}
		})
	}
}
