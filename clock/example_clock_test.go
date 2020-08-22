package clock

import (
	"fmt"
	"log"
	"testing"
)

func TestExampleClock_new(t *testing.T) {
	// a new clock
	clock1 := New(10, 30)
	log.Println(clock1.String())

	if got := clock1.String(); got != "10:30" {
		t.Fatalf("New(%d, %d) = %q, want %q", 10, 30, got, "10:30")
	}
	// Output: 10:30
}

func TestExampleClock_Add(t *testing.T) {
	// create a clock
	clock := New(10, 30)

	// add 30 minutes to it
	clock = clock.Add(30)
	if got := clock.String(); got != "11:00" {
		t.Fatalf("New(%d, %d) = %q, want %q", 10, 30, got, "11:00")
	}

	// Output: 11:00
}

func ExampleClock_Subtract() {
	// create a clock
	clock := New(10, 30)

	// subtract an hour and a half from it
	clock = clock.Subtract(90)
	fmt.Println(clock.String())

	// Output: 09:00
}

func ExampleClock_compare() {
	// a new clock
	clock1 := New(10, 30)

	// a second clock, same as the first
	clock2 := New(10, 30)

	// are the clocks equal?
	fmt.Println(clock2 == clock1)

	// change the second clock
	clock2 = clock2.Add(30)

	// are the clocks equal now?
	fmt.Println(clock2 == clock1)

	// Output:
	// true
	// false
}
