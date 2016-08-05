package ptime_test

import (
	"fmt"
	. "github.com/yaa110/go-persian-calendar/ptime"
	"testing"
	"time"
)

type pMonthName struct {
	month Month
	name  string
}

type amPmName struct {
	ap   AmPm
	name string
}

type pdate struct {
	year  int
	month Month
	day   int
}

type gdate struct {
	year  int
	month time.Month
	day   int
}

type dateConversion struct {
	persian   pdate
	gregorian gdate
}

type dayFunctions struct {
	day1 pdate
	day2 pdate
}

var monthPersianNames = []pMonthName{
	{Farvardin, "فروردین"},
	{Ordibehesht, "اردیبهشت"},
	{Khordad, "خرداد"},
	{Tir, "تیر"},
	{Mordad, "مرداد"},
	{Shahrivar, "شهریور"},
	{Mehr, "مهر"},
	{Aban, "آبان"},
	{Azar, "آذر"},
	{Dey, "دی"},
	{Bahman, "بهمن"},
	{Esfand, "اسفند"},
}

var monthDariNames = []pMonthName{
	{Hamal, "حمل"},
	{Sur, "ثور"},
	{Jauza, "جوزا"},
	{Saratan, "سرطان"},
	{Asad, "اسد"},
	{Sonboleh, "سنبله"},
	{Mizan, "میزان"},
	{Aqrab, "عقرب"},
	{Qos, "قوس"},
	{Jady, "جدی"},
	{Dolv, "دلو"},
	{Hut, "حوت"},
}

var amPmNames = []amPmName{
	{Am, "قبل از ظهر"},
	{Pm, "بعد از ظهر"},
}

var amPmSNames = []amPmName{
	{Am, "ق.ظ"},
	{Pm, "ب.ظ"},
}

var dateConversions = []dateConversion{
	{
		persian:   pdate{1383, Tir, 15},
		gregorian: gdate{2004, time.July, 5},
	},
	{
		persian:   pdate{1394, Dey, 11},
		gregorian: gdate{2016, time.January, 1},
	},
	{
		persian:   pdate{1394, Esfand, 9},
		gregorian: gdate{2016, time.February, 28},
	},
	{
		persian:   pdate{1394, Esfand, 11},
		gregorian: gdate{2016, time.March, 1},
	},
	{
		persian:   pdate{1394, Esfand, 29},
		gregorian: gdate{2016, time.March, 19},
	},
	{
		persian:   pdate{1395, Farvardin, 1},
		gregorian: gdate{2016, time.March, 20},
	},
	{
		persian:   pdate{1395, Farvardin, 2},
		gregorian: gdate{2016, time.March, 21},
	},
	{
		persian:   pdate{1395, Farvardin, 3},
		gregorian: gdate{2016, time.March, 22},
	},
	{
		persian:   pdate{1395, Dey, 11},
		gregorian: gdate{2016, time.December, 31},
	},
}

var dayFunctionsSlice = []dayFunctions{
	{
		pdate{1394, Tir, 31},
		pdate{1394, Mordad, 1},
	},
	{
		pdate{1394, Esfand, 29},
		pdate{1395, Farvardin, 1},
	},
	{
		pdate{1395, Esfand, 29},
		pdate{1395, Esfand, 30},
	},
	{
		pdate{1395, Ordibehesht, 12},
		pdate{1395, Ordibehesht, 13},
	},
}

func TestPersianMonthName(t *testing.T) {
	for _, p := range monthPersianNames {
		if p.month.String() != p.name {
			t.Error(
				"Expected", p.name,
				"got", p.month.String(),
			)
		}
	}
}

func TestDariMonthName(t *testing.T) {
	for _, p := range monthDariNames {
		if p.month.Dari() != p.name {
			t.Error(
				"Expected", p.name,
				"got", p.month.Dari(),
			)
		}
	}
}

func TestAmPmName(t *testing.T) {
	for _, p := range amPmNames {
		if p.ap.String() != p.name {
			t.Error(
				"Expected", p.name,
				"got", p.ap.String(),
			)
		}
	}
}

func TestAmPmShortName(t *testing.T) {
	for _, p := range amPmSNames {
		if p.ap.Short() != p.name {
			t.Error(
				"Expected", p.name,
				"got", p.ap.Short(),
			)
		}
	}
}

func TestLocations(t *testing.T) {
	if Iran().String() != "Asia/Tehran" {
		t.Error(
			"For", "Iran",
			"expected", "Asia/Tehran",
			"got", Iran().String(),
		)
	}

	if Afghanistan().String() != "Asia/Kabul" {
		t.Error(
			"For", "Afghanistan",
			"expected", "Asia/Kabul",
			"got", Afghanistan().String(),
		)
	}
}

