package main

import (
	"github.com/d2r2/go-i2c"
	"github.com/d2r2/go-logger"
	"github.com/d2r2/go-si7021"
	"log"
)

func main() {
	_ = logger.ChangePackageLogLevel("i2c", logger.InfoLevel)
	_ = logger.ChangePackageLogLevel("si7021", logger.InfoLevel)

	// Create new connection to i2c-bus on 1 line with address 0x40.
	// Use i2cdetect utility to find device address over the i2c-bus
	i2c, err := i2c.NewI2C(0x40, 1)
	if err != nil {
		log.Fatal(err)
	}
	defer i2c.Close()

	sensor := si7021.NewSi7021()

	rh, t, err := sensor.ReadRelativeHumidityAndTemperature(i2c)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Relative humidity and temperature = %v%%, %v*C\n", rh, t)
}
