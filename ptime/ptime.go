// In the name of Allah

// Persian Calendar v0.1
// Please visit https://github.com/yaa110/go-persian-calendar for more information.
//
// Copyright (c) 2016 Navid Fathollahzade
// This source code is licensed under MIT license that can be found in the LICENSE file.

// Package ptime provides functionality for implementation of Persian (Jalali) Calendar.
package ptime

import (
	"time"
	"math"
)

// A Month specifies a month of the year in Persian calendar starting from 1.
type Month int

// A Weekday specifies a day of the week in Persian starting from 0.
type Weekday int

// A PersianDate represents a day in Persian (Jalali) Calendar.
type Time struct {
	year int
	month Month
	day int
	hour int
	min int
	sec int
	nsec int
	loc *time.Location
}

const (
	Farvardin Month = 1 + iota
	Ordibehesht
	Khordad
	Tir
	Mordad
	Shahrivar
	Mehr
	Aban
	Azar
	Dey
	Bahman
	Esfand
)

const (
	Shanbe Weekday = iota
	Yekshanbe
	Doshanbe
	Seshanbe
	Charshanbe
	Panjshanbe
	Jome
)

var months = [12]string{
	"فروردین",
	"اردیبهشت",
	"خرداد",
	"تیر",
	"مرداد",
	"شهریور",
	"مهر",
	"آبان",
	"آذر",
	"دی",
	"بهمن",
	"اسفند",
}

var days = [7]string{
	"شنبه",
	"یک‌شنبه",
	"دوشنبه",
	"سه‌شنبه",
	"چهارشنبه",
	"پنج‌شنبه",
	"جمعه",
}

//  {days,   leap_days}
var p_month_count = [12][2]int {
	{31,     31},     // Farvardin
	{31,     31},     // Ordibehesht
	{31,     31},     // Khordad
	{31,     31},     // Tir
	{31,     31},     // Mordad
	{31,     31},     // Shahrivar
	{30,     30},     // Mehr
	{30,     30},     // Aban
	{30,     30},     // Azar
	{30,     30},     // Dey
	{30,     30},     // Bahman
	{29,     30},     // Esfand
}

// Returns the Persian name of the month.
func (m Month) String() string {
	return months[m - 1]
}

// Returns the Persian name of the day in week.
func (d Weekday) String() string {
	return days[d]
}

func Time(t time.Time) Time {
	pt := Time{}
	&pt.SetTime(t)

	return pt
}

// Converts Persian date to Gregorian date and returns an instance of time.Time
func (t Time) Time() time.Time {
	var year, month, day int

	jdn := getJdn(t.year, t.month, t.day)

	if jdn > 2299160 {
		l := jdn + 68569
		n := 4 * l / 146097
		l = l - (146097 * n + 3) / 4
		i := 4000 * (l + 1) / 1461001
		l = l - 1461 * i / 4 + 31
		j := 80 * l / 2447
		day = l - 2447 * j / 80
		l = j / 11
		month = j + 2 - 12 * l
		year = 100 * (n - 49) + i + l
	} else {
		j := jdn + 1402
		k := (j - 1) / 1461
		l := j - 1461 * k
		n := (l - 1) / 365 - l / 1461
		i := l - 365 * n + 30
		j = 80 * i / 2447
		day = i - 2447 * j / 80
		i = j / 11
		month = j + 2 - 12 * i
		year = 4 * k + n + i - 4716
	}

	return time.Date(year, month, day, t.hour, t.min, t.sec, t.nsec, t.loc)
}

// Returns a new instance of PersianDate.
func Date(year int, month Month, day, hour, min, sec, nsec int, loc *time.Location) Time {
	if loc == nil {
		panic("ptime: the Location must not be nil in call to Date")
	}

	t := Time{year, month, day, hour, min, sec, nsec, loc}
	&t.norm()

	return t
}

// Returns a new instance of PersianDate from unix timestamp.
// sec seconds and nsec nanoseconds since January 1, 1970 UTC.
func Unix(sec, nsec int64, loc *time.Location) Time {
	if loc == nil {
		panic("ptime: the Location must not be nil in call to Unix")
	}

	return Time(time.Unix(sec, nsec).In(loc))
}

func Now(loc *time.Location) Time {
	if loc == nil {
		panic("ptime: the Location must not be nil in call to Now")
	}

	return Time(time.Now().In(loc))
}

