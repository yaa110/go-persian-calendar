// In the name of Allah

// Persian Calendar
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
	"unicode/utf8"
)

// A Month specifies a month of the year starting from Farvardin = 1.
type Month int

// A Weekday specifies a day of the week starting from Shanbe = 0.
type Weekday int

// A AmPm specifies the 12-Hour marker.
type AmPm int

// A DayTime represents a part of the day based on hour.
type DayTime int

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

// List of day times.
const (
	Midnight DayTime = iota
	Dawn
	Morning
	BeforeNoon
	Noon
	AfterNoon
	Evening
	Night
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
	"یک\u200cشنبه",
	"دوشنبه",
	"سه\u200cشنبه",
	"چهارشنبه",
	"پنج\u200cشنبه",
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

var daytimes = [8]string{
	"نیمه\u200cشب",
	"سحر",
	"صبح",
	"قبل از ظهر",
	"ظهر",
	"بعد از ظهر",
	"عصر",
	"شب",
}

// {days, leap_days, days_before_start}
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

// String returns t in RFC3339Nano format.
func (t Time) String() string {
	return t.Format("yyyy-MM-ddTHH:mm:ss.nsZ")
}

// Dari returns the Dari name of the month.
func (m Month) Dari() string {
	switch {
	case m < 1:
		return dmonths[0]
	case m > 11:
		return dmonths[11]
	default:
		return dmonths[m-1]
	}
}

// String returns the Persian name of the month.
func (m Month) String() string {
	switch {
	case m < 1:
		return months[0]
	case m > 11:
		return months[11]
	default:
		return months[m-1]
	}
}

// String returns the Persian name of the day in week.
func (d Weekday) String() string {
	switch {
	case d < 0:
		return days[0]
	case d > 6:
		return days[6]
	default:
		return days[d]
	}
}

// Short returns the Persian short name of the day in week.
func (d Weekday) Short() string {
	switch {
	case d < 0:
		return sdays[0]
	case d > 6:
		return sdays[6]
	default:
		return sdays[d]
	}
}

// String returns the Persian name of 12-Hour marker.
func (a AmPm) String() string {
	if a <= 0 { // Am
		return amPm[0]
	}

	return amPm[1] // Pm
}

// Short returns the Persian short name of 12-Hour marker.
func (a AmPm) Short() string {
	switch {
	case a < 0:
		return sAmPm[0]
	case a > 1:
		return sAmPm[1]
	default:
		return sAmPm[a]
	}
}

// String returns the Persian name of day time.
func (d DayTime) String() string {
	switch {
	case d < 0:
		return daytimes[0]
	case d > 7:
		return daytimes[7]
	default:
		return daytimes[d]
	}
}

// New converts Gregorian calendar to Persian calendar and
//
// returns a new instance of Time corresponding to the time of t or a zero instance of time if Gregorian year is less than 1097.
//
// t is an instance of time.Time in Gregorian calendar.
func New(t time.Time) Time {
	if t.Year() < 1097 {
		return Time{}
	}

	pt := new(Time)
	pt.SetTime(t)

	return *pt
}

// Time converts the Shamsi (Solar Hijri) testDate stored in the Time struct to the corresponding
// Gregorian testDate and returns it as a Go time.Time object.
func (t Time) Time() time.Time {
	var year, month, day int

	// Convert the Shamsi testDate to the corresponding Julian Day Number (JDN)
	jdn := convertShamsiToJDN(t.year, int(t.month), t.day)

	// Convert the JDN to a Gregorian testDate
	if jdn > gregorianReformJulianDay {
		year, month, day = convertJDNToGregorianPostReform(jdn)
	} else {
		year, month, day = convertJDNToGregorianPreReform(jdn)
	}

	// Use the location stored in the Time struct, or default to the local time zone
	loc := t.loc
	if loc == nil {
		loc = time.Local
	}

	// Return the corresponding time.Time object
	return time.Date(year, time.Month(month), day, t.hour, t.min, t.sec, t.nsec, loc)
}

// Date returns a new instance of Time.
//
// year, month and day represent a day in Persian calendar.
//
// hour, min minute, sec seconds, nsec nanoseconds offsets represent a moment in time.
//
// loc is a pointer to time.Location, if loc is nil then the local time is used.
func Date(year int, month Month, day, hour, min, sec, nsec int, loc *time.Location) Time {
	if loc == nil {
		loc = time.Local
	}

	t := new(Time)
	t.Set(year, month, day, hour, min, sec, nsec, loc)

	return *t
}

// Unix returns a new instance of PersianDate from unix timestamp.
//
// sec seconds and nsec nanoseconds since January 1, 1970 UTC.
func Unix(sec, nsec int64) Time {
	return New(time.Unix(sec, nsec))
}

// Now returns a new instance of Time corresponding to the current time.
func Now() Time {
	return New(time.Now())
}

// SetTime sets the time and testDate for the `Time` struct based on the input `time.Time` object.
// This function converts a Gregorian testDate (as provided by `ti`) to the Shamsi (Persian) calendar.
// It first calculates the Julian Day Number (JDN), a continuous count of days since the beginning
// of the Julian Period, and then converts this JDN to a Shamsi testDate.
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

	if isAfterGregorianReform(gy, gm, gd) {
		jdn = convertGregorianPostReformToJDN(gy, gm, gd)
	} else {
		jdn = convertGregorianPreReformToJDN(gy, gm, gd)
	}

	year, month, day = convertJDNToShamsi(jdn)

	t.year = year
	t.month = Month(month)
	t.day = day
}

