package ptime

import (
	"fmt"
	"testing"
)

func TestCalculateJDNToGregorian(t *testing.T) {
	for _, tc := range julianToGregorianMapping {
		t.Run(fmt.Sprintf("%d/%d/%d", tc.gregorianDate.year, tc.gregorianDate.month, tc.gregorianDate.day), func(t *testing.T) {
			// Calculate JDN from the Gregorian testDate
			calculatedJDN := convertGregorianPostReformToJDN(tc.gregorianDate.year, tc.gregorianDate.month, tc.gregorianDate.day)

			// Compare the calculated JDN with the expected JDN
			if calculatedJDN != tc.julianDay {
				t.Errorf("Test failed for testDate %d-%d-%d: expected JDN %d, got %d\n",
					tc.gregorianDate.year, tc.gregorianDate.month, tc.gregorianDate.day, tc.julianDay, calculatedJDN)
			}
		})

	}
}

func TestCalculateGregorianToJDN(t *testing.T) {
	for _, tc := range julianToGregorianMapping {
		t.Run(fmt.Sprintf("%d/%d/%d", tc.gregorianDate.year, tc.gregorianDate.month, tc.gregorianDate.day), func(t *testing.T) {
			// calculate shamsi testDate by jdn
			year, month, day := convertJDNToGregorianPostReform(tc.julianDay)

			if year != tc.gregorianDate.year {
				t.Errorf("Test failed for testDate %d: expected Year %d, got %d\n", tc.julianDay, tc.gregorianDate.year, year)
			}

			if month != tc.gregorianDate.month {
				t.Errorf("Test failed for testDate %d: expected Month %d, got %d\n", tc.julianDay, tc.gregorianDate.month, month)
			}

			if day != tc.gregorianDate.day {
				t.Errorf("Test failed for testDate %d: expected Day %d, got %d\n", tc.julianDay, tc.gregorianDate.day, day)
			}
		})

	}
}

func TestCalculateJDNToShamsi(t *testing.T) {
	for _, tc := range julianToShamsiMapping {
		t.Run(fmt.Sprintf("%d/%d/%d", tc.shamsiDate.year, tc.shamsiDate.month, tc.shamsiDate.day), func(t *testing.T) {
			// calculate shamsi testDate by jdn
			year, month, day := convertJDNToShamsi(tc.julianDay)

			if year != tc.shamsiDate.year {
				t.Errorf("Test failed for testDate %d: expected Year %d, got %d\n", tc.julianDay, tc.shamsiDate.year, year)
			}

			if month != tc.shamsiDate.month {
				t.Errorf("Test failed for testDate %d: expected Month %d, got %d\n", tc.julianDay, tc.shamsiDate.month, month)
			}

			if day != tc.shamsiDate.day {
				t.Errorf("Test failed for testDate %d: expected Day %d, got %d\n", tc.julianDay, tc.shamsiDate.day, day)
			}
		})

	}
}

func TestCalculateShamsiToJDN(t *testing.T) {
	for _, tc := range julianToShamsiMapping {
		t.Run(fmt.Sprintf("%d/%d/%d", tc.shamsiDate.year, tc.shamsiDate.month, tc.shamsiDate.day), func(t *testing.T) {
			// calculate shamsi testDate by jdn
			julianDay := convertShamsiToJDN(tc.shamsiDate.year, tc.shamsiDate.month, tc.shamsiDate.day)

			if julianDay != tc.julianDay {
				t.Errorf("Test failed for testDate %d-%d-%d: expected JDN %d, got %d\n",
					tc.shamsiDate.year, tc.shamsiDate.month, tc.shamsiDate.day, tc.julianDay, julianDay)
			}
		})
	}

}

type testDate struct {
	year, month, day int
}

