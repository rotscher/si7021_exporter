package main

import (
	"github.com/d2r2/go-logger"
	"github.com/rotscher/si7021_exporter/internal/sensordata"
	"time"
)

func main() {
	_ = logger.ChangePackageLogLevel("i2c", logger.InfoLevel)
	_ = logger.ChangePackageLogLevel("si7021", logger.InfoLevel)

	sd := sensordata.New()
	defer func(sd *sensordata.SensorData) {
		err := sd.Close()
		if err != nil {

		}
	}(&sd)

	for {
		sd.ReadRelativeHumidityAndTemperature()
		time.Sleep(10 * time.Second)
	}

}