// SetUnix sets t to represent the corresponding unix timestamp of
//
// sec seconds and nsec nanoseconds since January 1, 1970 UTC.
func (t *Time) SetUnix(sec, nsec int64) {
	t.SetTime(time.Unix(sec, nsec))
}

// norm returns nhi, nlo such that
//
//	hi * base + lo == nhi * base + nlo
//	0 <= nlo < base
func norm(hi, lo, base int) (int, int) {
	if lo < 0 {
		n := (-lo-1)/base + 1
		hi -= n
		lo += n * base
	}
	if lo >= base {
		n := lo / base
		hi += n
		lo -= n * base
	}
	return hi, lo
}

// norm returns nhi, nlo such that
//
//	hi * base + lo == nhi * base + nlo
//	0 <= nlo < base
func normDay(hi, lo, base int) (int, int) {
	if lo < 1 {
		n := (-lo-1)/base + 1
		hi -= n
		lo += n * base
	}
	if lo > base {
		n := lo / base
		hi += n
		lo -= n * base
	}
	return hi, lo
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

	// Normalize nsec, sec, min, hour, overflowing into day.
	sec, nsec = norm(sec, nsec, 1e9)
	min, sec = norm(min, sec, 60)
	hour, min = norm(hour, min, 60)
	day, hour = norm(day, hour, 24)

	// Normalize month, overflowing into year.
	m := int(month) - 1
	year, m = norm(year, m, 12)

	if m < 0 {
		m = 0
	} else if m > 11 {
		m = 11
	}

	if isLeap(year) {
		m, day = normDay(m, day, pMonthCount[m][1])
	} else {
		m, day = normDay(m, day, pMonthCount[m][0])
	}
	year, m = norm(year, m, 12)
	month = Month(m) + 1
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
	t.normDay()
	t.resetWeekday()
}

// SetMonth sets the month of t.
func (t *Time) SetMonth(month Month) {
	t.month = month
	t.normMonth()
	t.normDay()
	t.resetWeekday()
}

// SetDay sets the day of t.
func (t *Time) SetDay(day int) {
	t.day = day
	t.normDay()
	t.resetWeekday()
}

// SetHour sets the hour of t.
func (t *Time) SetHour(hour int) {
	t.hour = hour
	t.normHour()
}

// SetMinute sets the minute offset of t.
func (t *Time) SetMinute(min int) {
	t.min = min
	t.normMinute()
}

// SetSecond sets the second offset of t.
func (t *Time) SetSecond(sec int) {
	t.sec = sec
	t.normSecond()
}

// SetNanosecond sets the nanosecond offset of t.
func (t *Time) SetNanosecond(nsec int) {
	t.nsec = nsec
	t.normNanosecond()
}

// In sets the location of t.
//
// loc is a pointer to time.Location and must not be nil.
func (t Time) In(loc *time.Location) Time {
	if loc == nil {
		panic("ptime: the Location must not be nil in call to In")
	}

	t.loc = loc
	t.resetWeekday()
	return t
}

// At sets the hour, min minute, sec second and nsec nanoseconds offsets of t.
func (t *Time) At(hour, min, sec, nsec int) {
	t.SetHour(hour)
	t.SetMinute(min)
	t.SetSecond(sec)
	t.SetNanosecond(nsec)
}

