package fizzbuzz_test

import (
	"fmt"
	"testing"

	"github.com/hariprasadlaveti/fizzbuzz"
)

func TestFizz(t *testing.T) {
	_, ok := fizzbuzz.Fizzbuzz(1)

	if ok {
		t.Logf("the value %v should not have retruned true", 1)
		t.Fail()
	}

	result, ok := fizzbuzz.Fizzbuzz(3)
	if !ok {
		t.Logf("the value %v should have returned true", 3)
		t.Fail()
	}

	result1, ok1 := fizzbuzz.Fizzbuzz(5)
	if !ok1 {
		t.Logf("the value %v should have returned true", 5)
		t.Fail()
	}

	result2, ok2 := fizzbuzz.Fizzbuzz(45)
	if !ok2 {
		t.Logf("the value %v should have returned true", 45)
		t.Fail()
	}

	if result != "fizz" || result1 != "buzz" || result2 != "fizzbuzz" {
		t.Log("Result should be fizz")
		t.Fail()
	}
}

func ExampleFizzBuzz() {
	result, _ := fizzbuzz.Fizzbuzz(3)
	fmt.Println(result)
	// Output:fizz
}
