// In the name of Allah

// Copyright (c) 2016 Navid Fathollahzade
// This source code is licensed under MIT license that can be found in the LICENSE file.

// Version: 0.1
// Please visit https://github.com/yaa110/go-persian-calendar for more information.

// Package ptime provides functionality for implementation of Persian (Jalali) Calendar.
package ptime

import "time"

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
	Chaharshanbe
	Panjshanbe
	Jome
)

const (
	persian_epoch = 226899

	month_count_normal = 0
	month_count_leap = 1
	month_count_normal_before = 2
	month_count_leap_before = 3
)

var months = [...]string{
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

var days = [...]string{
	"شنبه",
	"یک‌شنبه",
	"دوشنبه",
	"سه‌شنبه",
	"چهارشنبه",
	"پنج‌شنبه",
	"جمعه",
}

var p_month_count = [...][...]int {
	{31,     31,      0},       // Farvardin
	{31,     31,      31},      // Ordibehesht
	{31,     31,      62},      // Khordad
	{31,     31,      93},      // Tir
	{31,     31,      124},     // Mordad
	{31,     31,      155},     // Shahrivar
	{30,     30,      186},     // Mehr
	{30,     30,      216},     // Aban
	{30,     30,      246},     // Azar
	{30,     30,      276},     // Dey
	{30,     30,      306},     // Bahman
	{29,     30,      336},     // Esfand
}

var g_month_count = [...][...]int {
	{31,     31,      0,        0},       // Jan
	{28,     29,      31,       31},      // Feb
	{31,     31,      59,       60},      // Mar
	{30,     30,      90,       91},      // Apr
	{31,     31,      120,      121},     // May
	{30,     30,      151,      152},     // Jun
	{31,     31,      181,      182},     // Jul
	{31,     31,      212,      213},     // Aug
	{30,     30,      243,      244},     // Sep
	{31,     31,      273,      274},     // Oct
	{30,     30,      304,      305},     // Nov
	{31,     31,      334,      335},     // Dec
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
	// TODO convert time.Time (Gregorian) to Persian Time
	return nil
}

// Returns a new instance of time.Time from t.
func (t Time) Time() time.Time {
	// TODO convert Persian date to time.Time (Gregorian)
	return nil
}

// Returns a new instance of PersianDate.
func Date(year int, month Month, day, hour, min, sec, nsec int, loc *time.Location) Time {
	if loc == nil {
		panic("ptime: the Location must not be nil in call to Date")
	}

	return Time{year, month, day, hour, min, sec, nsec, loc}.normalize()
}

// Returns a new instance of PersianDate from unix timestamp.
// sec seconds and nsec nanoseconds since January 1, 1970 UTC.
func Unix(sec, nsec int64, loc *time.Location) Time {
	if loc == nil {
		panic("ptime: the Location must not be nil in call to Unix")
	}

	return Time(time.Unix(sec, nsec).In(loc))
}

// Returns unix timestamp (the number of seconds) of t.
func (t Time) Unix() int64 {
	return t.Time().Unix()
}

// Returns unix timestamp (the number of nanoseconds) of t.
func (t Time) UnixNano() int64 {
	return t.Time().UnixNano()
}

func Now(loc *time.Location) Time {
	if loc == nil {
		panic("ptime: the Location must not be nil in call to Now")
	}

	return Time(time.Now().In(loc))
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

// Returns the day in year of t.
func (t Time) YearDay() int {
	// TODO YearDay of PersianDate
	return 0
}

// Returns the weekday of t.
func (t Time) Weekday() Weekday {
	// TODO Weekday of PersianDate
	return 0
}

// Returns a new instance of Time for t+d.
func (t Time) Add(d time.Duration) Time {
	return Time(t.Time().Add(d))
}

// Returns true if the year of t is a leap year.
func (t Time) IsLeap() bool {
	return IsLeap(t.year)
}

func IsLeap(year int) bool {
	// TODO IsPersianLeap
	return false
}

// Normalizes the year, month and day if they were outside their usual ranges.
func (date Time) normalize() Time {
	// TODO Validate PersianDate
	return date
}