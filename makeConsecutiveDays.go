package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"math"
	"strconv"
	"time"
)

func diff(a, b time.Time) (year, month, day, hour, min, sec int) {
	if a.Location() != b.Location() {
		b = b.In(a.Location())
	}
	if a.After(b) {
		a, b = b, a
	}
	y1, M1, d1 := a.Date()
	y2, M2, d2 := b.Date()

	h1, m1, s1 := a.Clock()
	h2, m2, s2 := b.Clock()

	year = int(y2 - y1)
	month = int(M2 - M1)
	day = int(d2 - d1)
	hour = int(h2 - h1)
	min = int(m2 - m1)
	sec = int(s2 - s1)

	// Normalize negative values
	if sec < 0 {
		sec += 60
		min--
	}
	if min < 0 {
		min += 60
		hour--
	}
	if hour < 0 {
		hour += 24
		day--
	}
	if day < 0 {
		// days in month:
		t := time.Date(y1, M1, 32, 0, 0, 0, 0, time.UTC)
		day += 32 - t.Day()
		month--
	}
	if month < 0 {
		month += 12
		year--
	}

	return
}

func Date(year, month, day int) time.Time {
    fmt.Printf("Inside Date year: %d month: %d day:%d\n", year, month, day)
    return time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
}

func main() {
    csvfile, err := os.Open("ConscotizaWITHOUTHEADER.csv")
   
    if err != nil {
	    log.Fatalln("could not open the file", err)
    }

    r := csv.NewReader(csvfile)
    indexRow := 0
    nextPreviousDate := 0.0
    var cDate time.Time
    var pDate time.Time
    var year int
    var month int
    var day int

    for {
        record, err := r.Read()
	if err == io.EOF {
	    break
	}


        currentDate, err := strconv.ParseFloat(record[0], 64)

	year, err = strconv.Atoi(record[2]) 
	day, err= strconv.Atoi(record[4])
	month, err= strconv.Atoi(record[3])
	cDate = Date(day, month, year)

        fmt.Printf("Current Row %s \n", record[0])
        fmt.Printf("COMPARING Previous vs Current | Previous: %2f Current %2f\n", nextPreviousDate, currentDate)
        fmt.Printf("COMPARING Previous vs Current Date Format | Previous: %s Current %s\n", pDate.Format("2006 01 02"), cDate.Format("2006 01 02"))
        fmt.Printf("COMPARING Previous vs Current Date Format | Previous: %s Current %s\n", pDate.String(), cDate.String())
        fmt.Printf("COMPARING Date Format | Day: %d Month %d Year %d\n", year, month, day)

	if (indexRow!=0) {
             var differencialDate = (nextPreviousDate - currentDate) - 1.0 
	     year, month, day, hour, min, sec := diff(pDate, cDate)
             fmt.Printf("COMPARING Cdate sub pdate: %d Years %d Months %d Days %d hours %d min %d secs\n", year, month, day, hour, min, sec)

	     if ((day-1)>1) {
        	fmt.Printf("I should add new rows new VERSION: %d \n", (day-1)) 
	     }
             if (math.Abs(differencialDate)>1) {
                if (math.Abs(differencialDate) > 10) {
			differencialDate = math.Abs(differencialDate) - 69
        	        fmt.Printf("I should add new rows: %2f \n", math.Abs(differencialDate)) 
		}
             }
 	}

	if err != nil {
            log.Fatal(err)
	}
	nextPreviousDate = currentDate
	pDate = Date(day, month, year)
        fmt.Printf("Count: %d \n", indexRow)
	indexRow ++
	fmt.Printf("NEW nextPreviousDate: %2f \n", nextPreviousDate)
	fmt.Printf("=====================================================\n")
    }
    fmt.Printf("FINAL VALUE nextPreviousDate %f \n", nextPreviousDate)
}