func TestPersianToGregorian(t *testing.T) {
	for _, p := range dateConversions {
		gt := Date(p.persian.year, p.persian.month, p.persian.day, 11, 59, 59, 0, Iran()).Time()

		if gt.Year() != p.gregorian.year || gt.Month() != p.gregorian.month || gt.Day() != p.gregorian.day {
			t.Error(
				"For", fmt.Sprintf("%d %s %d", p.persian.year, p.persian.month.String(), p.persian.day),
				"expected", fmt.Sprintf("%d %s %d", p.gregorian.year, p.gregorian.month.String(), p.gregorian.day),
				"got", fmt.Sprintf("%d %s %d", gt.Year(), gt.Month().String(), gt.Day()),
			)
		}
	}
}

func TestGregorianToPersian(t *testing.T) {
	for _, p := range dateConversions {
		pt := New(time.Date(p.gregorian.year, p.gregorian.month, p.gregorian.day, 11, 59, 59, 0, Iran()))

		if pt.Year() != p.persian.year || pt.Month() != p.persian.month || pt.Day() != p.persian.day {
			t.Error(
				"For", fmt.Sprintf("%d %s %d", p.gregorian.year, p.gregorian.month.String(), p.gregorian.day),
				"expected", fmt.Sprintf("%d %s %d", p.persian.year, p.persian.month.String(), p.persian.day),
				"got", fmt.Sprintf("%d %s %d", pt.Year(), pt.Month().String(), pt.Day()),
			)
		}
	}
}

func TestToUnixTimeStamp(t *testing.T) {
	pu := Now(Iran()).Unix()
	tu := time.Now().In(Iran()).Unix()
	if pu != tu {
		t.Error(
			"Expected", tu,
			"got", pu,
		)
	}
}

func TestFromUnixTimeStamp(t *testing.T) {
	tu := time.Now().In(Iran()).Unix()
	now := Now(Iran())

	fu := Unix(tu, int64(now.Nanosecond()), Iran())

	if fu.String() != now.String() {
		t.Error(
			"Expected", now.String(),
			"got", fu.String(),
		)
	}
}

func TestYesterday(t *testing.T) {
	for _, p := range dayFunctionsSlice {
		day := Date(p.day2.year, p.day2.month, p.day2.day, 12, 59, 59, 0, Iran())
		yesterday := day.Yesterday()
		if yesterday.Year() != p.day1.year || yesterday.Month() != p.day1.month || yesterday.Day() != p.day1.day {
			t.Error(
				"For", fmt.Sprintf("%d %s %d", p.day2.year, p.day2.month.String(), p.day2.day),
				"expected", fmt.Sprintf("%d %s %d", p.day1.year, p.day1.month.String(), p.day1.day),
				"got", fmt.Sprintf("%d %s %d", yesterday.Year(), yesterday.Month().String(), yesterday.Day()),
			)
		}
	}
}

func TestTomorrow(t *testing.T) {
	for _, p := range dayFunctionsSlice {
		day := Date(p.day1.year, p.day1.month, p.day1.day, 12, 59, 59, 0, Iran())
		tomorrow := day.Tomorrow()
		if tomorrow.Year() != p.day2.year || tomorrow.Month() != p.day2.month || tomorrow.Day() != p.day2.day {
			t.Error(
				"For", fmt.Sprintf("%d %s %d", p.day1.year, p.day1.month.String(), p.day1.day),
				"expected", fmt.Sprintf("%d %s %d", p.day2.year, p.day2.month.String(), p.day2.day),
				"got", fmt.Sprintf("%d %s %d", tomorrow.Year(), tomorrow.Month().String(), tomorrow.Day()),
			)
		}
	}
}

func TestWeekday(t *testing.T) {
	ti := Date(1394, Mehr, 2, 12, 59, 59, 0, Iran())

	if ti.Weekday() != Panjshanbeh {
		t.Error(
			"For", "Weekday()",
			"expected", Panjshanbeh.String(),
			"got", ti.Weekday().String(),
		)
	}
}

func TestYearDay(t *testing.T) {
	ti := Date(1394, Mehr, 2, 12, 59, 59, 0, Iran())

	if ti.YearDay() != 188 {
		t.Error(
			"For", "YearDay()",
			"expected", 188,
			"got", ti.YearDay(),
		)
	}

	if ti.RYearDay() != 177 {
		t.Error(
			"For", "RYearDay()",
			"expected", 177,
			"got", ti.RYearDay(),
		)
	}
}

