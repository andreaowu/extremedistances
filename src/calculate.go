package main

import (
    "bufio"
    "encoding/csv"
    "io"
    "log"
    "math"
    "os"
    "strconv"
)

type Point struct {
    lat float64 // latitude
    lng float64 // longitude
    id int64    // given id in the file
    dist float64 // distance from office
}

const (
    // According to Wikipedia, the Earth's radius is about 6371km
    EARTH_RADIUS = 6371
)

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

// IsZero returns whether p is a real data point
func IsZero(p Point) bool {
    if p.id == 0 && p.lat == 0 && p.lng == 0 && p.dist == 0 {
        return true
    }
    return false
}

// CheckFill checks whether arr has all data or whether some fields have
// not been written in yet.
func CheckFill(arr [5]Point) bool {
    first := arr[0]
    if IsZero(first) {
        return false
    }
    return true
}


// MoveDown shifts all arr values to the left and puts val at ind in arr.
func MoveDown(val Point, arr [5]Point, ind int) [5]Point {
    for i := 0; i < ind; i++ {
        arr[i] = arr[i + 1]
    }
    arr[ind] = val
    return arr
}

// MoveUp shifts values to the right and inserts val such that arr's order 
// is maintained.
func MoveUp(arr [5]Point, val Point) [5]Point {
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
func AddToArray(arr [5]Point, val Point) [5]Point {
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
func Reverse(arr [5]Point) [5]Point {
    for i, j := 0, len(arr) - 1; i < j; i, j = i + 1, j - 1 {
        arr[i], arr[j] = arr[j], arr[i]
    }
    return arr
}

// PrintToFile prints arr into a file named fileName.
func PrintToFile(fileName string, arr [5]Point) {
    file, err := os.Create(fileName)
    Check(err)
    defer file.Close()
   
    writerClose := csv.NewWriter(file)
    // Add column names as first line to file
    tags := []string{"id", "lat", "lng", "distance in km from office"}
    err = writerClose.Write(tags)

    // Change all numbers into strings and write to file
    for i := 0; i < len(arr); i++ {
        if !IsZero(arr[i]) {
            record := []string{strconv.FormatInt(arr[i].id, 10),
                      strconv.FormatFloat(arr[i].lat, 'f', -1, 64),
                      strconv.FormatFloat(arr[i].lng, 'f', -1, 64), 
                      strconv.FormatFloat(arr[i].dist, 'f', -1, 64),}
            err = writerClose.Write(record)
        }
    }

    defer writerClose.Flush()
}

// Main function of the program reads the file, processes the data, and 
// outputs appropriately.
func main() {
    // Open and read given CSV file
    fileNameWithData := os.Args[1]
    file, err := os.Open(fileNameWithData)
    Check(err)
    defer file.Close()
    r := csv.NewReader(bufio.NewReader(file))

    var closest [5]Point // array to track closest data points to office
    close := math.MaxFloat64 // furthest point in top 5 closest data points
    var furthest [5]Point // array to track furthest data points to office
    far := float64(-1) // closest point in top 5 furthest data points
    office := Point{51.9204549, 4.5099984, -1, 0}
    count := 0
    for {
        record, err := r.Read()
        if err == io.EOF {
            break
        }
        if err != nil {
            log.Fatal(err)
        }
        if count < 1 {
            // skip the first line of the record, which says 'id', 'lat', 'lng'
            count += 1
            continue
        }
        lat, err := strconv.ParseFloat(record[1], 64)
        lng, err := strconv.ParseFloat(record[2], 64)
        id, err := strconv.ParseInt(record[0], 10, 0)
        place := Point{lat, lng, id, 0}
        dist := Distance(place, office)
        place.dist = dist
        if !CheckFill(closest) || dist < close {
            // add it to closest, but add it in the right spot
            closest = AddToArray(closest, place)
            close = Distance(closest[4], office)
        }
        if  !CheckFill(furthest) || dist > far {
            // add it to furthest, but delete first element
            furthest[0] = Point{0, 0, 0, 0}
            furthest = AddToArray(furthest, place)
            far = Distance(furthest[0], office)
        }
        count += 1
    }

    PrintToFile("closest.csv", closest)
    PrintToFile("furthest.csv", Reverse(furthest))
}
