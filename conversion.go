package ptime

// (October 15, 1582)
const gregorianReformJulianDay = 2299160

// isAfterGregorianReform checks if the testDate is after the Gregorian calendar reform (October 15, 1582).
func isAfterGregorianReform(year, month, day int) bool {
	return year > 1582 || (year == 1582 && month > 10) || (year == 1582 && month == 10 && day > 14)
}

// convertGregorianPostReformToJDN calculates the Julian Day Number (JDN) for dates after the Gregorian reform.
// This function is based on the standard algorithm for converting a Gregorian calendar testDate into a Julian Day Number.
// The Gregorian reform was implemented on October 15, 1582, which corrected the drift of the Julian calendar by modifying
// leap year rules and adjusting the calendar by 10 days.
//
// The function uses several components to calculate the JDN:
//   - adjustedYear: The year is adjusted to accommodate the shift caused by the Gregorian reform, adding 4800 and
//     adjusting the month for the year calculation.
//   - leapYearFactor: This factor accounts for the leap years by multiplying the adjusted year by 1461 and dividing by 4.
//   - adjustedMonth: The month is adjusted to align with the calendar calculations, considering the calendar's
//     month structure.
//   - monthFactor: The adjusted month is multiplied by 367 and divided by 12 to align the month correctly.
//   - centuryFactor: A century correction factor is calculated to account for the Gregorian reform's century rules.
//
// Finally, these components are combined with the day of the month (gd) and a constant offset (32075) to compute the JDN.
// https://aa.usno.navy.mil/faq/JD_formula
func convertGregorianPostReformToJDN(year, month, day int) int {
	const (
		// The specific value 1461 is derived from the fact that there are 365.25 days in a year on average,
		// and 1461 represents the number of days in a 4-year cycle (365.25 * 4).
		// This value is chosen to align with the cycles in the Gregorian calendar,
		// especially the leap year cycle, and it's commonly used in algorithms that involve date calculations.
		daysInFourYearCycle     = 1461
		yearOffset              = 4800  // Offset to adjust the year for calculations
		centuryAdjustmentOffset = 4900  // Offset to adjust the century for calculations
		monthCycleFactor        = 367   // Multiplier used in the month cycle calculation
		baseDayAdjustment       = 32075 // Adjustment factor for the base day calculation
	)

	adjustedYear := year + yearOffset + ((month - 14) / 12)
	leapYearFactor := (daysInFourYearCycle * adjustedYear) / 4

	adjustedMonth := month - 2 - 12*((month-14)/12)
	monthFactor := (monthCycleFactor * adjustedMonth) / 12

	centuryFactor := (3 * ((year + centuryAdjustmentOffset + ((month - 14) / 12)) / 100)) / 4

	return leapYearFactor + monthFactor - centuryFactor + day - baseDayAdjustment
}

// convertGregorianPreReformToJDN calculates the Julian Day Number (JDN) for dates before the Gregorian reform.
// Before the Gregorian calendar was introduced in 1582, the Julian calendar was used, which had a simpler rule for leap years
// and no century corrections. This function uses the Julian calendar's rules to calculate the JDN.
//
// The function uses several components to calculate the JDN:
//   - adjustedYear: The year is adjusted to accommodate the Julian calendar's structure, adding 5001 and adjusting the month.
//   - leapYearFactor: This factor accounts for the leap years under the Julian rules, multiplying the adjusted year by 7
//     and dividing by 4.
//   - monthFactor: The month is multiplied by 275 and divided by 9 to align the month correctly.
//
// These components are combined with the day of the month (gd) and a constant offset (1729777) to compute the JDN for dates
// before the Gregorian reform.
// https://aa.usno.navy.mil/faq/JD_formula
func convertGregorianPreReformToJDN(year, month, day int) int {
	adjustedYear := year + 5001 + (month-9)/7
	leapYearFactor := (7 * adjustedYear) / 4

	monthFactor := (275 * month) / 9

	return 367*year - leapYearFactor + monthFactor + day + 1729777
}

// convertJDNToGregorianPostReform converts a Julian Day Number (JDN) to the corresponding
// Gregorian testDate for dates after the Gregorian calendar reform (after October 15, 1582).
// https://aa.usno.navy.mil/faq/JD_formula
func convertJDNToGregorianPostReform(jdn int) (year, month, day int) {
	const (
		daysInFourYearCycle = 1461
		// The specific value 2447 is chosen based on the characteristics of the Gregorian calendar and its various cycles.
		// It plays a role in the algorithm to determine the month and day components of the Gregorian date during the conversion process from Julian Day.
		daysInMonthMultiplier = 2447
		// Offset used to adjust Julian Day
		julianDayOffset = 68569
		// The specific value 1461001 is derived from the fact that there are 365.25 days in a year on average,
		// and 1461001 represents the number of days in a 4000-year cycle (365.25 * 4000).
		// This value is chosen to align with the cycles in the Gregorian calendar,
		// facilitating the conversion between Julian Day and Gregorian Date.
		julianDay4000YearCycleDayOffset = 1461001 // 365.25 * 4000
		// The specific value 146097 is derived from the fact that there are 365.25 days in a year on average,
		// and 146097 represents the number of days in a 400-year cycle (365.25 * 400).
		// This value is commonly used in algorithms that involve date calculations, especially when dealing with leap years.
		julianDayOf400Years = 146097 // 365.25 * 400
	)

	offsetJDN := jdn + julianDayOffset

	// Calculate century
	century := 4 * offsetJDN / julianDayOf400Years
	offsetJDN = offsetJDN - (julianDayOf400Years*century+3)/4

	// Calculate year
	yearBase := 4000 * (offsetJDN + 1) / julianDay4000YearCycleDayOffset
	offsetJDN = offsetJDN - daysInFourYearCycle*yearBase/4 + 31

	// Calculate month and day
	monthFactor := 80 * offsetJDN / daysInMonthMultiplier
	day = offsetJDN - daysInMonthMultiplier*monthFactor/80
	offsetJDN = monthFactor / 11
	month = monthFactor + 2 - 12*offsetJDN
	year = 100*(century-49) + yearBase + offsetJDN

	return year, month, day
}

