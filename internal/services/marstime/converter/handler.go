package converter

import (
	"fmt"
	"gopkg.in/matryer/respond.v1"
	"math"
	"net/http"
	"time"
)

type EarthTime struct {
	UTC string
}

type ErrorMessage struct {
	Error string `json:"error"`
}

type Result struct {
	MSD float64 `json:"MSD"`
	MTC string  `json:"MTC"`
}

const diff = 1.027491252
const leapSeconds = 0.00080074074
const epochDays = 2451545.0
const MSDPositiveFixer = 44796.0
const deltaDays = 4.5
const Mars24Adjustment = 0.0009626

func Handler(w http.ResponseWriter, r *http.Request) {

	query := r.URL.Query()
	timeUTC, ok := query["UTC"]
	if !ok || len(timeUTC) == 0 {
		message := ErrorMessage{fmt.Sprintf("You should provide UTC time in format %s", time.RFC3339)}
		respond.With(w, r, http.StatusBadRequest, message)
		return
	}

	EarthTime, err := time.Parse(time.RFC3339, timeUTC[0])

	if err != nil {
		message := ErrorMessage{err.Error()}
		respond.With(w, r, http.StatusBadRequest, message)
		return
	}

	JulianTimeUT := julian(EarthTime)
	JulianTimeTT := JulianTimeUT + leapSeconds
	DeltaTime := JulianTimeTT - epochDays
	MarsTime := ((DeltaTime - deltaDays) / diff) + MSDPositiveFixer - Mars24Adjustment

	MSDh := math.Mod(24*MarsTime, 24)
	MSDm := math.Mod(60*MSDh, 60)
	MSDs := math.Round(math.Mod(60*MSDm, 60))
	MSDFormatted := fmt.Sprintf("%d:%d:%d", int64(MSDh), int64(MSDm), int64(MSDs))

	result := Result{
		MSD: MarsTime,
		MTC: MSDFormatted,
	}

	respond.With(w, r, http.StatusOK, result)
}

func julian(t time.Time) float64 {
	const julian = 2453738.4195
	unix := time.Unix(1136239445, 0)
	const oneDay = float64(86400. * time.Second)
	return julian + float64(t.Sub(unix))/oneDay
}