func TestRMonthDay(t *testing.T) {
	ti := Date(1394, Mehr, 2, 12, 59, 59, 0, Iran())

	if ti.RMonthDay() != 28 {
		t.Error(
			"For", "RMonthDay()",
			"expected", 28,
			"got", ti.RMonthDay(),
		)
	}
}

func TestFirstLast(t *testing.T) {
	ti := Date(1394, Mehr, 2, 12, 59, 59, 0, Iran())

	if ti.FirstMonthDay().Weekday() != Charshanbeh {
		t.Error(
			"For", "FirstMonthDay().Weekday()",
			"expected", Charshanbeh.String(),
			"got", ti.FirstMonthDay().Weekday(),
		)
	}

	if ti.LastMonthDay().Weekday() != Panjshanbeh {
		t.Error(
			"For", "LastMonthDay().Weekday()",
			"expected", Panjshanbeh.String(),
			"got", ti.LastMonthDay().Weekday(),
		)
	}

	if ti.FirstYearDay().Weekday() != Shanbeh {
		t.Error(
			"For", "FirstYearDay().Weekday()",
			"expected", Shanbeh.String(),
			"got", ti.FirstYearDay().Weekday(),
		)
	}

	if ti.LastYearDay().Weekday() != Shanbeh {
		t.Error(
			"For", "LastYearDay().Weekday()",
			"expected", Shanbeh.String(),
			"got", ti.LastYearDay().Weekday(),
		)
	}

	if ti.LastWeekday().Weekday() != Jomeh {
		t.Error(
			"For", "LastWeekday().Weekday()",
			"expected", Jomeh.String(),
			"got", ti.LastWeekday().Weekday(),
		)
	}
}

func TestAddDate(t *testing.T) {
	ti := Date(1394, Mehr, 2, 12, 59, 59, 0, Iran())

	if ti.AddDate(0, 0, 20).Weekday() != Charshanbeh {
		t.Error(
			"For", "AddDate(0, 0, 20).Weekday()",
			"expected", Charshanbeh.String(),
			"got", ti.AddDate(0, 0, 20).Weekday(),
		)
	}

	if ti.AddDate(0, 1, 0).Weekday() != Shanbeh {
		t.Error(
			"For", "AddDate(0, 1, 0).Weekday()",
			"expected", Shanbeh.String(),
			"got", ti.AddDate(0, 1, 0).Weekday(),
		)
	}

	if ti.AddDate(2, 0, 0).Weekday() != Yekshanbeh {
		t.Error(
			"For", "AddDate(2, 0, 0).Weekday()",
			"expected", Yekshanbeh.String(),
			"got", ti.AddDate(2, 0, 0).Weekday(),
		)
	}
}

func TestWeeks(t *testing.T) {
	ti := Date(1394, Mehr, 2, 12, 59, 59, 0, Iran())

	if ti.YearWeek() != 27 {
		t.Error(
			"For", "YearWeek()",
			"expected", 27,
			"got", ti.YearWeek(),
		)
	}

	if ti.RYearWeek() != 25 {
		t.Error(
			"For", "RYearWeek()",
			"expected", 25,
			"got", ti.RYearWeek(),
		)
	}

	if ti.MonthWeek() != 1 {
		t.Error(
			"For", "MonthWeek()",
			"expected", 1,
			"got", ti.MonthWeek(),
		)
	}

	if ti.IsLeap() {
		t.Error(
			"For", "IsLeap()",
			"expected", false,
			"got", ti.IsLeap(),
		)
	}
}

func TestFormat(t *testing.T) {
	ti := Date(1394, Mehr, 2, 12, 59, 59, 50260050, Iran())

	s := ti.Format("d MMM yyyy")
	if s != "2 مهر 1394" {
		t.Error(
			"Expected", "2 مهر 1394",
			"got", s,
		)
	}

	s = ti.Format("d MMI yyyy")
	if s != "2 میزان 1394" {
		t.Error(
			"Expected", "2 میزان 1394",
			"got", s,
		)
	}

	s = ti.Format("yyyy yyy yy y MM M dd d HH H kk k hh h KK K S ns Z")
	if s != "1394 1394 94 1394 07 7 02 2 12 12 12 12 12 12 00 0 050 50260050 +03:30" {
		t.Error(
			"Expected", "1394 1394 94 1394 07 7 02 2 12 12 12 12 12 12 00 0 050 50260050 +03:30",
			"got", s,
		)
	}
}