// convertJDNToGregorianPreReform converts a Julian Day Number (JDN) to the corresponding
// Gregorian testDate for dates before the Gregorian calendar reform (before October 15, 1582).
func convertJDNToGregorianPreReform(jdn int) (year, month, day int) {
	const (
		daysInFourYearCycle   = 1461
		daysInMonthMultiplier = 2447 // Multiplier used for month calculation
		julianDayOffset       = 1402 // Offset used to adjust Julian Day for pre-Gregorian dates
	)

	offsetJDN := jdn + julianDayOffset

	// Calculate year
	quadrennialCycle := (offsetJDN - 1) / daysInFourYearCycle
	remainingDays := offsetJDN - daysInFourYearCycle*quadrennialCycle
	yearAdjustment := (remainingDays-1)/365 - remainingDays/daysInFourYearCycle
	dayOfYear := remainingDays - 365*yearAdjustment + 30

	// Calculate month and day
	monthFactor := 80 * dayOfYear / daysInMonthMultiplier
	day = dayOfYear - daysInMonthMultiplier*monthFactor/80
	yearFraction := monthFactor / 11
	month = monthFactor + 2 - 12*yearFraction
	year = 4*quadrennialCycle + yearAdjustment + yearFraction - 4716

	return year, month, day
}

// convertJDNToShamsi converts a Julian Day Number (JDN) to the Shamsi (Solar Hijri) calendar testDate.
// The conversion is based on the offset between the Julian calendar and the Shamsi calendar.
// The calculation is performed as follows:
// - The JDN is adjusted by subtracting a constant offset to align it with the Shamsi calendar.
// - The resulting value is used to calculate the year by accounting for cycles of 33 years and leap years.
// - The month and day are then calculated based on the remaining days within the year.
//
// Parameters:
// - jdn: The Julian Day Number to be converted.
//
// Returns:
// - year: The calculated year in the Shamsi calendar.
// - month: The calculated month in the Shamsi calendar.
// - day: The calculated day in the Shamsi calendar.
func convertJDNToShamsi(jdn int) (year, month, day int) {
	const (
		julianDayToShamsiOffset = 1365393
		cyclesOf33YearsCount    = 12053 // 33 * 364.24
		daysInFourYearCycle     = 1461
		middleDayInYear         = 186 // 6 * 31
	)

	// Align the Julian Day Number with the Shamsi calendar
	daysSinceStartOfShamsi := jdn - julianDayToShamsiOffset

	// Calculate the Shamsi year
	cyclesOf33Years := daysSinceStartOfShamsi / cyclesOf33YearsCount
	year = -1595 + 33*cyclesOf33Years
	remainingDays := daysSinceStartOfShamsi % cyclesOf33YearsCount

	cyclesOf4Years := remainingDays / daysInFourYearCycle
	year += 4 * cyclesOf4Years
	remainingDays %= daysInFourYearCycle

	// Adjust for remaining days within the current cycle
	if remainingDays > 365 {
		year += (remainingDays - 1) / 365
		remainingDays = (remainingDays - 1) % 365
	}

	// Determine the Shamsi month and day
	if remainingDays < middleDayInYear {
		month = 1 + remainingDays/31
		day = 1 + remainingDays%31
	} else {
		month = 7 + (remainingDays-middleDayInYear)/30
		day = 1 + (remainingDays-middleDayInYear)%30
	}

	return year, month, day
}

// convertShamsiToJDN converts a Shamsi (Solar Hijri) calendar testDate to the corresponding Julian Day Number (JDN).
// The calculation takes into account the specific offset and adjustments needed for leap years in the Shamsi calendar.
func convertShamsiToJDN(year, month, day int) int {
	const (
		shamsiToJulianOffset = 1365392
		leapYearCycle        = 33
		leapYearContribution = 8
		middleDayInYear      = 186 // 6 * 31
		daysInFirstSixMonths = 31  // Months 1-6: each month has 31 days
		daysInNextSixMonths  = 30  // Months 7-12: each month has 30 days
	)

	// Adjust the Shamsi year for the calculation
	adjustedShamsiYear := year + 1595

	// Calculate the number of leap years that have occurred up to the given year
	leapYearContributionCount := (adjustedShamsiYear/leapYearCycle)*leapYearContribution +
		((adjustedShamsiYear%leapYearCycle + 3) / 4)

	// Determine the day of the year within the Shamsi calendar
	var dayOfYear int
	if month < 7 {
		dayOfYear = (month - 1) * daysInFirstSixMonths
	} else {
		dayOfYear = (month-7)*daysInNextSixMonths + middleDayInYear
	}

	// Calculate the Julian Day Number (JDN)
	jdn := shamsiToJulianOffset + 365*adjustedShamsiYear +
		leapYearContributionCount + dayOfYear + day

	return jdn
}
