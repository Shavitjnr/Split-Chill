package core

import "fmt"


type CoordinateDisplayType byte


const (
	COORDINATE_DISPLAY_TYPE_DEFAULT                                    CoordinateDisplayType = 0
	COORDINATE_DISPLAY_TYPE_LATITUDE_LONGITUDE_DECIMAL_DEGREES         CoordinateDisplayType = 1
	COORDINATE_DISPLAY_TYPE_LONGITUDE_LATITUDE_DECIMAL_DEGREES         CoordinateDisplayType = 2
	COORDINATE_DISPLAY_TYPE_LATITUDE_LONGITUDE_DECIMAL_MINUTES         CoordinateDisplayType = 3
	COORDINATE_DISPLAY_TYPE_LONGITUDE_LATITUDE_DECIMAL_MINUTES         CoordinateDisplayType = 4
	COORDINATE_DISPLAY_TYPE_LATITUDE_LONGITUDE_DEGREES_MINUTES_SECONDS CoordinateDisplayType = 5
	COORDINATE_DISPLAY_TYPE_LONGITUDE_LATITUDE_DEGREES_MINUTES_SECONDS CoordinateDisplayType = 6
	COORDINATE_DISPLAY_TYPE_INVALID                                    CoordinateDisplayType = 255
)


func (d CoordinateDisplayType) String() string {
	switch d {
	case COORDINATE_DISPLAY_TYPE_DEFAULT:
		return "Default"
	case COORDINATE_DISPLAY_TYPE_LATITUDE_LONGITUDE_DECIMAL_DEGREES:
		return "Latitude Longitude (Decimal Degrees)"
	case COORDINATE_DISPLAY_TYPE_LONGITUDE_LATITUDE_DECIMAL_DEGREES:
		return "Longitude Latitude (Decimal Degrees)"
	case COORDINATE_DISPLAY_TYPE_LATITUDE_LONGITUDE_DECIMAL_MINUTES:
		return "Latitude Longitude (Decimal Minutes)"
	case COORDINATE_DISPLAY_TYPE_LONGITUDE_LATITUDE_DECIMAL_MINUTES:
		return "Longitude Latitude (Decimal Minutes)"
	case COORDINATE_DISPLAY_TYPE_LATITUDE_LONGITUDE_DEGREES_MINUTES_SECONDS:
		return "Latitude Longitude (Degrees Minutes Seconds)"
	case COORDINATE_DISPLAY_TYPE_LONGITUDE_LATITUDE_DEGREES_MINUTES_SECONDS:
		return "Longitude Latitude (Degrees Minutes Seconds)"
	case COORDINATE_DISPLAY_TYPE_INVALID:
		return "Invalid"
	default:
		return fmt.Sprintf("Invalid(%d)", int(d))
	}
}
