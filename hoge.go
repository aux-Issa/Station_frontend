package main

import (
	"fmt"
)

type car struct {
	horn    string
	enabled bool
}

type bike struct {
	name  string
	heavy bool
}

type vehicle interface {
	drive() string
}

func getDrive(v vehicle) string {
	return v.drive()
}

func (c *car) drive() string {
	return c.horn
}

func (b *bike) drive() string {
	return b.name
}

func main() {
	car := car{
		horn:    "honk",
		enabled: true,
	}
	bike := bike{
		name:  "bike",
		heavy: true,
	}
	vehicles := []vehicle{&car, &bike}

	for _, v := range vehicles {
		fmt.Println(getDrive(v))
	}
}