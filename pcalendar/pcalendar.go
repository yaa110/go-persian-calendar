// In the name of Allah

// Copyright (c) 2015 Navid Fathollahzade
// This source code is licensed under MIT license that can be found in the LICENSE file.

// Package ptime provides functionality for implementation of Persian (Jalali) Calendar.
package pcalendar

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
	yearday int
	weekday Weekday
}

// A PersianDate represents a day in Gregorian Calendar.
type GregorianDate struct {
	Year int
	Month time.Month
	Day int
}

// A PersianCalendar provides functionality for conversion of Persian (Jalali) and Gregorian Calendars.
type PersianCalendar struct {
	Persian PersianDate
	Gregorian GregorianDate
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

// Returns new instance of PersianCalendar using PersianDate.
func NewPersianDate(date PersianDate) (PersianCalendar, errors) {
	err := ValidatePersianDate(date)
	if err != nil {
		return nil, err
	}

	date.calculateDays()

	return PersianCalendar {
		date,
		toGregorian(date),
	}, nil
}

// Returns new instance of PersianCalendar using GregorianDate.
func NewGregorianDate(date GregorianDate) (PersianCalendar, errors) {
	err := ValidateGregorianDate(date)
	if err != nil {
		return nil, err
	}

	pdate := toPersian(date)
	pdate.calculateDays()

	return PersianCalendar {
		pdate,
		date,
	}, nil
}

// Changes the instance of PersianCalendar (pc) using Persian date.
func (pc *PersianCalendar) SetPersianDate(date PersianDate) errors {
	err := ValidatePersianDate(date)
	if err != nil {
		return err
	}

	date.calculateDays()

	pc.Persian = date
	pc.Gregorian = toGregorian(date)

	return nil
}

// Changes the instance of PersianCalendar (pc) using Gregorian date.
func (pc *PersianCalendar) SetGregorianDate(date GregorianDate) errors {
	err := ValidateGregorianDate(date)
	if err != nil {
		return err
	}

	pdate := toPersian(date)
	pdate.calculateDays()

	pc.Persian = pdate
	pc.Gregorian = date

	return nil
}

// Validates the PersianDate to represent a correct day.
func ValidatePersianDate(date PersianDate) errors {
	// TODO
	return nil
}

// Validates the GregorianDate to represent a correct day.
func ValidateGregorianDate(date GregorianDate) errors {
	// TODO
	return nil
}

func toPersian(date GregorianDate) PersianDate {
	// TODO
	return nil
}

func toGregorian(date PersianDate) GregorianDate {
	// TODO
	return nil
}

func (pd *PersianDate) calculateDays() {
	// TODO set year and week days of pd
}