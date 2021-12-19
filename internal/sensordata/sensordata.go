package sensordata

import (
	"github.com/d2r2/go-i2c"
	"github.com/d2r2/go-si7021"
	"log"
)

type SensorData struct {
	sensor *si7021.Si7021
	i2cBus *i2c.I2C
}

func New() SensorData {

	sensorData := SensorData{}
	var err error
	// Create new connection to i2c-bus on 1 line with address 0x40.
	// Use i2cdetect utility to find device address over the i2c-bus
	sensorData.i2cBus, err = i2c.NewI2C(0x40, 1)
	if err != nil {
		log.Fatal(err)
	}

	sensorData.sensor = si7021.NewSi7021()
	return sensorData
}

func (sd *SensorData) ReadRelativeHumidityAndTemperature() {
	rh, t, err := sd.sensor.ReadRelativeHumidityAndTemperature(sd.i2cBus)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Relative humidity and temperature = %v%%, %v*C\n", rh, t)

}

func (sd *SensorData) Close() error {
	return sd.i2cBus.Close()
}
