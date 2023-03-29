package tpms

import (
	"testing"
)

type alarmTest struct {
	it       string
	expected bool
	actual   func() bool
}

type TestSensor struct {
	nextVal int
}

func (t TestSensor) popNextPressurePsiValue() int {
	return 20
}

func TestNoAlarmMidrange(t *testing.T) {
	t.Run("Do something", func(t *testing.T) {

		s := TestSensor{
			nextVal: 20,
		}
		a := NewAlarm(s)
		a.check()

		if a.isAlarmOn() {
			t.Error("Expected no alarm but was Alarm")
		}
	})
}

func TestAlarmBelowRange(t *testing.T) {
	t.Run("Do something", func(t *testing.T) {

		s := TestSensor{
			nextVal: 10,
		}
		a := NewAlarm(s)
		a.check()

		if !a.isAlarmOn() {
			t.Error("Expected alarm did not get one")
		}
	})
}
