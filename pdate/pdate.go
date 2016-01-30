// In the name of Allah

// Copyright (c) 2016 Navid Fathollahzade
// This source code is licensed under MIT license that can be found in the LICENSE file.

// Version: 0.1
// Please visit https://github.com/yaa110/go-persian-calendar for more information.

// Package pdate provides functionality for implementation of Persian (Jalali) Calendar.
package pdate

import "time"

// A Month specifies a month of the year in Persian calendar starting from 1.
type Month int

// A Weekday specifies a day of the week in Persian starting from 0.
type Weekday int

// A PersianDate represents a day in Persian (Jalali) Calendar.
type PersianDate struct {
	year int
	month Month
	day int
}

// A PersianDate represents a day in Gregorian Calendar.
type GregorianDate struct {
	year int
	month time.Month
	day int
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

// Returns a new instance of PersianDate.
func PersianDate(year int, month Month, day int) PersianDate {
	return PersianDate{year, month, day}.normalize()
}

// Returns a new instance of GregorianDate.
func GregorianDate(year int, month Month, day int) GregorianDate {
	return GregorianDate{year, month, day}.normalize()
}

// Returns a new instance of PersianDate from unix timestamp.
// seconds since January 1, 1970 UTC.
func Unix(seconds int64, location *time.Location) PersianDate {
	return GregorianDate(time.Unix(seconds, 0).In(location).Date()).PersianDate()
}

// Returns unix timestamp (the number of seconds) of date.
func (date PersianDate) Unix(hour, minute, seconds, location *time.Location) int64 {
	gdate := date.GregorianDate()
	return time.Date(gdate.year, gdate.month, gdate.day, hour, minute, seconds, 0, location).Unix()
}

func Now(location *time.Location) PersianDate {
	return GregorianDate(time.Now().In(location).Date()).PersianDate()
}

// Returns the year, month, day of date.
func (date PersianDate) Date() (int, Month, int) {
	return date.year, date.month, date.day
}

// Returns the year, month, day of date.
func (date GregorianDate) Date() (int, time.Month, int) {
	return date.year, date.month, date.day
}

// Returns the year of date.
func (date PersianDate) Year() int {
	return date.year
}

// Returns the year of date.
func (date GregorianDate) Year() int {
	return date.year
}

// Returns the month of date.
func (date PersianDate) Month() Month {
	return date.month
}

// Returns the month of date.
func (date GregorianDate) Month() time.Month {
	return date.month
}

// Returns the day in month of date.
func (date PersianDate) Day() int {
	return date.day
}

// Returns the day in month of date.
func (date GregorianDate) Day() int {
	return date.day
}

// Returns the day in year of date.
func (date PersianDate) YearDay() int {
	// TODO YearDay of PersianDate
	return 0
}

// Returns the weekday of date.
func (date PersianDate) Weekday() Weekday {
	// TODO Weekday of PersianDate
	return 0
}

// Returns a new instance of time.Time from date.
func (date GregorianDate) Time(hours, minutes, seconds, nanoseconds int, location *time.Location) time.Time {
	return time.Date(date.year, date.month, date.day, hours, minutes, seconds, nanoseconds, location)
}

// Returns a new instance of time.Time from date.
func (date PersianDate) Time(hours, minutes, seconds, nanoseconds int, location *time.Location) time.Time {
	return date.GregorianDate().Time(hours, minutes, seconds, nanoseconds, location)
}

// Returns a new instance of PersianDate representing the day after or before date.
func (date PersianDate) Add(days int) PersianDate {
	return Unix(date.Unix(12, 0, 0, time.UTC) + (days * 86400), time.UTC)
}

func (date GregorianDate) PersianDate() PersianDate {
	var year, month, day int
	// TODO ToPersianDate
	return PersianDate(year, month, day)
}

// Returns the weekday of date.
func (date PersianDate) GregorianDate() GregorianDate {
	var year, month, day int
	// TODO ToGregorianDate
	return GregorianDate(year, month, day)
}

// Returns true if the year of date is a leap year.
func (date PersianDate) IsLeap() bool {
	return IsPersianLeap(date.year)
}

// Returns true if the year of date is a leap year.
func (date GregorianDate) IsLeap() bool {
	return IsGregorianLeap(date.year)
}

func IsPersianLeap(year int) bool {
	// TODO IsPersianLeap
	return false
}

func IsGregorianLeap(year int) bool {
	// TODO IsGregorianLeap
	return false
}

// Normalizes the year, month and day if they were outside their usual ranges.
func (date PersianDate) normalize() PersianDate {
	// TODO Validate PersianDate
	return date
}

// Normalizes the year, month and day if they were outside their usual ranges.
func (date GregorianDate) normalize() GregorianDate {
	// TODO Validate GregorianDate
	return date
}