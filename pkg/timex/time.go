package timex

import "time"

const (
	nanoToMilliDivisor int64 = 1e6
	milliToSecDivisor  int64 = 1e3
)

var (
	now = time.Now
)

func CurrentUnixMillisecond() int64 {
	return ToUnixMillisecond(now())
}

func ToUnixMillisecond(t time.Time) int64 {
	return t.UnixNano() / nanoToMilliDivisor
}

func FromUnixMillisecond(milli int64) time.Time {
	sec := milli / milliToSecDivisor
	nano := (milli % milliToSecDivisor) * nanoToMilliDivisor
	return time.Unix(sec, nano)
}

func GetTimeString(t time.Time) string {
	return t.Format("2006-01-02T15:04:05")
}

func GetDateString(t time.Time) string {
	return t.Format("2006-1-2")
}

func GetDayTimeString(t time.Time) string {
	return t.Format("15:04:05")
}

func GetTimeFromString(timeString string) (time.Time, error) {
	t, err := time.Parse("2006-01-02T15:04:05", timeString)
	if err != nil {
		return time.Time{}, err
	}

	return t, nil
}

func GetDateFromString(timeString string) (time.Time, error) {
	t, err := time.Parse("2006-1-2", timeString)
	if err != nil {
		return time.Time{}, err
	}

	return t, nil
}

func GetDayTimeFromString(timeString string) (time.Time, error) {
	t, err := time.Parse("15:04:05", timeString)
	if err != nil {
		return time.Time{}, err
	}

	return t, nil
}
