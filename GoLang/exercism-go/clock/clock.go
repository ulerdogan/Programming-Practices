package clock

import "fmt"

type Clock struct {
	hour   int
	minute int
}

func (c Clock) FixTime() Clock {

	for c.minute > 59 {
		c.minute -= 60
		c.hour += 1
	}

	for c.minute < 0 {
		c.minute += 60
		c.hour -= 1
	}

	for c.hour > 23 {
		c.hour -= 24
	}

	for c.hour < 0 {
		c.hour += 24
	}

	return c
}

func New(hour, minute int) Clock {
	var c Clock
	c.hour = hour
	c.minute = minute
	return c.FixTime()
}

func (c Clock) Add(minutes int) Clock {
	c.minute += minutes
	return New(c.hour, c.minute)
}

func (c Clock) Subtract(minutes int) Clock {
	c.minute -= minutes
	return New(c.hour, c.minute)
}

func (c Clock) String() string {
	return fmt.Sprintf("%.2d:%.2d", c.hour, c.minute)
}
