# Weather Station CLI (Go)

This project is a command-line Weather Station program written in Go.  
It was developed as part of a programming task at **Hive Helsinki**.

The application stores and updates weather measurements using sensor IDs and allows the user to update, view, and clear the data interactively through standard input.

## Overview

The program reads measurement inputs in `id,value` format, maps them to specific weather fields, and stores them in a structured data model. Values are stored as pointers so the program can distinguish between missing data (`NULL`) and actual numeric values (including `0`).


## Features

- Interactive command-line interface
- Update weather measurements by sensor ID
- Support for NULL values
- Display all stored weather data
- Clear all stored measurements
- Pointer-based storage to differentiate missing vs zero values


## Data Model

Weather data is stored in a `WeatherData` struct with pointer fields:

- Air temperature
- Air pressure
- Precipitation
- Wind speed
- Wind direction
- Humidity
- Dew point
- Soil moisture
- Cloud cover

Using `*float64` allows the program to represent missing values as `nil`.


## Supported Sensor IDs

| ID | Field |
|-----|----------------|
| 1   | airTemp |
| 2   | airPressure |
| 7   | precipitation |
| 11  | windSpeed |
| 12  | windDirection |
| 13  | humidity |
| 14  | dewPoint |
| 15  | soilMoisture |
| 22  | cloudCover |


## How to Run

Make sure Go is installed, then run:

```bash
go run .
```

or:

```bash
go run main.go
```

The program will start and wait for user input.

---

## Input Format

Measurements must be entered as:

```
id,value
```

### Examples

```
1,21.5
13,55
7,NULL
```

Rules:

- The first value is the sensor ID
- The second value is a number or `NULL`
- `NULL` clears that measurement


## Commands

### Show stored data

```
get
```

Prints all weather fields. Missing values are shown as `NULL`.

### Clear all data

```
clear
```

Resets the entire weather struct to empty values.

### Exit program

```
exit
```

Terminates the application.


## Implementation Notes

- Input is read using `bufio.Reader`
- Values are parsed with `strconv`
- Struct fields are updated through an ID-based switch
- Reset is implemented by assigning a zero-value struct:

```go
*s = WeatherData{}
```

- Printing checks for `nil` pointers and outputs `NULL` when no value is stored



## Possible Extensions

- Input validation and error handling
- Timestamped measurements
- Multiple station support
