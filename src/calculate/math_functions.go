package main

import (
    "math"
)

// According to Wikipedia, the Earth's radius is about 6371km
const EARTH_RADIUS float64 = 6371

// ToRadians takes a degrees and returns a in in radians.
func ToRadians(a float64) float64 {
    return a * math.Pi / 180
}

// Distance takes two points, a and b, and returns the great circle distance
// the two points.
func Distance(a Point, b Point) float64 {
    latDiff := ToRadians(b.lat - a.lat)
    lngDiff := ToRadians(b.lng - a.lng)
    aLatRad := ToRadians(a.lat)
    bLatRad := ToRadians(b.lat)

    x := math.Sin(latDiff / 2) * math.Sin(latDiff / 2)
    y := math.Sin(lngDiff / 2) * math.Sin(lngDiff / 2) * math.Cos(aLatRad) * math.Cos(bLatRad)
    z := x + y
    c := 2 * math.Atan2(math.Sqrt(z), math.Sqrt(1 - z))

    return EARTH_RADIUS * c
}
