package calendar

import "errors"

// Date is used to hold dates
type Date struct {
	year  int
	month int
	day   int
}

// SetYear sets the year
func (d *Date) SetYear(year int) error {
	if year < 1 {
		return errors.New("Invalid Year")
	}
	d.year = year
	return nil
}

// SetMonth sets the month
func (d *Date) SetMonth(month int) error {
	if month < 1 || month > 12 {
		return errors.New("Invalid Month")
	}
	d.month = month
	return nil
}

// SetDay sets the day
func (d *Date) SetDay(day int) error {
	if day < 1 || day > 31 {
		return errors.New("Invalid Day")
	}
	d.day = day
	return nil
}

// Year gets the year
func (d *Date) Year() int {
	return d.year
}

// Month gets the month
func (d *Date) Month() int {
	return d.month
}

// Day gets the day
func (d *Date) Day() int {
	return d.day
}
