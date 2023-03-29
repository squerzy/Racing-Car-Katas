package tpms

type Alarm interface {
	check()
	isAlarmOn() bool
}

type Sensor interface {
	popNextPressurePsiValue() int
}

type alarm struct {
	lowPressureThreshold  int
	highPressureThreshold int
	sensor                Sensor
	alarmOn               bool
}

func (a *alarm) check() {
	p := a.sensor.popNextPressurePsiValue()

	if p < a.lowPressureThreshold || a.highPressureThreshold < p {
		a.alarmOn = true
	}
}

func (a *alarm) isAlarmOn() bool {
	return a.alarmOn
}

func NewAlarm() Alarm {
	return &alarm{
		lowPressureThreshold:  17,
		highPressureThreshold: 21,
		alarmOn:               false,
		sensor:                NewSensor(), // SOLID - Open Closed Violation
	}

}
