// In the name of Allah

// Copyright (c) 2016 Navid Fathollahzade
// This source code is licensed under MIT license that can be found in the LICENSE file.

// Version: 0.1
// Please visit https://github.com/yaa110/go-persian-calendar for more information.

// Package pdate provides functionality for implementation of Persian (Jalali) Calendar.
package pdate

import (
	"errors"
	"time"
)

// A Month specifies a month of the year in Persian calendar starting from 1.
type Month int

// A Weekday specifies a day of the week in Persian starting from 0.
type Weekday int

// A PersianDate represents a day in Persian (Jalali) Calendar.
type PersianDate struct {
	Year int
	Month Month
	Day int
}

// A PersianDate represents a day in Gregorian Calendar.
type GregorianDate struct {
	Year int
	Month time.Month
	Day int
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

// Validates the date to represent a correct day.
func (date PersianDate) Validate() errors {
	// TODO
	return nil
}

// Validates the date to represent a correct day.
func (date GregorianDate) Validate() errors {
	// TODO
	return nil
}

// Returns the day in year of date.
func (date PersianDate) YearDay() int {
	// TODO set year and week days of pd
	return 0
}

// Returns the weekday of date.
func (date PersianDate) Weekday() int {
	// TODO set year and week days of pd
	return 0
}

func (date GregorianDate) ToPersianDate() PersianDate {
	// TODO
	return nil
}

// Returns the weekday of date.
func (date PersianDate) ToGregorianDate() GregorianDate {
	// TODO
	return nil
}