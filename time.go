package logger

import (
	"strings"
	"time"
)

const timeFormat = "2006-02-01::3:04PM"

type TimeFormatColor struct {
	Year  Attribute
	Sep   Attribute // separator ::
	Date  Attribute
	Style Attribute
}

func (t *TimeFormatColor) getTimeColorFormat(time string) string {
	timeSplit := strings.Split(time, "::")
	year := GetColorFormat(t.Year, t.Style, timeSplit[0])
	date := GetColorFormat(t.Date, t.Style, timeSplit[1])
	sep := GetColorFormat(t.Sep, t.Style, "::")

	result := strings.Join([]string{year, date}, sep) + sep
	return result

}

func getTimeNowFormat() string {
	return time.Now().UTC().Format(timeFormat)
}

func getColorTime(time string) string {

	timeColor := TimeFormatColor{
		Year:  Blue,
		Sep:   Yellow,
		Date:  White,
		Style: Faint,
	}

	time = timeColor.getTimeColorFormat(time)

	return time
}
