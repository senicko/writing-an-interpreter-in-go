package object

import "testing"

func TestStringHashKey(t *testing.T) {
	hello1 := &String{Value: "Hello World"}
	hello2 := &String{Value: "Hello World"}

	diff1 := &String{Value: "Test1"}

	if hello1.HashKey() != hello2.HashKey() {
		t.Errorf("strings with same content have different hash keys")
	}

	if hello1.HashKey() == diff1.HashKey() {
		t.Errorf("strings with different content have same hash keys")
	}
}

func TestBooleanHashKey(t *testing.T) {
	trueBoolean := &Boolean{Value: true}
	falseBoolean := &Boolean{Value: false}

	if trueBoolean.HashKey() == falseBoolean.HashKey() {
		t.Errorf("true has the same hash key sa false")
	}
}

func TestIntegerHashKey(t *testing.T) {
	a1 := &Integer{Value: 2}
	a2 := &Integer{Value: 2}

	b := &Integer{Value: 20}

	if a1.HashKey() != a2.HashKey() {
		t.Errorf("strings with same content have different hash keys")
	}

	if a1.HashKey() == b.HashKey() {
		t.Errorf("strings with different content have same hash keys")
	}
}
