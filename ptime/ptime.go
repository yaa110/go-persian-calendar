// In the name of Allah

// Persian Calendar v0.3
// Please visit https://github.com/yaa110/go-persian-calendar for more information.
//
// Copyright (c) 2016 Navid Fathollahzade
// This source code is licensed under MIT license that can be found in the LICENSE file.

// Package ptime provides functionality for implementation of Persian (Solar Hijri) Calendar.
package ptime

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	"time"
)

// A Month specifies a month of the year starting from Farvardin = 1.
type Month int

// A Weekday specifies a day of the week starting from Shanbe = 0.
type Weekday int

// A AmPm specifies the 12-Hour marker.
type AmPm int

// A Time represents a moment in time in Persian (Jalali) Calendar.
type Time struct {
	year  int
	month Month
	day   int
	hour  int
	min   int
	sec   int
	nsec  int
	loc   *time.Location
	wday  Weekday
}

// List of months in Persian calendar.
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

// List of Dari months in Persian calendar.
const (
	Hamal Month = 1 + iota
	Sur
	Jauza
	Saratan
	Asad
	Sonboleh
	Mizan
	Aqrab
	Qos
	Jady
	Dolv
	Hut
)

// List of days in a week.
const (
	Shanbeh Weekday = iota
	Yekshanbeh
	Doshanbeh
	Seshanbeh
	Charshanbeh
	Panjshanbeh
	Jomeh
)

// List of 12-Hour markers.
const (
	Am AmPm = 0 + iota
	Pm
)

var amPm = [2]string{
	"قبل از ظهر",
	"بعد از ظهر",
}

var sAmPm = [2]string{
	"ق.ظ",
	"ب.ظ",
}

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