// IsZero returns true if t is zero time instance
func (t Time) IsZero() bool {
	return t == Time{}
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
	if t.hour >= 12 {
		return t.hour - 12
	}

	return t.hour
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

// DayTime returns the dayTime of that part of the day.
// [0,3)   -> midnight
// [3,6)   -> dawn
// [6,9)   -> morning
// [9,12)  -> before noon
// [12,15) -> noon
// [15,18) -> afternoon
// [18,21) -> evening
// [21,24) -> night
func (t Time) DayTime() DayTime {
	return DayTime(t.hour / 3)
}

// Location returns a pointer to time.Location of t.
func (t Time) Location() *time.Location {
	return t.loc
}

// YearDay returns the day of year of t.
func (t Time) YearDay() int {
	m := int(t.month - 1)
	switch { // isInBounds()
	case m < 0:
		return pMonthCount[0][2] + t.day
	case m > 11:
		return pMonthCount[11][2] + t.day
	default:
		return pMonthCount[m][2] + t.day
	}
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

	m := t.month - 1

	switch { // isInBounds()
	case m < 0:
		return pMonthCount[0][i] - t.day
	case m > 11:
		return pMonthCount[11][i] - t.day
	default:
		return pMonthCount[m][i] - t.day
	}
}

// BeginningOfWeek returns a new instance of Time representing the first day of the week of t.
// The time is reset to 00:00:00
func (t Time) BeginningOfWeek() Time {
	nt := t.AddDate(0, 0, int(Shanbeh-t.wday))
	nt.SetHour(0)
	nt.SetMinute(0)
	nt.SetSecond(0)
	nt.SetNanosecond(0)
	return nt
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

// BeginningOfMonth returns a new instance of Time representing the first day of the month of t.
// The time is reset to 00:00:00
func (t Time) BeginningOfMonth() Time {
	return Date(t.year, t.month, 1, 0, 0, 0, 0, t.loc)
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

	m := t.month - 1
	if m < 0 {
		m = 0
	} else if m > 11 {
		m = 11
	}

	ld := pMonthCount[m][i]
	if ld == t.day {
		return t
	}
	return Date(t.year, t.month, ld, t.hour, t.min, t.sec, t.nsec, t.loc)
}

// BeginningOfYear returns a new instance of Time representing the first day of the year of t.
// The time is reset to 00:00:00
func (t Time) BeginningOfYear() Time {
	return Date(t.year, Farvardin, 1, 0, 0, 0, 0, t.loc)
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
	return int(math.Ceil(float64(t.YearDay()+int(t.FirstYearDay().Weekday())) / 7.0))
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
	t.Set(t.year+years, Month(int(t.month)+months), t.day+days, t.hour, t.min, t.sec, t.nsec, t.loc)
	return t
}

// Since returns the number of seconds between t and t2.
func (t Time) Since(t2 Time) int64 {
	return int64(math.Abs(float64(t2.Unix() - t.Unix())))
}

// IsLeap returns true if the year of t is a leap year.
func (t Time) IsLeap() bool {
	return isLeap(t.year)
}

func isLeap(year int) bool {
	return divider(25*year+11, 33) < 8
}

// AmPm returns the 12-Hour marker of t.
func (t Time) AmPm() AmPm {
	if t.hour > 12 || (t.hour == 12 && (t.min > 0 || t.sec > 0)) {
		return Pm
	}
	return Am
}

// Zone returns the zone name and its offset in seconds east of UTC of t.
func (t Time) Zone() (string, int) {
	return t.Time().Zone()
}

// ZoneOffset returns the zone offset of t in the format of [+|-]HH:mm.
// If `f` is set, then return format is based on `f`.
func (t Time) ZoneOffset(f ...string) string {
	format := "-07:00"
	if len(f) > 0 {
		format = f[0]
		if format != "-0700" && format != "-07" && format != "-07:00" && format != "Z0700" && format != "Z07:00" {
			format = "-07:00"
		}
	}

	_, offset := t.Zone()

	if offset == 0 {
		switch format {
		case "-0700":
			return "+0000"
		case "-07":
			return "+00"
		case "-07:00":
			return "+00:00"
		case "Z0700", "Z07:00":
			return "Z"
		}
	}

	h := offset / 3600
	m := (offset - h*3600) / 60

	switch format {
	case "-0700", "Z0700":
		return fmt.Sprintf("%+03d%02d", h, m)
	case "-07":
		return fmt.Sprintf("%+03d", h)
	default:
		return fmt.Sprintf("%+03d:%02d", h, m)
	}
}

// Format returns the formatted representation of t.
//
//	yyyy, yyy, y     year (e.g. 1394)
//	yy               2-digits representation of year (e.g. 94)
//	MMM              the Persian name of month (e.g. فروردین)
//	MMI              the Dari name of month (e.g. حمل)
//	MM               2-digits representation of month (e.g. 01)
//	M                month (e.g. 1)
//	rw               remaining weeks of year
//	w                week of year
//	RW               remaining weeks of month
//	W                week of month
//	RD               remaining days of year
//	D                day of year
//	rd               remaining days of month
//	dd               2-digits representation of day (e.g. 01)
//	d                day (e.g. 1)
//	E                the Persian name of weekday (e.g. شنبه)
//	e                the Persian short name of weekday (e.g. ش)
//	A                the Persian name of 12-Hour marker (e.g. قبل از ظهر)
//	a                the Persian short name of 12-Hour marker (e.g. ق.ظ)
//	HH               2-digits representation of hour [00-23]
//	H                hour [0-23]
//	kk               2-digits representation of hour [01-24]
//	k                hour [1-24]
//	hh               2-digits representation of hour [01-12]
//	h                hour [1-12]
//	KK               2-digits representation of hour [00-11]
//	K                hour [0-11]
//	mm               2-digits representation of minute [00-59]
//	m                minute [0-59]
//	ss               2-digits representation of seconds [00-59]
//	s                seconds [0-59]
//	n				 hour name (e.g. صبح)
//	ns               nanoseconds
//	S                3-digits representation of milliseconds (e.g. 001)
//	z                the name of location
//	Z                zone offset (e.g. +03:30)
func (t Time) Format(format string) string {
	if format == "" {
		return ""
	}

	var (
		i  int
		sb strings.Builder
	)

	sb.Grow(2 * len(format)) // double the format len, the formatted value likely to be longer than format

	writeD2 := func(v int) {
		if v < 10 {
			sb.WriteByte('0')
		}

		sb.WriteString(strconv.Itoa(v))
	}

	writeD3 := func(v int) {
		switch {
		case v >= 100: // noop
		case v >= 10:
			sb.WriteByte('0')
		case v >= 0:
			sb.WriteString("00")
		}

		sb.WriteString(strconv.Itoa(v))
	}

	writeD4 := func(v int) {
		switch {
		case v >= 1000: // noop
		case v >= 100:
			sb.WriteByte('0')
		case v >= 10:
			sb.WriteString("00")
		case v >= 0:
			sb.WriteString("000")
		}

		sb.WriteString(strconv.Itoa(v))
	}

	// peek returns next character
	peek := func() byte {
		j := i + 1

		if j < 0 || j >= len(format) { // IsInBounds()
			return 0
		}

		return format[j]
	}

	for {
		if i < 0 || i >= len(format) { // isSliceInBounds()
			break
		}

		current := format[i:]

		switch format[i] {
		case 'A':
			sb.WriteString(t.AmPm().String())
			i++
		case 'D':
			sb.WriteString(strconv.Itoa(t.YearDay()))
			i++
		case 'E':
			sb.WriteString(t.wday.String())
			i++
		case 'H':
			if peek() == 'H' { // HH
				writeD2(t.hour)
				i += 2
			} else { // H
				sb.WriteString(strconv.Itoa(t.hour))
				i++
			}
		case 'K':
			if peek() == 'K' { // KK
				writeD2(t.Hour12())
				i += 2
			} else { // K
				sb.WriteString(strconv.Itoa(t.Hour12()))
				i++
			}
		case 'M':
			switch {
			default: // M
				sb.WriteString(strconv.Itoa(int(t.month)))
				i++
			case strings.HasPrefix(current, "MMM"):
				sb.WriteString(t.month.String())
				i += 3
			case strings.HasPrefix(current, "MMI"):
				sb.WriteString(t.month.Dari())
				i += 3
			case peek() == 'M': // MM
				writeD2(int(t.month))
				i += 2
			}
		case 'R':
			if peek() == 'D' { // RD
				sb.WriteString(strconv.Itoa(t.RYearDay()))
				i += 2
			} else { // R
				sb.WriteByte('R')
				i++
			}
		case 'S':
			writeD3(t.nsec / 1e6)
			i++
		case 'W':
			sb.WriteString(strconv.Itoa(t.MonthWeek()))
			i++
		case 'Z':
			sb.WriteString(t.ZoneOffset())
			i++
		case 'a':
			sb.WriteString(t.AmPm().Short())
			i++
		case 'd':
			if peek() == 'd' { // dd
				writeD2(t.day)
				i += 2
			} else { // d
				sb.WriteString(strconv.Itoa(t.day))
				i++
			}
		case 'e':
			sb.WriteString(t.wday.Short())
			i++
		case 'h':
			if peek() == 'h' { // hh
				writeD2(modifyHour(t.Hour12(), 12))
				i += 2
			} else { // h
				sb.WriteString(strconv.Itoa(modifyHour(t.Hour12(), 12)))
				i++
			}
		case 'k':
			if peek() == 'k' { // kk
				writeD2(modifyHour(t.hour, 24))
				i += 2
			} else { // k
				sb.WriteString(strconv.Itoa(modifyHour(t.hour, 24)))
				i++
			}
		case 'm':
			if peek() == 'm' { // mm
				writeD2(t.min)
				i += 2
			} else { // m
				sb.WriteString(strconv.Itoa(t.min))
				i++
			}
		case 'n':
			if peek() == 's' { // ns
				sb.WriteString(strconv.Itoa(t.nsec))
				i += 2
			} else { // n
				sb.WriteString(t.DayTime().String())
				i++
			}
		case 'r':
			switch peek() {
			default: // r
				sb.WriteByte('r')
				i++
			case 'w': // rw
				sb.WriteString(strconv.Itoa(t.RYearWeek()))
				i += 2
			case 'd': // rd
				sb.WriteString(strconv.Itoa(t.RMonthDay()))
				i += 2
			}
		case 's':
			if peek() == 's' { // ss
				writeD2(t.sec)
				i += 2
			} else { // s
				sb.WriteString(strconv.Itoa(t.sec))
				i++
			}
		case 'w':
			sb.WriteString(strconv.Itoa(t.YearWeek()))
			i++
		case 'y':
			switch {
			default: // y
				writeD4(t.year)
				i++
			case strings.HasPrefix(current, "yyyy"):
				writeD4(t.year)
				i += 4
			case strings.HasPrefix(current, "yyy"):
				writeD4(t.year)
				i += 3
			case peek() == 'y': // yy
				switch s := strconv.Itoa(t.year); len(s) {
				default:
					sb.WriteString(s[len(s)-2:])
				case 1:
					sb.WriteString("0" + s)
				case 2:
					sb.WriteString(s)
				}

				i += 2
			}
		case 'z':
			sb.WriteString(t.loc.String())
			i++
		default:
			r, n := utf8.DecodeRuneInString(current)
			sb.WriteRune(r)
			i += n
		}
	}

	return sb.String()
}

// TimeFormat formats in standard time format.
//
//	2006        four digit year (e.g. 1399)
//	06          two digit year (e.g. 99)
//	01          two digit month (e.g. 01)
//	1           one digit month (e.g. 1)
//	Jan         month name (e.g. آذر)
//	January     month name (e.g. آذر)
//	02          two digit day (e.g. 07)
//	2           one digit day (e.g. 7)
//	_2          right justified two character day (e.g.  7)
//	Mon         weekday (e.g. شنبه)
//	Monday      weekday (e.g. شنبه)
//	Morning     hour name (e.g. صبح)
//	03          two digit 12 hour format (e.g. 03)
//	3           one digit 12 hour format (e.g. 3)
//	15          two digit 24 hour format (e.g. 15)
//	04          two digit minute (e.g. 03)
//	4           one digit minute (e.g. 03)
//	05          two digit minute (e.g. 09)
//	5           one digit minute (e.g. 9)
//	.000        millisecond (e.g. .120)
//	.000000     microsecond (e.g. .123400)
//	.000000000  nanosecond (e.g. .123456000)
//	.999        trailing zeros removed millisecond (e.g. .12)
//	.999999     trailing zeros removed microsecond (e.g. .1234)
//	.999999999  trailing zeros removed nanosecond (e.g. .123456)
//	PM          full 12-Hour marker (e.g. قبل از ظهر)
//	pm          short 12-Hour marker (e.g. ق.ظ)
//	MST         the name of location
//	-0700       zone offset (e.g. +0330)
//	-07         zone offset (e.g. +03)
//	-07:00      zone offset (e.g. +03:30)
//	Z0700       zone offset (e.g. +0330)
//	Z07:00      zone offset (e.g. +03:30)
func (t Time) TimeFormat(format string) string {
	initializer := []string{
		"January", "{MMMM}",
		"Jan", "{MMM}",
		"Monday", "{WD}",
		"Mon", "{W}",
		"Morning", "{n}",
		".000000000", "{ns}",
		".000000", "{ms}",
		".000", "{mls}",
		".999999999", "{nst}",
		".999999", "{mst}",
		".999", "{mlst}",
		"2006", "{YYYY}",
		"PM", "{AFTER}",
		"pm", "{after}",
		"MST", "{LOC}",
		"Z0700", "{Z0700}",
		"Z07:00", "{Z07:00}",
		"-0700", "{-0700}",
		"-07:00", "{-07:00}",
		"-07", "{-07}",
		"15", "{HH}",
		"06", "{YY}",
		"01", "{MM}",
		"02", "{DD}",
		"03", "{hh}",
		"04", "{mm}",
		"05", "{ss}",
		"_2", "{_D}",
		"1", "{M}",
		"2", "{D}",
		"3", "{h}",
		"4", "{m}",
		"5", "{s}",
	}

	r := strings.NewReplacer(initializer...)
	formatted := r.Replace(format)

	nsec := strconv.Itoa(t.nsec)
	if len(nsec) < 6 {
		nsec = fmt.Sprintf("%06d", t.nsec)
	}

	year := strconv.Itoa(t.year)
	if len(year) < 4 {
		year = fmt.Sprintf("%04d", t.year)
	}

	params := []string{
		"{YYYY}", year,
		"{YY}", year[2:],
		"{MMMM}", t.locMonthName(),
		"{MMM}", t.locMonthName(),
		"{MM}", fmt.Sprintf("%02d", int(t.month)),
		"{M}", strconv.Itoa(int(t.month)),
		"{DD}", fmt.Sprintf("%02d", t.day),
		"{_D}", fmt.Sprintf("%2d", t.day),
		"{D}", strconv.Itoa(t.day),
		"{WD}", t.wday.String(),
		"{W}", t.wday.Short(),
		"{n}", t.DayTime().String(),
		"{HH}", fmt.Sprintf("%02d", t.hour),
		"{hh}", fmt.Sprintf("%02d", t.Hour12()),
		"{h}", strconv.Itoa(t.Hour12()),
		"{mm}", fmt.Sprintf("%02d", t.min),
		"{m}", strconv.Itoa(t.min),
		"{ss}", fmt.Sprintf("%02d", t.sec),
		"{s}", strconv.Itoa(t.sec),
		"{ns}", "." + nsec,
		"{ms}", "." + nsec[:6],
		"{mls}", "." + nsec[:3],
		"{nst}", strings.TrimRight("."+nsec, "0"),
		"{mst}", strings.TrimRight("."+nsec[:6], "0"),
		"{mlst}", strings.TrimRight("."+nsec[:3], "0"),
		"{AFTER}", t.AmPm().String(),
		"{after}", t.AmPm().Short(),
		"{LOC}", t.loc.String(),
		"{Z0700}", t.ZoneOffset("Z0700"),
		"{Z07:00}", t.ZoneOffset("Z07:00"),
		"{-0700}", t.ZoneOffset("-0700"),
		"{-07:00}", t.ZoneOffset("-07:00"),
		"{-07}", t.ZoneOffset("-07"),
	}

	r = strings.NewReplacer(params...)
	return r.Replace(formatted)
}

func (t *Time) locMonthName() string {
	if t.Location().String() == Afghanistan().String() {
		return t.month.Dari()
	}
	return t.month.String()
}

func (t *Time) norm() {
	t.normNanosecond()
	t.normSecond()
	t.normMinute()
	t.normHour()
	t.normMonth()
	t.normDay()
}

func (t *Time) normNanosecond() {
	between(&t.nsec, 0, 999999999)
}

func (t *Time) normSecond() {
	between(&t.sec, 0, 59)
}

func (t *Time) normMinute() {
	between(&t.min, 0, 59)
}

func (t *Time) normHour() {
	between(&t.hour, 0, 23)
}

func (t *Time) normMonth() {
	betweenMonth(&t.month, Farvardin, Esfand)
}

func (t *Time) normDay() {
	i := 0
	if t.IsLeap() {
		i = 1
	}

	m := t.month - 1
	if m < 0 {
		m = 0
	} else if m > 11 {
		m = 11
	}

	between(&t.day, 1, pMonthCount[m][i])
}

func modifyHour(value, max int) int {
	if value == 0 {
		return max
	}
	return value
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
