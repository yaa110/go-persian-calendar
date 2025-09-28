package ptime_test

import (
	"fmt"
	"runtime"
	"testing"
	"time"

	ptime "github.com/yaa110/go-persian-calendar"
)

type pMonthName struct {
	month ptime.Month
	name  string
}

type amPmName struct {
	ap   ptime.AmPm
	name string
}

type comparisonTest struct {
	name    string
	t1, t2  ptime.Time
	before  bool
	after   bool
	equal   bool
	compare int
}

type pdate struct {
	year  int
	month ptime.Month
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

type dayTime struct {
	hour    []int
	daytime ptime.DayTime
}

var monthPersianNames = []pMonthName{
	{ptime.Farvardin, "فروردین"},
	{ptime.Ordibehesht, "اردیبهشت"},
	{ptime.Khordad, "خرداد"},
	{ptime.Tir, "تیر"},
	{ptime.Mordad, "مرداد"},
	{ptime.Shahrivar, "شهریور"},
	{ptime.Mehr, "مهر"},
	{ptime.Aban, "آبان"},
	{ptime.Azar, "آذر"},
	{ptime.Dey, "دی"},
	{ptime.Bahman, "بهمن"},
	{ptime.Esfand, "اسفند"},
}

var monthDariNames = []pMonthName{
	{ptime.Hamal, "حمل"},
	{ptime.Sur, "ثور"},
	{ptime.Jauza, "جوزا"},
	{ptime.Saratan, "سرطان"},
	{ptime.Asad, "اسد"},
	{ptime.Sonboleh, "سنبله"},
	{ptime.Mizan, "میزان"},
	{ptime.Aqrab, "عقرب"},
	{ptime.Qos, "قوس"},
	{ptime.Jady, "جدی"},
	{ptime.Dolv, "دلو"},
	{ptime.Hut, "حوت"},
}

var amPmNames = []amPmName{
	{ptime.Am, "قبل از ظهر"},
	{ptime.Pm, "بعد از ظهر"},
}

var amPmSNames = []amPmName{
	{ptime.Am, "ق.ظ"},
	{ptime.Pm, "ب.ظ"},
}

var dateConversions = []dateConversion{
	{
		persian:   pdate{1383, ptime.Tir, 15},
		gregorian: gdate{2004, time.July, 5},
	},
	{
		persian:   pdate{1394, ptime.Dey, 11},
		gregorian: gdate{2016, time.January, 1},
	},
	{
		persian:   pdate{1394, ptime.Esfand, 9},
		gregorian: gdate{2016, time.February, 28},
	},
	{
		persian:   pdate{1394, ptime.Esfand, 11},
		gregorian: gdate{2016, time.March, 1},
	},
	{
		persian:   pdate{1394, ptime.Esfand, 29},
		gregorian: gdate{2016, time.March, 19},
	},
	{
		persian:   pdate{1395, ptime.Farvardin, 1},
		gregorian: gdate{2016, time.March, 20},
	},
	{
		persian:   pdate{1395, ptime.Farvardin, 2},
		gregorian: gdate{2016, time.March, 21},
	},
	{
		persian:   pdate{1395, ptime.Farvardin, 3},
		gregorian: gdate{2016, time.March, 22},
	},
	{
		persian:   pdate{1395, ptime.Dey, 11},
		gregorian: gdate{2016, time.December, 31},
	},
}

var dayFunctionsSlice = []dayFunctions{
	{
		pdate{1394, ptime.Tir, 31},
		pdate{1394, ptime.Mordad, 1},
	},
	{
		pdate{1394, ptime.Esfand, 29},
		pdate{1395, ptime.Farvardin, 1},
	},
	{
		pdate{1395, ptime.Esfand, 29},
		pdate{1395, ptime.Esfand, 30},
	},
	{
		pdate{1395, ptime.Ordibehesht, 12},
		pdate{1395, ptime.Ordibehesht, 13},
	},
}

var daytimes = []dayTime{
	{[]int{0, 1, 2}, ptime.Midnight},
	{[]int{3, 4, 5}, ptime.Dawn},
	{[]int{6, 7, 8}, ptime.Morning},
	{[]int{9, 10, 11}, ptime.BeforeNoon},
	{[]int{12, 13, 14}, ptime.Noon},
	{[]int{15, 16, 17}, ptime.AfterNoon},
	{[]int{18, 19, 20}, ptime.Evening},
	{[]int{21, 22, 23}, ptime.Night},
}

var comparisonTests = []comparisonTest{
	{
		name:    "same time equal",
		t1:      ptime.Date(1394, ptime.Mehr, 2, 12, 0, 0, 0, ptime.Iran()),
		t2:      ptime.Date(1394, ptime.Mehr, 2, 12, 0, 0, 0, ptime.Iran()),
		before:  false,
		after:   false,
		equal:   true,
		compare: 0,
	},
	{
		name:    "t1 before t2 by hour",
		t1:      ptime.Date(1394, ptime.Mehr, 2, 12, 0, 0, 0, ptime.Iran()),
		t2:      ptime.Date(1394, ptime.Mehr, 2, 13, 0, 0, 0, ptime.Iran()),
		before:  true,
		after:   false,
		equal:   false,
		compare: -1,
	},
	{
		name:    "t1 after t2 by hour",
		t1:      ptime.Date(1394, ptime.Mehr, 2, 13, 0, 0, 0, ptime.Iran()),
		t2:      ptime.Date(1394, ptime.Mehr, 2, 12, 0, 0, 0, ptime.Iran()),
		before:  false,
		after:   true,
		equal:   false,
		compare: 1,
	},
	{
		name:    "t1 before t2 by day",
		t1:      ptime.Date(1394, ptime.Mehr, 2, 12, 0, 0, 0, ptime.Iran()),
		t2:      ptime.Date(1394, ptime.Mehr, 3, 12, 0, 0, 0, ptime.Iran()),
		before:  true,
		after:   false,
		equal:   false,
		compare: -1,
	},
	{
		name:    "t1 after t2 by day",
		t1:      ptime.Date(1394, ptime.Mehr, 3, 12, 0, 0, 0, ptime.Iran()),
		t2:      ptime.Date(1394, ptime.Mehr, 2, 12, 0, 0, 0, ptime.Iran()),
		before:  false,
		after:   true,
		equal:   false,
		compare: 1,
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
	if ptime.Iran().String() != "Asia/Tehran" {
		t.Error(
			"For", "Iran",
			"expected", "Asia/Tehran",
			"got", ptime.Iran().String(),
		)
	}

	if ptime.Afghanistan().String() != "Asia/Kabul" {
		t.Error(
			"For", "Afghanistan",
			"expected", "Asia/Kabul",
			"got", ptime.Afghanistan().String(),
		)
	}

	loc, _ := time.LoadLocation("Asia/Baghdad")
	expected := loc.String()

	actual := ptime.Now().In(loc).Location().String()

	if actual != expected {
		t.Error(
			"For", "Baghdad",
			"expected", expected,
			"got", actual,
		)
	}
}

func TestPersianToGregorian(t *testing.T) {
	for _, p := range dateConversions {
		gt := ptime.Date(p.persian.year, p.persian.month, p.persian.day, 11, 59, 59, 0, ptime.Iran()).Time()

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
		pt := ptime.New(time.Date(p.gregorian.year, p.gregorian.month, p.gregorian.day, 11, 59, 59, 0, ptime.Iran()))

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
	pu := ptime.Now().Unix()

	tu := time.Now().Unix()

	if pu != tu {
		t.Error(
			"Expected", tu,
			"got", pu,
		)
	}
}

func TestFromUnixTimeStamp(t *testing.T) {
	tu := time.Now().Unix()
	now := ptime.Now()

	fu := ptime.Unix(tu, int64(now.Nanosecond()))

	if fu.String() != now.String() {
		t.Error(
			"Expected", now.String(),
			"got", fu.String(),
		)
	}
}

func TestYesterday(t *testing.T) {
	for _, p := range dayFunctionsSlice {
		day := ptime.Date(p.day2.year, p.day2.month, p.day2.day, 12, 59, 59, 0, ptime.Iran())

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
		day := ptime.Date(p.day1.year, p.day1.month, p.day1.day, 12, 59, 59, 0, ptime.Iran())

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
	ti := ptime.Date(1394, ptime.Mehr, 2, 12, 59, 59, 0, ptime.Iran())

	if ti.Weekday() != ptime.Panjshanbeh {
		t.Error(
			"For", "Weekday()",
			"expected", ptime.Panjshanbeh.String(),
			"got", ti.Weekday().String(),
		)
	}
}

func TestYearDay(t *testing.T) {
	ti := ptime.Date(1394, ptime.Mehr, 2, 12, 59, 59, 0, ptime.Iran())

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
	ti := ptime.Date(1394, ptime.Mehr, 2, 12, 59, 59, 0, ptime.Iran())

	if ti.RMonthDay() != 28 {
		t.Error(
			"For", "RMonthDay()",
			"expected", 28,
			"got", ti.RMonthDay(),
		)
	}
}

func TestFirstLast(t *testing.T) {
	ti := ptime.Date(1394, ptime.Mehr, 2, 12, 59, 59, 0, ptime.Iran())

	if ti.FirstMonthDay().Weekday() != ptime.Charshanbeh {
		t.Error(
			"For", "FirstMonthDay().Weekday()",
			"expected", ptime.Charshanbeh.String(),
			"got", ti.FirstMonthDay().Weekday(),
		)
	}

	if ti.BeginningOfMonth().Weekday() != ptime.Charshanbeh {
		t.Error(
			"For", "BeginningOfMonth().Weekday()",
			"expected", ptime.Charshanbeh.String(),
			"got", ti.BeginningOfMonth().Weekday(),
		)
	}

	if ti.LastMonthDay().Weekday() != ptime.Panjshanbeh {
		t.Error(
			"For", "LastMonthDay().Weekday()",
			"expected", ptime.Panjshanbeh.String(),
			"got", ti.LastMonthDay().Weekday(),
		)
	}

	if ti.FirstYearDay().Weekday() != ptime.Shanbeh {
		t.Error(
			"For", "FirstYearDay().Weekday()",
			"expected", ptime.Shanbeh.String(),
			"got", ti.FirstYearDay().Weekday(),
		)
	}

	if ti.BeginningOfYear().Weekday() != ptime.Shanbeh {
		t.Error(
			"For", "BeginningOfYear().Weekday()",
			"expected", ptime.Shanbeh.String(),
			"got", ti.BeginningOfYear().Weekday(),
		)
	}

	if ti.LastYearDay().Weekday() != ptime.Shanbeh {
		t.Error(
			"For", "LastYearDay().Weekday()",
			"expected", ptime.Shanbeh.String(),
			"got", ti.LastYearDay().Weekday(),
		)
	}

	if ti.LastWeekday().Weekday() != ptime.Jomeh {
		t.Error(
			"For", "LastWeekday().Weekday()",
			"expected", ptime.Jomeh.String(),
			"got", ti.LastWeekday().Weekday(),
		)
	}
}

func TestAddDate(t *testing.T) {
	ti := ptime.Date(1394, ptime.Mehr, 2, 12, 59, 59, 0, ptime.Iran())

	if ti.AddDate(0, 0, 20).Weekday() != ptime.Charshanbeh {
		t.Error(
			"For", "AddDate(0, 0, 20).Weekday()",
			"expected", ptime.Charshanbeh.String(),
			"got", ti.AddDate(0, 0, 20).Weekday(),
		)
	}

	if ti.AddDate(0, 1, 0).Weekday() != ptime.Shanbeh {
		t.Error(
			"For", "AddDate(0, 1, 0).Weekday()",
			"expected", ptime.Shanbeh.String(),
			"got", ti.AddDate(0, 1, 0).Weekday(),
		)
	}

	if ti.AddDate(2, 0, 0).Weekday() != ptime.Yekshanbeh {
		t.Error(
			"For", "AddDate(2, 0, 0).Weekday()",
			"expected", ptime.Yekshanbeh.String(),
			"got", ti.AddDate(2, 0, 0).Weekday(),
		)
	}
}

func TestWeeks(t *testing.T) {
	ti := ptime.Date(1394, ptime.Mehr, 2, 12, 59, 59, 0, ptime.Iran())

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

func TestPanic(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Error("the test has paniced")
		}
	}()

	for y := 0; y < 3000; y++ {
		for m := 0; m < 13; m++ {
			for d := 0; d < 32; d++ {
				//nolint:gosmopolitan
				ti := time.Date(y, time.Month(m), d, 0, 0, 0, 0, time.Local)
				pt := ptime.New(ti)
				_ = pt.Month().String()
			}
		}
	}

	pt := ptime.Time{}
	_ = pt.String()
}

func TestZero(t *testing.T) {
	pt := ptime.New(time.Time{})
	if !pt.IsZero() {
		t.Error("time must be zero")
	}
}

func TestComparison(t *testing.T) {
	for _, tt := range comparisonTests {
		t.Run(tt.name, func(t *testing.T) {
			// Test Before
			if got := tt.t1.Before(tt.t2); got != tt.before {
				t.Error(
					"For", tt.name,
					"Before expected", tt.before,
					"got", got,
				)
			}

			// Test After
			if got := tt.t1.After(tt.t2); got != tt.after {
				t.Error(
					"For", tt.name,
					"After expected", tt.after,
					"got", got,
				)
			}

			// Test Equal
			if got := tt.t1.Equal(tt.t2); got != tt.equal {
				t.Error(
					"For", tt.name,
					"Equal expected", tt.equal,
					"got", got,
				)
			}

			// Test Compare
			if got := tt.t1.Compare(tt.t2); got != tt.compare {
				t.Error(
					"For", tt.name,
					"Compare expected", tt.compare,
					"got", got,
				)
			}
		})
	}
}

func TestFormat(t *testing.T) {
	ti := ptime.Date(1394, ptime.Mehr, 2, 12, 59, 59, 50260050, ptime.Iran())

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

func TestTimeFormat(t *testing.T) {
	ti := ptime.Date(1394, 7, 2, 14, 7, 8, 52065090, ptime.Iran())

	s := ti.TimeFormat("2 Jan 2006")
	if s != "2 مهر 1394" {
		t.Error(
			"Expected", "2 مهر 1394",
			"got", s,
		)
	}

	tid := ptime.Date(1394, ptime.Mizan, 2, 12, 59, 59, 52065090, ptime.Afghanistan())

	s = tid.TimeFormat("2 Jan 2006")
	if s != "2 میزان 1394" {
		t.Error(
			"Expected", "2 میزان 1394",
			"got", s,
		)
	}

	vals := map[string]string{
		"2006":       "1394",
		"06":         "94",
		"01":         "07",
		"1":          "7",
		"Jan":        "مهر",
		"January":    "مهر",
		"02":         "02",
		"2":          "2",
		"_2":         " 2",
		"Mon":        "پ",
		"Monday":     "پنج\u200cشنبه",
		"03":         "02",
		"3":          "2",
		"15":         "14",
		"04":         "07",
		"4":          "7",
		"05":         "08",
		"5":          "8",
		".000":       ".520",
		".000000":    ".520650",
		".000000000": ".52065090",
		".999":       ".52",
		".999999":    ".52065",
		".999999999": ".5206509",
		"PM":         "بعد از ظهر",
		"pm":         "ب.ظ",
		"MST":        "Asia/Tehran",
		"-0700":      "+0330",
		"-07":        "+03",
		"-07:00":     "+03:30",
		"Z0700":      "+0330",
		"Z07:00":     "+03:30",
	}
	for k, v := range vals {
		if s := ti.TimeFormat(k); s != v {
			t.Error(
				"Expected", k+"=>"+v,
				"got", s,
			)
		}
	}
}

func TestHourName(t *testing.T) {
	for _, dayPart := range daytimes {
		for _, hour := range dayPart.hour {
			ti := ptime.Date(1394, 7, 2, hour, 7, 8, 52065090, ptime.Iran())

			daytime := ti.DayTime()
			if daytime != dayPart.daytime {
				t.Error("Expected ", dayPart.daytime, "got ", daytime)
			}
		}
	}
}

func BenchmarkFormat(b *testing.B) {
	now := ptime.Now()

	var s string

	for i := 0; i < b.N; i++ {
		s = now.Format("A D E H HH K KK MM MMM MMI MM RD R S W Z a dd d e hh h kk k mm m ns nr rw rd ss s w y yyyy yyy yy z")
	}

	runtime.KeepAlive(s)
}