var dmonths = [12]string{
	"حمل",
	"ثور",
	"جوزا",
	"سرطان",
	"اسد",
	"سنبله",
	"میزان",
	"عقرب",
	"قوس",
	"جدی",
	"دلو",
	"حوت",
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

var sdays = [7]string{
	"ش",
	"ی",
	"د",
	"س",
	"چ",
	"پ",
	"ج",
}

//  {days, leap_days, days_before_start}
var pMonthCount = [12][3]int{
	{31, 31, 0},   // Farvardin
	{31, 31, 31},  // Ordibehesht
	{31, 31, 62},  // Khordad
	{31, 31, 93},  // Tir
	{31, 31, 124}, // Mordad
	{31, 31, 155}, // Shahrivar
	{30, 30, 186}, // Mehr
	{30, 30, 216}, // Aban
	{30, 30, 246}, // Azar
	{30, 30, 276}, // Dey
	{30, 30, 306}, // Bahman
	{29, 30, 336}, // Esfand
}

// Iran returns a pointer to time.Location of Asia/Tehran
func Iran() *time.Location {
	loc, err := time.LoadLocation("Asia/Tehran")
	if err != nil {
		loc = time.FixedZone("Asia/Tehran", 12600) // UTC + 03:30
	}
	return loc
}

// Afghanistan returns a pointer to time.Location of Asia/Kabul
func Afghanistan() *time.Location {
	loc, err := time.LoadLocation("Asia/Kabul")
	if err != nil {
		loc = time.FixedZone("Asia/Kabul", 16200) // UTC + 04:30
	}
	return loc
}

// Returns t in RFC3339Nano format.
func (t Time) String() string {
	return t.Format("yyyy-MM-ddTHH:mm:ss.nsZ")
}

// Dari returns the Dari name of the month.
func (m Month) Dari() string {
	return dmonths[m-1]
}

// Returns the Persian name of the month.
func (m Month) String() string {
	return months[m-1]
}

// Returns the Persian name of the day in week.
func (d Weekday) String() string {
	return days[d]
}

// Short returns the Persian short name of the day in week.
func (d Weekday) Short() string {
	return sdays[d]
}

// Returns the Persian name of 12-Hour marker.
func (a AmPm) String() string {
	return amPm[a]
}

// Short returns the Persian short name of 12-Hour marker.
func (a AmPm) Short() string {
	return sAmPm[a]
}

// New converts Gregorian calendar to Persian calendar and
//
// returns a new instance of Time corresponding to the time of t.
//
// t is an instance of time.Time in Gregorian calendar.
func New(t time.Time) Time {
	pt := new(Time)
	pt.SetTime(t)

	return *pt
}

// Time converts Persian date to Gregorian date and returns a new instance of time.Time
func (t Time) Time() time.Time {
	var year, month, day int

	jdn := getJdn(t.year, int(t.month), t.day)

	if jdn > 2299160 {
		l := jdn + 68569
		n := 4 * l / 146097
		l = l - (146097*n+3)/4
		i := 4000 * (l + 1) / 1461001
		l = l - 1461*i/4 + 31
		j := 80 * l / 2447
		day = l - 2447*j/80
		l = j / 11
		month = j + 2 - 12*l
		year = 100*(n-49) + i + l
	} else {
		j := jdn + 1402
		k := (j - 1) / 1461
		l := j - 1461*k
		n := (l-1)/365 - l/1461
		i := l - 365*n + 30
		j = 80 * i / 2447
		day = i - 2447*j/80
		i = j / 11
		month = j + 2 - 12*i
		year = 4*k + n + i - 4716
	}

	return time.Date(year, time.Month(month), day, t.hour, t.min, t.sec, t.nsec, t.loc)
}

// Date returns a new instance of Time.
//
// year, month and day represent a day in Persian calendar.
//
// hour, min minute, sec seconds, nsec nanoseconds offsets represent a moment in time.
//
// loc is a pointer to time.Location and must not be nil.
func Date(year int, month Month, day, hour, min, sec, nsec int, loc *time.Location) Time {
	if loc == nil {
		panic("ptime: the Location must not be nil in call to Date")
	}

	t := new(Time)
	t.Set(year, month, day, hour, min, sec, nsec, loc)

	return *t
}

// Unix returns a new instance of PersianDate from unix timestamp.
//
// sec seconds and nsec nanoseconds since January 1, 1970 UTC.
//
// loc is a pointer to time.Location and must not be nil.
func Unix(sec, nsec int64, loc *time.Location) Time {
	if loc == nil {
		panic("ptime: the Location must not be nil in call to Unix")
	}

	return New(time.Unix(sec, nsec).In(loc))
}

// Now returns a new instance of Time corresponding to the current time.
//
// loc is a pointer to time.Location and must not be nil.
func Now(loc *time.Location) Time {
	if loc == nil {
		panic("ptime: the Location must not be nil in call to Now")
	}

	return New(time.Now().In(loc))
}

// SetTime sets t to the time of ti.
func (t *Time) SetTime(ti time.Time) {
	var year, month, day int

	t.nsec = ti.Nanosecond()
	t.sec = ti.Second()
	t.min = ti.Minute()
	t.hour = ti.Hour()
	t.loc = ti.Location()
	t.wday = getWeekday(ti.Weekday())

	var jdn int
	gy, gmm, gd := ti.Date()
	gm := int(gmm)

	if gy > 1582 || (gy == 1582 && gm > 10) || (gy == 1582 && gm == 10 && gd > 14) {
		jdn = ((1461 * (gy + 4800 + ((gm - 14) / 12))) / 4) + ((367 * (gm - 2 - 12*((gm-14)/12))) / 12) - ((3 * ((gy + 4900 + ((gm - 14) / 12)) / 100)) / 4) + gd - 32075
	} else {
		jdn = 367*gy - ((7 * (gy + 5001 + ((gm - 9) / 7))) / 4) + ((275 * gm) / 9) + gd + 1729777
	}

	dep := jdn - getJdn(475, 1, 1)
	cyc := dep / 1029983
	rem := dep % 1029983

	var ycyc int
	if rem == 1029982 {
		ycyc = 2820
	} else {
		a := rem / 366
		ycyc = (2134*a+2816*(rem%366)+2815)/1028522 + a + 1
	}

	year = ycyc + 2820*cyc + 474
	if year <= 0 {
		year = year - 1
	}

	var dy = float64(jdn - getJdn(year, 1, 1) + 1)
	if dy <= 186 {
		month = int(math.Ceil(dy / 31.0))
	} else {
		month = int(math.Ceil((dy - 6) / 30.0))
	}

	day = jdn - getJdn(year, month, 1) + 1

	t.year = year
	t.month = Month(month)
	t.day = day
}

// SetUnix sets t to represent the corresponding unix timestamp of
//
// sec seconds and nsec nanoseconds since January 1, 1970 UTC.
//
// loc is a pointer to time.Location and must not be nil.
func (t *Time) SetUnix(sec, nsec int64, loc *time.Location) {
	if loc == nil {
		panic("ptime: the Location must not be nil in call to SetUnix")
	}

	t.SetTime(time.Unix(sec, nsec).In(loc))
}

// Set sets t.
//
// year, month and day represent a day in Persian calendar.
//
// hour, min minute, sec seconds, nsec nanoseconds offsets represent a moment in time.
//
// loc is a pointer to time.Location and must not be nil.
func (t *Time) Set(year int, month Month, day, hour, min, sec, nsec int, loc *time.Location) {
	if loc == nil {
		panic("ptime: the Location must not be nil in call to Set")
	}

	t.year = year
	t.month = month
	t.day = day
	t.hour = hour
	t.min = min
	t.sec = sec
	t.nsec = nsec
	t.loc = loc
	t.resetWeekday()

	t.norm()
}

// SetYear sets the year of t.
func (t *Time) SetYear(year int) {
	t.year = year
	normDay(t)
	t.resetWeekday()
}

// SetMonth sets the month of t.
func (t *Time) SetMonth(month Month) {
	t.month = month
	normMonth(t)
	normDay(t)
	t.resetWeekday()
}

// SetDay sets the day of t.
func (t *Time) SetDay(day int) {
	t.day = day
	normDay(t)
	t.resetWeekday()
}

// SetHour sets the hour of t.
func (t *Time) SetHour(hour int) {
	t.hour = hour
	normHour(t)
}

// SetMinute sets the minute offset of t.
func (t *Time) SetMinute(min int) {
	t.min = min
	normMinute(t)
}

// SetSecond sets the second offset of t.
func (t *Time) SetSecond(sec int) {
	t.sec = sec
	normSecond(t)
}

// SetNanosecond sets the nanosecond offset of t.
func (t *Time) SetNanosecond(nsec int) {
	t.nsec = nsec
	normNanosecond(t)
}

// In sets the location of t.
//
// loc is a pointer to time.Location and must not be nil.
func (t *Time) In(loc *time.Location) {
	if loc == nil {
		panic("ptime: the Location must not be nil in call to In")
	}

	t.loc = loc
	t.resetWeekday()
}

// At sets the hour, min minute, sec second and nsec nanoseconds offsets of t.
func (t *Time) At(hour, min, sec, nsec int) {
	t.hour = hour
	t.min = min
	t.sec = sec
	t.nsec = nsec
	normHour(t)
	normMinute(t)
	normSecond(t)
	normNanosecond(t)
}

// Unix returns the number of seconds since January 1, 1970 UTC.
func (t Time) Unix() int64 {
	return t.Time().Unix()
}

// UnixNano seturns the number of nanoseconds since January 1, 1970 UTC.
func (t Time) UnixNano() int64 {
	return t.Time().UnixNano()
}

// Date returns the year, month, day of t.
func (t Time) Date() (int, Month, int) {
	return t.year, t.month, t.day
}

// Clock returns the hour, minute, seconds offsets of t.
func (t Time) Clock() (int, int, int) {
	return t.hour, t.min, t.sec
}

// Year returns the year of t.
func (t Time) Year() int {
	return t.year
}

// Month returns the month of t in the range [1, 12].
func (t Time) Month() Month {
	return t.month
}

// Day returns the day of month of t.
func (t Time) Day() int {
	return t.day
}

// Hour returns the hour of t in the range [0, 23].
func (t Time) Hour() int {
	return t.hour
}

// Hour12 returns the hour of t in the range [0, 11].
func (t Time) Hour12() int {
	h := t.hour
	if h >= 12 {
		h -= 12
	}

	return h
}

// Minute returns the minute offset of t in the range [0, 59].
func (t Time) Minute() int {
	return t.min
}

// Second returns the seconds offset of t in the range [0, 59].
func (t Time) Second() int {
	return t.sec
}

// Nanosecond returns the nanoseconds offset of t in the range [0, 999999999].
func (t Time) Nanosecond() int {
	return t.nsec
}

// Location returns a pointer to time.Location of t.
func (t Time) Location() *time.Location {
	return t.loc
}

// YearDay returns the day of year of t.
func (t Time) YearDay() int {
	return pMonthCount[t.month-1][2] + t.day
}

// RYearDay returns the number of remaining days of the year of t.
func (t Time) RYearDay() int {
	y := 365
	if t.IsLeap() {
		y++
	}
	return y - t.YearDay()
}

// Weekday returns the weekday of t.
func (t Time) Weekday() Weekday {
	return t.wday
}

// RMonthDay returns the number of remaining days of the month of t.
func (t Time) RMonthDay() int {
	i := 0
	if t.IsLeap() {
		i = 1
	}
	return pMonthCount[t.month-1][i] - t.day
}

// FirstWeekDay returns a new instance of Time representing the first day of the week of t.
func (t Time) FirstWeekDay() Time {
	if t.wday == Shanbeh {
		return t
	}

	return t.AddDate(0, 0, int(Shanbeh-t.wday))
}

// LastWeekday returns a new instance of Time representing the last day of the week of t.
func (t Time) LastWeekday() Time {
	if t.wday == Jomeh {
		return t
	}
	return t.AddDate(0, 0, int(Jomeh-t.wday))
}

// FirstMonthDay returns a new instance of Time representing the first day of the month of t.
func (t Time) FirstMonthDay() Time {
	if t.day == 1 {
		return t
	}

	return Date(t.year, t.month, 1, t.hour, t.min, t.sec, t.nsec, t.loc)
}

// LastMonthDay returns a new instance of Time representing the last day of the month of t.
func (t Time) LastMonthDay() Time {
	i := 0
	if t.IsLeap() {
		i = 1
	}

	ld := pMonthCount[t.month-1][i]

	if ld == t.day {
		return t
	}

	return Date(t.year, t.month, ld, t.hour, t.min, t.sec, t.nsec, t.loc)
}

// FirstYearDay returns a new instance of Time representing the first day of the year of t.
func (t Time) FirstYearDay() Time {
	if t.month == Farvardin && t.day == 1 {
		return t
	}
	return Date(t.year, Farvardin, 1, t.hour, t.min, t.sec, t.nsec, t.loc)
}

// LastYearDay returns a new instance of Time representing the last day of the year of t.
func (t Time) LastYearDay() Time {
	i := 0
	if t.IsLeap() {
		i = 1
	}

	ld := pMonthCount[Esfand-1][i]

	if t.month == Esfand && t.day == ld {
		return t
	}

	return Date(t.year, Esfand, ld, t.hour, t.min, t.sec, t.nsec, t.loc)
}

// MonthWeek returns the week of month of t.
func (t Time) MonthWeek() int {
	return int(math.Ceil(float64(t.day+int(t.FirstMonthDay().Weekday())) / 7.0))
}

// YearWeek returns the week of year of t.
func (t Time) YearWeek() int {
	return int(math.Ceil(float64(t.YearDay()) / 7.0))
}

// RYearWeek returns the number of remaining weeks of the year of t.
func (t Time) RYearWeek() int {
	return 52 - t.YearWeek()
}

// Yesterday returns a new instance of Time representing a day before the day of t.
func (t Time) Yesterday() Time {
	return t.AddDate(0, 0, -1)
}

// Tomorrow returns a new instance of Time representing a day after the day of t.
func (t Time) Tomorrow() Time {
	return t.AddDate(0, 0, 1)
}

// Add returns a new instance of Time for t+d.
func (t Time) Add(d time.Duration) Time {
	return New(t.Time().Add(d))
}

// AddDate returns a new instance of Time for t.year+years, t.month+months and t.day+days.
func (t Time) AddDate(years, months, days int) Time {
	return New(t.Time().AddDate(years, months, days))
}

// Since returns the number of seconds between t and t2.
func (t Time) Since(t2 Time) int64 {
	return int64(math.Abs(float64(t2.Unix() - t.Unix())))
}

// IsLeap returns true if the year of t is a leap year.
func (t Time) IsLeap() bool {
	return divider(25*t.year+11, 33) < 8
}

// AmPm returns the 12-Hour marker of t.
func (t Time) AmPm() AmPm {
	m := Am
	if t.hour > 12 || (t.hour == 12 && (t.min > 0 || t.sec > 0)) {
		m = Pm
	}
	return m
}

// Zone returns the zone name and its offset in seconds east of UTC of t.
func (t Time) Zone() (string, int) {
	return t.Time().Zone()
}

// ZoneOffset returns the zone offset of t in the format of [+|-]HH:mm.
func (t Time) ZoneOffset() string {
	_, offset := t.Zone()

	sign := "+"
	if offset < 0 {
		sign = "-"
	}

	h := offset / 3600
	m := (offset - h*3600) / 60

	return fmt.Sprintf("%s%02d:%02d", sign, h, m)
}

// Format returns the formatted representation of t.
//
//		yyyy, yyy, y     year (e.g. 1394)
//		yy               2-digits representation of year (e.g. 94)
//		MMM              the Persian name of month (e.g. فروردین)
//		MMI              the Dari name of month (e.g. حمل)
//		MM               2-digits representation of month (e.g. 01)
//		M                month (e.g. 1)
//		rw               remaining weeks of year
//		w                week of year
//		RW               remaining weeks of month
//		W                week of month
//		RD               remaining days of year
//		D                day of year
//		rd               remaining days of month
//		dd               2-digits representation of day (e.g. 01)
//		d                day (e.g. 1)
//		E                the Persian name of weekday (e.g. شنبه)
//		e                the Persian short name of weekday (e.g. ش)
//		A                the Persian name of 12-Hour marker (e.g. قبل از ظهر)
//		a                the Persian short name of 12-Hour marker (e.g. ق.ظ)
//		HH               2-digits representation of hour [00-23]
//		H                hour [0-23]
//		kk               2-digits representation of hour [01-24]
//		k                hour [1-24]
//		hh               2-digits representation of hour [01-12]
//		h                hour [1-12]
//		KK               2-digits representation of hour [00-11]
//		K                hour [0-11]
//		mm               2-digits representation of minute [00-59]
//		m                minute [0-59]
//		ss               2-digits representation of seconds [00-59]
//		s                seconds [0-59]
//		ns               nanoseconds
//		S                3-digits representation of milliseconds (e.g. 001)
//		z                the name of location
//		Z                zone offset (e.g. +03:30)
func (t Time) Format(format string) string {
	r := strings.NewReplacer(
		"yyyy", strconv.Itoa(t.year),
		"yyy", strconv.Itoa(t.year),
		"yy", strconv.Itoa(t.year)[2:],
		"y", strconv.Itoa(t.year),
		"MMM", t.month.String(),
		"MMI", t.month.Dari(),
		"MM", fmt.Sprintf("%02d", t.month),
		"M", strconv.Itoa(int(t.month)),
		"rw", strconv.Itoa(t.RYearWeek()),
		"w", strconv.Itoa(t.YearWeek()),
		"W", strconv.Itoa(t.MonthWeek()),
		"RD", strconv.Itoa(t.RYearDay()),
		"D", strconv.Itoa(t.YearDay()),
		"rd", strconv.Itoa(t.RMonthDay()),
		"dd", fmt.Sprintf("%02d", t.day),
		"d", strconv.Itoa(t.day),
		"E", t.wday.String(),
		"e", t.wday.Short(),
		"A", t.AmPm().String(),
		"a", t.AmPm().Short(),
		"HH", fmt.Sprintf("%02d", t.hour),
		"H", strconv.Itoa(t.hour),
		"KK", fmt.Sprintf("%02d", t.Hour12()),
		"K", strconv.Itoa(t.Hour12()),
		"kk", fmt.Sprintf("%02d", modifyHour(t.hour, 24)),
		"k", strconv.Itoa(modifyHour(t.hour, 24)),
		"hh", fmt.Sprintf("%02d", modifyHour(t.Hour12(), 12)),
		"h", strconv.Itoa(modifyHour(t.Hour12(), 12)),
		"mm", fmt.Sprintf("%02d", t.min),
		"m", strconv.Itoa(t.min),
		"ns", strconv.Itoa(t.nsec),
		"ss", fmt.Sprintf("%02d", t.sec),
		"s", strconv.Itoa(t.sec),
		"S", fmt.Sprintf("%03d", t.nsec/1e6),
		"z", t.loc.String(),
		"Z", t.ZoneOffset(),
	)
	return r.Replace(format)
}

func modifyHour(value, max int) int {
	if value == 0 {
		value = max
	}
	return value
}

func (t *Time) norm() {
	normNanosecond(t)
	normSecond(t)
	normMinute(t)
	normHour(t)
	normMonth(t)
	normDay(t)
}

func normNanosecond(t *Time) {
	between(&t.nsec, 0, 999999999)
}

func normSecond(t *Time) {
	between(&t.sec, 0, 59)
}

func normMinute(t *Time) {
	between(&t.min, 0, 59)
}

func normHour(t *Time) {
	between(&t.hour, 0, 23)
}

func normMonth(t *Time) {
	betweenMonth(&t.month, Farvardin, Esfand)
}

func normDay(t *Time) {
	i := 0
	if t.IsLeap() {
		i = 1
	}
	between(&t.day, 1, pMonthCount[t.month-1][i])
}

func betweenMonth(value *Month, min, max Month) {
	if *value < min {
		*value = min
	} else if *value > max {
		*value = max
	}
}

func between(value *int, min, max int) {
	if *value < min {
		*value = min
	} else if *value > max {
		*value = max
	}
}

func divider(num, den int) int {
	if num > 0 {
		return num % den
	}
	return num - ((((num + 1) / den) - 1) * den)
}

func getJdn(year int, month int, day int) int {
	base := year - 473
	if year >= 0 {
		base--
	}

	epy := 474 + (base % 2820)

	var md int
	if month <= 7 {
		md = (month - 1) * 31
	} else {
		md = (month-1)*30 + 6
	}

	return day + md + (epy*682-110)/2816 + (epy-1)*365 + base/2820*1029983 + 1948320
}

func getWeekday(wd time.Weekday) Weekday {
	switch wd {
	case time.Saturday:
		return Shanbeh
	case time.Sunday:
		return Yekshanbeh
	case time.Monday:
		return Doshanbeh
	case time.Tuesday:
		return Seshanbeh
	case time.Wednesday:
		return Charshanbeh
	case time.Thursday:
		return Panjshanbeh
	case time.Friday:
		return Jomeh
	}
	return 0
}

func (t *Time) resetWeekday() {
	t.wday = getWeekday(t.Time().Weekday())
}
