package tpms

import (
	"math/rand"
	"time"
)

type sensor struct {
	offset         func() int
	samplePressure func() int
}

func (s sensor) popNextPressurePsiValue() int {
	return s.offset() + s.samplePressure()
}

func NewSensor() Sensor {
	return &sensor{
		offset: func() int {
			return 0
		},
		samplePressure: func() int {
			s := rand.NewSource(time.Now().UnixNano())
			r := rand.New(s)
			pressureTelemetryValue := r.Intn(20) + 10
			// fmt.Printf("New Sensor: %d", pressureTelemetryValue)
			return pressureTelemetryValue
		},
	}
}
