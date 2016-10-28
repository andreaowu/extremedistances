package main

import (
    "math"
    "strconv"
)

func InitializeVars(record []string, office Point) (Point, float64) {
    // Get each field in the record
    lat, err := strconv.ParseFloat(record[1], 64)
    Check(err)
    lng, err := strconv.ParseFloat(record[2], 64)
    Check(err)
    id, err := strconv.ParseInt(record[0], 10, 0)
    Check(err)
    place := Point{lat, lng, id, 0}
    dist := Distance(place, office)
    place.dist = dist

    return place, dist
}


// IsZero returns whether p is a real data point
func IsZero(p Point) bool {
    if p.id == 0 && p.lat == 0 && p.lng == 0 && p.dist == 0 {
        return true
    }
    return false
}

// CheckFill checks whether arr has all data or whether some fields have
// not been written in yet.
func CheckFill(arr []Point) bool {
    first := arr[0]
    if IsZero(first) {
        return false
    }
    return true
}


// MoveDown shifts all arr values to the left and puts val at ind in arr.
func MoveDown(val Point, arr []Point, ind int) []Point {
    for i := 0; i < ind; i++ {
        arr[i] = arr[i + 1]
    }
    arr[ind] = val
    return arr
}

// MoveUp shifts values to the right and inserts val such that arr's order 
// is maintained.
func MoveUp(arr []Point, val Point) []Point {
    for i := 0; i < len(arr); i++ {
        if IsZero(arr[i]) {
            // Reached end of array so need to insert value
            if i == len(arr) - 1 || (i < len(arr) - 1 && arr[i + 1].dist > val.dist) {
                arr[i] = val
                break
            }
            continue
        }
        if val.dist < arr[i].dist {
            arr[i - 1] = val
            break
        } else {
            arr[i - 1] = arr[i]
            if i == len(arr) - 1 {
                arr[i] = val
            }
        }
    }
    return arr
}

// AddToArray adds val to arr. 
func AddToArray(arr []Point, val Point) []Point {
    if !CheckFill(arr) {
        arr = MoveUp(arr, val)
        return arr
    }
    for i := len(arr) - 2; i > -1; i-- {
        if val.dist < arr[i].dist {
            arr[i + 1] = arr[i]
            if i == 0 {
                arr[i] = val
            }
        } else {
            arr[i + 1] = val
            break
        }
    }
    return arr
}

// Check checks for an error and panics if there's an error.
func Check(e error) {
    if e != nil {
        panic(e)
    }
}

// Reverse returns the reverse of arr.
func Reverse(arr []Point) []Point {
    for i, j := 0, len(arr) - 1; i < j; i, j = i + 1, j - 1 {
        arr[i], arr[j] = arr[j], arr[i]
    }
    return arr
}

func CalculateFurthest(data [][]string, office Point, top int) []Point {

    var furthest []Point
    furthest = make([]Point, top) // array to track furthest data points to office
    far := float64(-1) // closest point in top furthest data points
    count := 0

    for _, record := range data {
        place, dist := InitializeVars(record, office)

        if  !CheckFill(furthest) || dist > far {
            // add it to furthest, but delete first element
            furthest[0] = Point{0, 0, 0, 0}
            furthest = AddToArray(furthest, place)
            far = Distance(furthest[0], office)
        }
        count += 1
    }

    return Reverse(furthest)
}

// CalculateClosest function does all the calculations for getting the
// closest data points.
func CalculateClosest(data [][]string, office Point, top int) []Point {

    var closest []Point 
    closest = make([]Point, top) // array to track closest data points to office
    close := math.MaxFloat64 // furthest point in top closest data points
    count := 0

    for _, record := range data {
        place, dist := InitializeVars(record, office)

        if !CheckFill(closest) || dist < close {
            // add it to closest, but add it in the right spot
            closest = AddToArray(closest, place)
            close = Distance(closest[top - 1], office)
        }
        count += 1
    }

    return closest
}
