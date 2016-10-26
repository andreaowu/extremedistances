package main

import (
    "bufio"
    "encoding/csv"
    "io"
    "log"
    "os"
    "strconv"
)

// OpenFile opens and reads given CSV file.
func OpenFile(fileName string) [][] string {

    // Open the file
    file, err := os.Open(fileName)
    Check(err)
    defer file.Close()
    r := csv.NewReader(bufio.NewReader(file))

    var records [][]string
    count := 0

    // Loop through all records in the file
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

        records = append(records, record)
    }

    return records

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
