package optional

import "testing"

func TestEmpty(t *testing.T) {
	o := Empty()
	if o.value != nil {
		t.Errorf("Empty(): Expected true but got false")
	}
	if o.IsPresent() {
		t.Errorf("IsPresent(): Expected false but got true")
	}
}

func TestOf(t *testing.T) {
	var strPtr *string

	strPtr = nil
	opt := Of(strPtr)
	if opt.IsPresent() {
		t.Errorf("IsPresent(): Expected false (not present) but got true")
	}

	strPtr = new("hello world")
	opt = Of(strPtr)
	if !opt.IsPresent() {
		t.Errorf("IsPresent(): Expected true (present) but got false")
	}
	var value *string = opt.Get()
	if len(*value) != len(*strPtr) {

	}
}
