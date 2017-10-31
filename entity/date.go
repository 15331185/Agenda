package agenda

import (
  "time"
  "strings"
  "strconv"
)
type Date struct {
  DateTime time.Time
}

func changetoint(dateTimeString string, dateTimeArray *[5]int) bool {
    splitBySlash := strings.Split(dateTimeString, "/")
    if len(splitBySlash) < 2 {
    return false
}
    splitByLineThrough := strings.Split(splitBySlash[0], "-")
    splitByColon := strings.Split(splitBySlash[1], ":")
    if len(splitByLineThrough) != 3 && len(splitByColon) != 2 {
  		return false
  	}
    var err error
  	dateTimeArray[0], err = strconv.Atoi(splitByLineThrough[0])
  	if (err != nil) {
  		return false
  	}
  	dateTimeArray[1], err = strconv.Atoi(splitByLineThrough[1])
  	if (err != nil) {
  		return false
  	}
  	dateTimeArray[2], err = strconv.Atoi(splitByLineThrough[2])
  	if (err != nil) {
  		return false
  	}
  	dateTimeArray[3], err = strconv.Atoi(splitByColon[0])
  	if (err != nil) {
  		return false
  	}
  	dateTimeArray[4], err = strconv.Atoi(splitByColon[1])
  	if (err != nil) {
  		return false
  	}
    return true
}

func setDate(year, month, day, hour, min int) time.Time {
      loc , err := time.LoadLocation("")
      if (err != nil) {
    		  panic(err)
    	}
return time.Date(year, (time.Month)(month), day, hour, minute, 0, 0, loc)
}

func IsLeapYear(year int) bool {
	return year % 400 == 0 || (year % 4 == 0 && year % 100 != 0)
}

func IsValidYear(year int) bool {
	return year >= 1000 && year <= 9999
}

func IsValidMonth(month int) bool {
	return month >= (int)(time.January) && month <= (int)(time.December)
}

func IsValidDay(year, month, day int) bool{
	monthDays := [13]int{0, 31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
	if IsLeapYear(year) {	// 闰年二月
		monthDays[2] = 29
	}
	return day > 0 && day <= monthDays[month]
}

func IsValidHour(hour int) bool {
	return hour >= 0 && hour < 24
}

func IsValidMinute(minute int) bool {
	return minute >= 0 && minute < 60
}

func IsValidDateTime(year, month, day, hour, minute int) bool {
	return IsValidYear(year) && IsValidMonth(month) && IsValidDay(year, month, day) &&
			IsValidHour(hour) && IsValidMinute(minute)
}

func (self Date) Equal(other Date) bool {
	return self.DateTime.Equal(other.DateTime)
}

func (self Date) Before(other Date) bool {
	return self.DateTime.Before(other.DateTime)
}

func (self Date) After(other Date) bool {
	return self.DateTime.After(other.DateTime)
}

func (date Date) ToString() string {
	return strings.Join(strings.Split(date.DateTime.String()[0:16], " "), "/")
}
