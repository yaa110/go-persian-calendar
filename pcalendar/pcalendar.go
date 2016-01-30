// In the name of Allah

// Copyright (c) 2015 Navid Fathollahzade
// This source code is licensed under MIT license that can be found in the LICENSE file.

// Package ptime provides functionality for implementation of Persian (Jalali) Calendar.
package pcalendar

import "errors"

// A PersianMonth specifies a month of the year in Persian calendar starting from 1.
type PersianMonth int

// A GregorianMonth specifies a month of the year in Gregorian calendar starting from 1.
type GregorianMonth int

// A PersianWeekday specifies a day of the week in Persian calendar starting from 0.
type PersianWeekday int

// A GregorianWeekday specifies a day of the week in Gregorian calendar starting from 0.
type GregorianWeekday int

// A PersianCalendar represents a day in Persian (Jalali) Calendar.
type PersianCalendar struct {
	p_year int
	p_month PersianMonth
	p_month_day int
	p_year_day int
	p_week_day PersianWeekday
	g_year int
	g_month GregorianMonth
	g_month_day int
	g_year_day int
	g_week_day GregorianWeekday
}

const (
	Farvardin PersianMonth = 1 + iota
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
	January GregorianMonth = 1 + iota
	February
	March
	April
	May
	June
	July
	August
	September
	October
	November
	December
)

const (
	Shanbe PersianWeekday = iota
	Yekshanbe
	Doshanbe
	Seshanbe
	Chaharshanbe
	Panjshanbe
	Jome
)

const (
	Saturday GregorianWeekday = iota
	Sunday
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
)

const (
	persian_epoch = 226899

	month_count_normal = 0
	month_count_leap = 1
	month_count_normal_before = 2
	month_count_leap_before = 3
)

var p_months = [...]string{
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

var g_months = [...]string{
	"January",
	"February",
	"March",
	"April",
	"May",
	"June",
	"July",
	"August",
	"September",
	"October",
	"November",
	"December",
}

var p_days = [...]string{
	"شنبه",
	"یک‌شنبه",
	"دوشنبه",
	"سه‌شنبه",
	"چهارشنبه",
	"پنج‌شنبه",
	"جمعه",
}

var g_days = [...]string{
	"Saturday",
	"Sunday",
	"Monday",
	"Tuesday",
	"Wednesday",
	"Thursday",
	"Friday",
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
func (m PersianMonth) String() string {
	return p_months[m - 1]
}

// Returns the English name of the month.
func (m GregorianMonth) String() string {
	return g_months[m - 1]
}

// Returns the Persian name of the day in week.
func (d PersianWeekday) String() string {
	return p_days[d]
}

// Returns the English name of the day in week.
func (d GregorianWeekday) String() string {
	return g_days[d]
}

// Returns new instance of PersianCalendar using Persian date.
func NewPersianDate(year int, month PersianMonth, day int) (PersianCalendar, errors) {
	err := checkPersianDate(year, month, day)
	if err != nil {
		return nil, err
	}

	gy, gm, gmd, gyd, gwd := toGregorian(year, month, day)
	pyd := persianYearDay(year, month, day)

	return PersianCalendar {
		year,   month,  day,    pyd,    PersianWeekday(gwd),
		gy,     gm,     gmd,    gyd,    gwd,
	}, nil
}

// Returns new instance of PersianCalendar using Gregorian date.
func NewGregorianDate(year int, month GregorianMonth, day int) (PersianCalendar, errors) {
	err := checkGregorianDate(year, month, day)
	if err != nil {
		return nil, err
	}

	py, pm, pmd, pyd, pwd := toPersian(year, month, day)
	gyd := gregorianYearDay(year, month, day)

	return PersianCalendar {
		py,     pm,     pmd,    pyd,    pwd,
		year,   month,  day,    gyd,    GregorianWeekday(pwd),
	}, nil
}

// Changes the instance of PersianCalendar (pc) using Persian date.
func (pc PersianCalendar) SetPersianDate(year int, month PersianMonth, day int) errors {
	err := checkPersianDate(year, month, day)
	if err != nil {
		return err
	}

	pc.p_year = year
	pc.p_month = month
	pc.p_month_day = day

	pc.g_year, pc.g_month, pc.g_month_day, pc.g_year_day, pc.g_week_day = toGregorian(year, month, day)
	pc.p_year_day = persianYearDay(year, month, day)
	pc.p_week_day = PersianWeekday(pc.g_week_day)

	return nil
}

// Changes the instance of PersianCalendar (pc) using Gregorian date.
func (pc PersianCalendar) SetGregorianDate(year int, month GregorianMonth, day int) errors {
	err := checkGregorianDate(year, month, day)
	if err != nil {
		return err
	}

	pc.g_year = year
	pc.g_month = month
	pc.g_month_day = day

	pc.p_year, pc.p_month, pc.p_month_day, pc.p_year_day, pc.p_week_day = toPersian(year, month, day)
	pc.g_year_day = gregorianYearDay(year, month, day)
	pc.g_week_day = GregorianWeekday(pc.p_week_day)

	return nil
}

// Checks Persian date inputs to be a correct day.
func checkPersianDate(year int, month PersianMonth, day int) errors {
	// TODO
	return nil
}

// Checks Gregorian date inputs to be a correct day.
func checkGregorianDate(year int, month GregorianMonth, day int) errors {
	// TODO
	return nil
}

// Converts Gregorian date to Persian date (year, month, month_day, year_day, week_day)
func toPersian(year int, month GregorianMonth, day int) (y int, m PersianMonth, md int, yd int, wd PersianWeekday) {
	// TODO
	return
}

// Converts Persian date to Gregorian date (year, month, month_day, year_day, week_day)
func toGregorian(year int, month PersianMonth, day int) (y int, m GregorianMonth, md int, yd int, wd GregorianWeekday) {
	// TODO
	return
}

// Returns the year day of Persian date.
func persianYearDay(year int, month PersianMonth, day int) (yd int) {
	// TODO
	return
}

// Returns the year day of Gregorian date.
func gregorianYearDay(year int, month GregorianMonth, day int) (yd int) {
	// TODO
	return
}