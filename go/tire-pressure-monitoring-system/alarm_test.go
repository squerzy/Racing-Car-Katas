package tpms

import (
	"testing"
)

type alarmTest struct {
	it       string
	expected bool
	actual   func() bool
}

func TestNoAlarmTriggered(t *testing.T) {
	t.Run("Do something", func(t *testing.T) {
		a := NewAlarm(nil)
		a.check()

		if a.isAlarmOn() {
			t.Error("Expected no alarm but was Alarm")
		}
	})

}
