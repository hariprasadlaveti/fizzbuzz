package fizzbuzz

// Fizzbuzz takes a number and returns either string and boolean
// True or empty string an boolean false
//
// Callers are supposed to print the orginal number if false is returned

func Fizzbuzz(num int) (string, bool) {
	if num%15 == 0 {
		return "fizzbuzz", true
	}

	if num%3 == 0 {
		return "fizz", true
	}
	if num%5 == 0 {
		return "buzz", true
	}
	return "", false
}
