package main

import (
    "os"
)

// Constant for office location
var office Point = Point{51.9204549, 4.5099984, -1, 0}
var top int = 5

// Main will run the entire program
func main() {
	// Get the file with the data from the command line
    fileName := os.Args[1]
    records := OpenFile(fileName)

    // Calculate the closest and furthest data points
    closest := CalculateClosest(records, office, top)
    furthest := CalculateFurthest(records, office, top)

    // Print results to file
    PrintToFile("closest.csv", closest)
    PrintToFile("furthest.csv", Reverse(furthest))
}