// Converts Gregorian date to Persian date.
// TODO has bug
func (pt *Time) SetTime(t time.Time) {
	var year, month, day int

	pt.nsec = t.Nanosecond()
	pt.sec = t.Second()
	pt.min = t.Minute()
	pt.hour = t.Hour()
	pt.loc = t.Location()

	var jdn int
	gy, gm, gd := t.Date()

	if gy > 1582 || (gy == 1582 && gm > 10) || (gy == 1582 && gm == 10 && gd > 14) {
		jdn = ((1461 * (gy + 4800 + ((gm - 14) / 12))) / 4) + ((367 * (gm - 2 - 12 * (((gm - 14) / 12)))) / 12) - ((3 * (((gy + 4900 + ((gm - 14) / 12)) / 100))) / 4) + gd - 32075
	} else {
		jdn = 367 * gy - ((7 * (gy + 5001 + ((gm - 9) / 7))) / 4) + ((275 * gm) / 9) + gd + 1729777
	}

	dep := jdn - getJdn(475, 1, 1)
	cyc := dep / 1029983
	rem := dep % 1029983

	var ycyc int
	if rem == 1029982 {
		ycyc = 2820
	} else {
		a := rem / 366
		ycyc = (2134 * a + 2816 * (rem % 366) + 2815) / 1028522 + a + 1;
	}

	year = ycyc + 2820 * cyc + 474
	if year <= 0 {
		year = year - 1
	}

	var dy float64 = float64(jdn - getJdn(year, 1, 1) + 1)
	if dy <= 186 {
		month = int(math.Ceil(dy / 31.0))
	} else {
		month = int(math.Ceil((dy - 6) / 30.0))
	}

	day = jdn - getJdn(year, month, 1) + 1

	pt.year = year
	pt.month = month
	pt.day = day
}

// Changes t using unix timestamp
func (t *Time) SetUnix(sec, nsec int64, loc *time.Location) {
	if loc == nil {
		panic("ptime: the Location must not be nil in call to SetUnix")
	}

	t.SetTime(time.Unix(sec, nsec).In(loc))
}

func (t *Time) Set(year int, month Month, day, hour, min, sec, nsec int, loc *time.Location) {
	if loc == nil {
		panic("ptime: the Location must not be nil in call to Change")
	}

	t.year = year
	t.month = month
	t.day = day
	t.hour = hour
	t.min = min
	t.sec = sec
	t.nsec = nsec
	t.loc = loc

	t.norm()
}

func (t *Time) SetYear(year int) {
	t.year = year
	norm_day(t)
}

func (t *Time) SetMonth(month Month) {
	t.month = month
	norm_month(t)
	norm_day(t)
}

func (t *Time) SetDay(day int) {
	t.day = day
	norm_day(t)
}

func (t *Time) SetHour(hour int) {
	t.hour = hour
	norm_hour(t)
}

func (t *Time) SetMinute(min int) {
	t.min = min
	norm_minute(t)
}

func (t *Time) SetSecond(sec int) {
	t.sec = sec
	norm_second(t)
}

func (t *Time) SetNanosecond(nsec int) {
	t.nsec = nsec
	norm_nanosecond(t)
}

func (t *Time) In(loc *time.Location) {
	if loc == nil {
		panic("ptime: the Location must not be nil in call to In")
	}

	t.loc = loc
}

func (t *Time) At(hour, min, sec, nsec int) {
	t.hour = hour
	t.min = min
	t.sec = sec
	t.nsec = nsec
	norm_hour(t)
	norm_minute(t)
	norm_second(t)
	norm_nanosecond(t)
}

// Returns unix timestamp (the number of seconds) of t.
func (t Time) Unix() int64 {
	return t.Time().Unix()
}

// Returns unix timestamp (the number of nanoseconds) of t.
func (t Time) UnixNano() int64 {
	return t.Time().UnixNano()
}

// Returns the year, month, day of t.
func (t Time) Date() (int, Month, int) {
	return t.year, t.month, t.day
}

// Returns the year of t.
func (t Time) Year() int {
	return t.year
}

// Returns the month of t.
func (t Time) Month() Month {
	return t.month
}

// Returns the day in month of t.
func (t Time) Day() int {
	return t.day
}

