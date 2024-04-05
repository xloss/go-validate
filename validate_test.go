package go_validate

import "testing"

func TestToInt(t *testing.T) {
	var (
		valInt    = 5
		valFloat  = 10.0
		valStr    = "15"
		valBadStr = "ab20"
	)

	if toInt(valInt) != 5 {
		t.Errorf("toInt(%v) != 5", valInt)
	}

	if toInt(valFloat) != 10 {
		t.Errorf("toInt(%v) != 10", valFloat)
	}

	if toInt(valStr) != 15 {
		t.Errorf("toInt(%v) != 15", valStr)
	}

	if toInt(valBadStr) != 0 {
		t.Errorf("toInt(%v) != 0", valBadStr)
	}
}

func TestToFloat(t *testing.T) {
	var (
		valInt    = 5
		valFloat  = 10.0
		valStr    = "15.0"
		valBadStr = "ab20.0"
	)

	if toFloat(valInt) != 5.0 {
		t.Errorf("toFloat(%v) != 5.0", valInt)
	}

	if toFloat(valFloat) != 10.0 {
		t.Errorf("toFloat(%v) != 10.0", valFloat)
	}

	if toFloat(valStr) != 15.0 {
		t.Errorf("toFloat(%v) != 15.0", valStr)
	}

	if toFloat(valBadStr) != 0.0 {
		t.Errorf("toFloat(%v) != 0.0", valBadStr)
	}
}
