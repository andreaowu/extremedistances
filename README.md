#Distance Calculator

Written in Go.

Functionality:
Given a location with latitude and longitude both in degrees and a CSV file with record ID, latitude in degrees, and longitude in degrees, output two CSV files. Both files will contain 5 records, each with the record ID, latitude and longitude of the original point in degrees, as well as the distance in kilometers from the given location. One file will be of the top 5 closest points to the given location and the other will be the top 5 furthest points from the given location

Deployment
==========
## Build 
This command will build the executable and the executable file will be in the directory in which you invoke this command:
```
go build src/calculate.go
```

## Run
Executing the following will generate the two files mentioned above, and these two files will be in the directory where you execute this command. Make sure you give it an argument with the directory of the file containing the data points.
```
./calculate <file>
```

As a side note, the location from which the closest and furthest points are determined is hard-coded into the file.
