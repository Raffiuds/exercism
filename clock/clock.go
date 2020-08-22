package clock

import (
	"time"
)

type Clock string

func New(hour, min int) Clock {

	return Clock(time.Date(0, 0, 0, hour, min, 0, 0, time.UTC).Format("15:04"))
}

func (c Clock) Add(min int) Clock {

	t, err := time.Parse("15:04", c.String())
	if err != nil {
		return Clock("")
	}

	c = Clock(t.Add(time.Minute * time.Duration(min)).Format("15:04"))
	return c
}

func (c Clock) Subtract(min int) Clock {

	t, err := time.Parse("15:04", c.String())
	if err != nil {
		return Clock("")
	}

	c = Clock(t.Add(-time.Minute * time.Duration(min)).Format("15:04"))
	return c
}

func (c Clock) String() string {
	return string(c)
}