// Returns the hour of t in the range [0, 23].
func (t Time) Hour() int {
	return t.hour
}

// Returns the minute offset of t in the range [0, 59].
func (t Time) Minute() int {
	return t.min
}

// Returns the second offset of t in the range [0, 59].
func (t Time) Second() int {
	return t.sec
}

// Returns the nanosecond offset of t in the range [0, 999999999].
func (t Time) Nanosecond() int {
	return t.nsec
}

// Returns the time zone information of t.
// For more information check the documentation of time.Location
func (t Time) Location() *time.Location {
	return t.loc
}

// Returns the day in the year of t.
func (t Time) YearDay() int {
	// TODO YearDay of PersianDate
	return 0
}

// Returns the weekday of t.
func (t Time) Weekday() Weekday {
	// TODO Weekday of PersianDate
	return 0
}

// Returns the number of remaining days in the year of t.
func (t Time) RYearDay() int {
	// TODO RYearDay
	return 0
}

// Returns the number of remaining days in the month of t.
func (t Time) RMonthDay() int {
	// TODO RMonthDay
	return 0
}

// Returns the number of remaining days in the week of t.
func (t Time) RWeekday() int {
	// TODO RWeekday
	return 0
}

func (t Time) FirstDayInWeek() Time {
	// TODO return the first day in the week
	return nil
}

func (t Time) FirstDayInMonth() Time {
	// TODO return the first day in the month
	return nil
}

func (t Time) FirstDayInYear() Time {
	// TODO return the first day in the year
	return nil
}

func (t Time) LastDayInWeek() Time {
	// TODO return the last day in the week
	return nil
}

func (t Time) LastDayInMonth() Time {
	// TODO return the last day in the month
	return nil
}

func (t Time) LastDayInYear() Time {
	// TODO return the last day in the year
	return nil
}

func (t Time) MonthWeek() int {
	// TODO return the week number in the month
	return 0
}

func (t Time) YearWeek() int {
	// TODO return the week number in the year
	return 0
}

func (t Time) Yesterday() int {
	// TODO return Yesterday
	return 0
}

func (t Time) Tomorrow() int {
	// TODO return Tomorrow
	return 0
}

// Returns a new instance of Time for t+d.
func (t Time) Add(d time.Duration) Time {
	return Time(t.Time().Add(d))
}

func (t Time) AddDate(years, months, days int) Time {
	return Time(t.Time().AddDate(years, months, days))
}

// Returns the time.Duration between t and t2
func (t Time) Since(t2 Time) time.Duration {
	return math.Abs(t2.Unix() - t.Unix()) * time.Second
}

// Returns true if the year of t is a leap year.
func (t Time) IsLeap() bool {
	return divider(25 * t.year + 11, 33) < 8
}

// Modifies the year, month and day if they were outside their usual ranges.
func (t *Time) norm() {
	norm_nanosecond(t)
	norm_second(t)
	norm_minute(t)
	norm_hour(t)
	norm_month(t)
	norm_day(t)
}

func norm_nanosecond(t *Time)  {
	between(&t.nsec, 0, 999999999)
}

func norm_second(t *Time)  {
	between(&t.sec, 0, 59)
}

func norm_minute(t *Time)  {
	between(&t.min, 0, 59)
}

func norm_hour(t *Time)  {
	between(&t.hour, 0, 23)
}

func norm_month(t *Time)  {
	between(&t.month, 1, 12)
}

func norm_day(t *Time)  {
	i := 0
	if t.IsLeap() {
		i = 1
	}
	between(&t.day, 1, p_month_count[t.month - 1][i])
}

func between(value *int, min, max int) {
	if *value < min {
		*value = min
	} else if *value > max {
		*value = max
	}
}

func divider(num, den int) int {
	if (num > 0) {
		return num % den
	}
	return num - ((((num + 1) / den) - 1) * den)
}

func getJdn(year, month, day int) int {
	base := year - 473
	if year >= 0 {
		base -= 1
	}

	epy := 474 + (base % 2820)

	var md int
	if month <= 7 {
		md = (month - 1) * 31
	} else {
		md = (month - 1) * 30 + 6
	}

	return day + md + (epy * 682 - 110) / 2816 + (epy - 1) * 365 + base / 2820 * 1029983 + 1948320
}