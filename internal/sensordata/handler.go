package sensordata

import (
	"fmt"
	"log"
	"net/http"
	"sync/atomic"
)

var errorCnt int32

func (sd *SensorData) Export(w http.ResponseWriter, _ *http.Request) {
	rh, t, err := sd.ReadRelativeHumidityAndTemperature()
	if err != nil {
		log.Print(err)
		_, _ = fmt.Fprintf(w, "si7021_error %d", atomic.AddInt32(&errorCnt, 1))
		return
	}
	_, _ = fmt.Fprintf(w, "si7021_temperature %f\n", t)
	_, _ = fmt.Fprintf(w, "si7021_relative_humidity %f\n", rh)
}
