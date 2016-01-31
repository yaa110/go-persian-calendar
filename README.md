Persian Calendar
================
The implementation of the Persian (Solar Hijri) Calendar in Go (golang)

## About
**Persian Calendar v0.1** provides functionality for conversion among Persian and Gregorian calendars. A Julian calendar has been used as an interface for all conversions. This package is compatible with the package [time](https://golang.org/pkg/time). The package name is `ptime`. All months are available with both Iranian and Dari Persian names. This source code is licensed under MIT license that can be found in the LICENSE file.

## Installation
```
$ go get github.com/yaa110/go-persian-calendar/ptime/...
```

## Getting started
1- Import the package `ptime`. Most of the time you need to import `time` and `fmt` packages, too.

```go
import (
    "github.com/yaa110/go-persian-calendar/ptime"
    "time"
    "fmt"
)
```

2- Convert Gregorian calendar to Persian calendar

```go
// Create a new instance of time.Time
var t time.Time = time.Date(2016, time.January, 1, 12, 1, 1, 0, ptime.Iran)

// Get a new instance of ptime.Time using time.Time
pt := ptime.New(t)

// Get the date in Persian calendar
fmt.Println(pt.Date()) // output: 1394 دی 11
```

3- Convert Persian calendar to Gregorian calendar

```go
// Create a new instance of ptime.Time
var pt ptime.Time = ptime.Date(1394, ptime.Mehr, 2, 12, 59, 59, 0, ptime.Iran)

// Get a new instance of time.Time
t := pt.Time()

// Get the date in Gregorian calendar
fmt.Println(t.Date()) // output: 2015 September 24
```

4- Get current time

```go
// Get a new instance of ptime.Time representing current time
pt := ptime.Now(ptime.Iran)

// Get year, month, day
fmt.Println(pt.Date()) // output: 1394 بهمن 11
fmt.Println(pt.Year(), pt.Month(), pt.Day()) // output: 1394 بهمن 11

// Get hour, minute, second
fmt.Println(pt.Clock()) // output: 21 54 30
fmt.Println(pt.Hour(), pt.Minute(), pt.Second()) // output: 21 54 30

// Get Unix timestamp (the number of seconds since January 1, 1970 UTC)
fmt.Println(pt.Unix()) // output: 1454277270

// Get yesterday, today and tomorrow
fmt.Println(pt.Yesterday().Weekday()) // output: شنبه
fmt.Println(pt.Weekday()) // output: یکشنبه
fmt.Println(pt.Tomorrow().Weekday()) // output: دوشنبه
```

5- Format the time

```go
// Get a new instance of ptime.Time using Unix timestamp
pt := ptime.Unix(1454277270, 0, ptime.Iran)

pt.Format("yyyy/MM/dd E hh:mm:ss a") // output: 1394/11/11 یکشنبه 09:54:30 ب.ظ

// yyyy, yyy, y     year (e.g. 1394)
// yy               2-digits representation of year (e.g. 94)
// MMM              the Persian name of month (e.g. فروردین)
// MMI              the Dari name of month (e.g. حمل)
// MM               2-digits representation of month (e.g. 01)
// M                month (e.g. 1)
// rw               remaining weeks of year
// w                week of year
// RW               remaining weeks of month
// W                week of month
// RD               remaining days of year
// D                day of year
// rd               remaining days of month
// dd               2-digits representation of day (e.g. 01)
// d                day (e.g. 1)
// E                the Persian name of weekday (e.g. شنبه)
// e                the Persian short name of weekday (e.g. ش)
// A                the Persian name of 12-Hour marker (e.g. قبل از ظهر)
// a                the Persian short name of 12-Hour marker (e.g. ق.ظ)
// HH               2-digits representation of hour [00-23]
// H                hour [0-23]
// kk               2-digits representation of hour [01-24]
// k                hour [1-24]
// hh               2-digits representation of hour [01-12]
// h                hour [1-12]
// KK               2-digits representation of hour [00-11]
// K                hour [0-11]
// mm               2-digits representation of minute [00-59]
// m                minute [0-59]
// ss               2-digits representation of seconds [00-59]
// s                seconds [0-59]
// ns               nanoseconds
// S                3-digits representation of milliseconds (e.g. 001)
// z                the name of location
// Z                zone offset (e.g. +03:30)
```

## Methods of `ptime.Time`
- **String()** RFC3339Nano format of `ptime.Time`.
- **Time()** Converts Persian date to Gregorian date and returns a new instance of `time.Time`
- **Unix()** Returns the number of seconds since January 1, 1970 UTC.
- **Date()** Returns the year, month, day of `ptime.Time`.
- **Clock()** Returns the hour, minute, seconds offsets of `ptime.Time`.
- **YearDay()** Returns the day of year of `ptime.Time`.
- **RYearDay()** Returns the number of remaining days of the year of `ptime.Time`.
- **RMonthDay()** Returns the number of remaining days of the month of `ptime.Time`.
- **FirstWeekDay()** Returns a new instance of Time representing the first day of the week of `ptime.Time`.
- **LastWeekday()** Returns a new instance of Time representing the last day of the week of `ptime.Time`.
- **FirstMonthDay()** Returns a new instance of Time representing the first day of the month of `ptime.Time`.
- **LastMonthDay()** Returns a new instance of Time representing the last day of the month of `ptime.Time`.
- **FirstYearDay()** Returns a new instance of Time representing the first day of the year of `ptime.Time`.
- **LastYearDay()** Returns a new instance of Time representing the last day of the year of `ptime.Time`.
- **MonthWeek()** Returns the week of month of `ptime.Time`.
- **RMonthWeek()** Returns the number of remaining weeks of the month of `ptime.Time`.
- **YearWeek()** Returns the week of year of `ptime.Time`.
- **RYearWeek()** Returns the number of remaining weeks of the year of `ptime.Time`.
- **Yesterday()** Returns a new instance of `ptime.Time` representing a day before the day of another `ptime.Time`.
- **Tomorrow()** Returns a new instance of `ptime.Time` representing a day after the day of another `ptime.Time`.
- **IsLeap()** Returns true if the year of `ptime.Time` is a leap year.
- **Zone()** Returns the zone name and its offset in seconds east of UTC of `ptime.Time`.
Use `go doc` for more methods and functionality available for `ptime.Time`, `ptime.Month`, `ptime.Weekday` and `AM_PM`.