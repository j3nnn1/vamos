package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"time"
)

func diff(a, b time.Time) (year, month, day, hour, min, sec int) {
/*	https://www.socketloop.com/tutorials/golang-minus-time-with-time-add-or-time-adddate-functions-to-calculate-past-date */
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
/*    fmt.Printf("Inside Date year: %d month: %d day:%d\n", year, month, day) */
    return time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
}

func isFirstRow(index int) bool {
    return index!=0
}

func initializeFromInput(row []string) time.Time {
    year, err := strconv.Atoi(row[2]) 
    day, err := strconv.Atoi(row[4])
    month, err:= strconv.Atoi(row[3])
    if err != nil {
        log.Fatalln("could not cast vars", err)
    }
    return Date(day, month, year)
}
func thereIsAGapBetweenDates(previosDate time.Time, cDate time.Time) bool {
    year, month, day, hour, min, sec := diff(previosDate, cDate)
    fmt.Printf("COMPARING Cdate sub pdate: %d Years %d Months %d Days %d hours %d min %d secs\n", year, month, day, hour, min, sec)
    /* get Yesterday time.Now().AddDate(0, 0, -1) */
    return (day-1)>1
}
func getDaysOfSubtractionBetween(a time.Time, b time.Time) int {
    if a.After(b) {
        a, b = b, a
    }
    d1 := a.Day()
    d2 := b.Day()
    return int(d2 - d1)
}

func main() {
    var value []string
    var cDate time.Time
    var previosDate time.Time

    csvfileInput, err := os.Open("ConscotizaWITHOUTHEADER.csv")

    csvfileOutput, err := os.Create("filledDates.csv")
    writer := csv.NewWriter(csvfileOutput)
    defer csvfileOutput.Close() 
    defer writer.Flush() 
    
    if err != nil {
        log.Fatalln("could not open the file", err)
    }

    r := csv.NewReader(csvfileInput)
    indexRow := 0

    for {
        record, err := r.Read()
	if err == io.EOF {
	    break
	}
	cDate = initializeFromInput(record)
        fmt.Printf("Current Row %s \n", record[0])
        fmt.Printf("COMPARING Previous vs Current Date Format | Previous: %s Current %s\n", previosDate.String(), cDate.String())

	if (isFirstRow(indexRow)) {
	     if (thereIsAGapBetweenDates(previosDate, cDate)) {
		day := getDaysOfSubtractionBetween(previosDate, cDate)
                fmt.Printf("I should add new rows new VERSION: %d \n", (day-1)) 

		var newTmp [10] time.Time
		var prevTmp time.Time
		prevTmp = previosDate

		for i := 0; i < (day -1); i++ {
		     newTmp[i] = prevTmp.AddDate(0, 0, 1)
		     prevTmp = newTmp[i]
	             value = []string{newTmp[i].String(),record[0], record[5], record[6]}
		     writer.Write(value)
		}
	     }
 	}

	value := []string{cDate.String(), record[0], record[5], record[6]}
	writer.Write(value)
	fmt.Printf("=========SAVING IN FILE =============================\n")
	previosDate = cDate
	indexRow++
	fmt.Printf("NEW nextPreviousDate: %s \n", cDate.String())
	fmt.Printf("=====================================================\n")
    }
}
