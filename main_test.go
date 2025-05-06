package main

import "testing"

func Test_fizzbuzz_3(t *testing.T) {
	got := fizzbuzz(3)
	want := "Fizz"
	if got != want {
		t.Errorf("fizzbuzz(3) \n got: %v \n want: \n%v", got, want)
	}
}

func Test_fizzbuzz_5(t *testing.T) {
	got := fizzbuzz(5)
	want := "Buzz"
	if got != want {
		t.Errorf("fizzbuzz(5) \n got: %v \n want: \n%v", got, want)
	}
}

func Test_fizzbuzz_15(t *testing.T) {
	got := fizzbuzz(15)
	want := "FizzBuzz"
	if got != want {
		t.Errorf("fizzbuzz(15) \n got: %v \n want: \n%v", got, want)
	}
}
