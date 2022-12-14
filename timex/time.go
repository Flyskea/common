package timex

import "time"

const (
	DateTimeF     = "2006-01-02 15:04:05"
	DateF         = "01-02"
	DateWithYearF = "2006-01-02"
	TimeF         = "15:04"
	TimeWithSecsF = "15:04:05"
)

func DateTimeFx(t time.Time) string {
	return t.Format(DateTimeF)
}

func DateFx(t time.Time) string {
	return t.Format(DateF)
}

func TimeFx(t time.Time) string {
	return t.Format(TimeF)
}

func DateWithYearFx(t time.Time) string {
	return t.Format(DateWithYearF)
}

func TimeWithSecsFx(t time.Time) string {
	return t.Format(TimeWithSecsF)
}