// also these test cases where calculated by hand and the reference of time.ir to check for validation of them
var julianToShamsiMapping = []struct {
	julianDay  int
	shamsiDate testDate
}{
	{
		shamsiDate: testDate{year: 1403, month: 12, day: 30},
		julianDay:  2460755,
	},
	{
		shamsiDate: testDate{year: 1400, month: 6, day: 15},
		julianDay:  2459464,
	},
	{
		shamsiDate: testDate{year: 1395, month: 1, day: 1},
		julianDay:  2457468,
	},
	{
		shamsiDate: testDate{year: 1422, month: 12, day: 29},
		julianDay:  2467694,
	},
	{
		shamsiDate: testDate{year: 1388, month: 7, day: 12},
		julianDay:  2455109,
	},
	{
		shamsiDate: testDate{year: 1415, month: 4, day: 5},
		julianDay:  2464870,
	},
	{
		shamsiDate: testDate{year: 1390, month: 10, day: 20},
		julianDay:  2455937,
	},
	{
		shamsiDate: testDate{year: 1435, month: 2, day: 9},
		julianDay:  2472117,
	},
	{
		shamsiDate: testDate{year: 1377, month: 5, day: 3},
		julianDay:  2451020,
	},
	{
		shamsiDate: testDate{year: 1405, month: 7, day: 25},
		julianDay:  2461331,
	},
	{
		shamsiDate: testDate{year: 1399, month: 12, day: 29},
		julianDay:  2459293,
	},
	{
		shamsiDate: testDate{year: 1385, month: 8, day: 15},
		julianDay:  2454046,
	},
	{
		shamsiDate: testDate{year: 1418, month: 3, day: 7},
		julianDay:  2465937,
	},
	{
		shamsiDate: testDate{year: 1372, month: 2, day: 1},
		julianDay:  2449099,
	},
	{
		shamsiDate: testDate{year: 1429, month: 11, day: 12},
		julianDay:  2470204,
	},
	{
		shamsiDate: testDate{year: 1407, month: 6, day: 8},
		julianDay:  2462013,
	},
	{
		shamsiDate: testDate{year: 1393, month: 4, day: 25},
		julianDay:  2456855,
	},
	{
		shamsiDate: testDate{year: 1437, month: 7, day: 3},
		julianDay:  2472997,
	},
	{
		shamsiDate: testDate{year: 1380, month: 1, day: 1},
		julianDay:  2451990,
	},
	{
		shamsiDate: testDate{year: 1410, month: 9, day: 18},
		julianDay:  2463210,
	},
	{
		shamsiDate: testDate{year: 1398, month: 11, day: 29},
		julianDay:  2458898,
	},
	{
		shamsiDate: testDate{year: 1376, month: 12, day: 29}, // Leap year
		julianDay:  2450893,
	},
	{
		shamsiDate: testDate{year: 1397, month: 6, day: 20},
		julianDay:  2458373,
	},
	{
		shamsiDate: testDate{year: 1414, month: 12, day: 29}, // Leap year
		julianDay:  2464772,
	},
	{
		shamsiDate: testDate{year: 1388, month: 11, day: 10},
		julianDay:  2455227,
	},
	{
		shamsiDate: testDate{year: 1400, month: 1, day: 1}, // Leap year
		julianDay:  2459295,
	},
	{
		shamsiDate: testDate{year: 1422, month: 7, day: 25},
		julianDay:  2467540,
	},
	{
		shamsiDate: testDate{year: 1375, month: 9, day: 8},
		julianDay:  2450416,
	},
	{
		shamsiDate: testDate{year: 1396, month: 2, day: 20}, // Leap year
		julianDay:  2457884,
	},
	{
		shamsiDate: testDate{year: 1411, month: 3, day: 1},
		julianDay:  2463374,
	},
	{
		shamsiDate: testDate{year: 1382, month: 8, day: 20},
		julianDay:  2452955,
	},
	{
		shamsiDate: testDate{year: 1378, month: 11, day: 24},
		julianDay:  2451588,
	},
	{
		shamsiDate: testDate{year: 1348, month: 10, day: 11},
		julianDay:  2440588,
	},
}

// the reference for these test cases are https://aa.usno.navy.mil/data/JulianDate
var julianToGregorianMapping = []struct {
	julianDay     int
	gregorianDate testDate
}{
	{
		julianDay:     2460755, // 1403/12/30
		gregorianDate: testDate{year: 2025, month: 3, day: 20},
	},
	{
		julianDay:     2451588,
		gregorianDate: testDate{year: 2000, month: 2, day: 13},
	},
	{
		julianDay:     2440588,
		gregorianDate: testDate{year: 1970, month: 1, day: 1},
	},
	{
		julianDay:     2406842,
		gregorianDate: testDate{year: 1877, month: 8, day: 10},
	},
	{
		julianDay:     2460493,
		gregorianDate: testDate{year: 2024, month: 7, day: 1},
	},
}
