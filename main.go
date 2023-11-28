package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	// "golang.org/x/text/cases"
)

type Grade string

const (
	A Grade = "A"
	B Grade = "B"
	C Grade = "C"
	F Grade = "F"
)

type student struct {
	firstName, lastName, university                string
	test1Score, test2Score, test3Score, test4Score int
}

type studentStat struct {
	student
	finalScore float32
	grade      Grade
}

func parseCSV(filePath string) ([]student, error) {
	// Pending
	// Defer block
	// Closing the csv

	file, err := os.Open(filePath)
	// fmt.Print(dat)
	// dat.read
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var students []student

	count := 0

	for scanner.Scan() {
		count++
		line := scanner.Text()
		if count > 1 {
			// fmt.Print(line)
			fields := strings.Split(line, ",")
			field3, err1 := strconv.Atoi(fields[3])

			field4, err2 := strconv.Atoi(fields[4])

			field5, err3 := strconv.Atoi(fields[5])

			field6, err4 := strconv.Atoi(fields[6])

			if err1 != nil || err2 != nil || err3 != nil || err4 != nil {

				return []student{}, errors.New("Invalid CSV data")
			}
			student := student{fields[0], fields[1], fields[2], field3, field4, field5, field6}
			students = append(students, student)
		}
		// fmt.Print(line)

	}

	return students, nil
}

func calculateGrade(students []student) []studentStat {

	// Avg of t1 to t4
	// Then bunch of if else cond on avg score

	var studentStats []studentStat
	for _, val := range students {

		AverageScore := (val.test1Score + val.test2Score + val.test3Score + val.test4Score) / 4
		floatValueOfAvg := float32(AverageScore)
		if AverageScore < 35 {
			studentStat := studentStat{val, floatValueOfAvg, F}
			studentStats = append(studentStats, studentStat)
		}
		if AverageScore >= 35 && AverageScore < 50 {
			studentStat := studentStat{val, floatValueOfAvg, C}
			studentStats = append(studentStats, studentStat)
		}
		if AverageScore >= 50 && AverageScore < 70 {
			studentStat := studentStat{val, floatValueOfAvg, B}
			studentStats = append(studentStats, studentStat)
		}
		if AverageScore >= 70 {
			studentStat := studentStat{val, floatValueOfAvg, A}
			studentStats = append(studentStats, studentStat)
		}

	}

	return studentStats
}

func findOverallTopper(gradedStudents []studentStat) studentStat {
	var Topper float32
	Topper = 0
	var NameOfTopper studentStat
	for _, val := range gradedStudents {

		if val.finalScore > Topper {
			Topper = val.finalScore
			NameOfTopper = val
		}

	}
	// fmt.Print("OverAll Topper is ", NameOfTopper)
	return NameOfTopper
}

func findTopperPerUniversity(gs []studentStat) map[string]studentStat {
	return nil
}

func main() {
	// fmt.Print(parseCSV("grades.csv"))
	student, err := parseCSV("grades.csv")
	if err != nil {
		fmt.Print("We have a err", err)
	}
	gradedStudents := calculateGrade(student)
	findOverallTopper(gradedStudents)
}
