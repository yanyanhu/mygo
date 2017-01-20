package main

import (
    "bufio"
    "errors"
    "flag"
    "fmt"
    "io"
    "os"
    "strconv"

    "algorithm/qsort"
)


var infile *string = flag.String("i", "infile", "File contains values for sorting")
var outfile *string = flag.String("o", "outfile", "File to receive sorted result")
var algorithm *string = flag.String("a", "qsort", "Sort algorithm")


func readValues(infile string) (values []int, err error) {
    file, err := os.Open(infile)
    if err != nil {
        fmt.Println("Failed to open the input file ", infile)
	return
    }

    defer file.Close()
    br := bufio.NewReader(file)
    values = make([]int, 0)

    fmt.Println("Reading lines...")
    for {
	line, isPrefix, err1 := br.ReadLine()

	if err1 != nil {
	    if err1 != io.EOF {
		err = err1
	    } else {
                fmt.Println("Reach the end of file, return.")
	    }
	    break
	}

	if isPrefix {
	    err_msg := "A too long line, seems unexpected."
	    fmt.Println(err_msg)
	    err = errors.New(err_msg)
	    break
	}

        str := string(line)

	value, err1 := strconv.Atoi(str)
	if err1 != nil {
	    err = err1
	    break
	}

	values = append(values, value)
    }

    return
}


func writeValues(values []int, outfile string) (err error) {
    file, err := os.Create(outfile)
    if err != nil {
	fmt.Println("Failed to create the output file ", outfile)
	return
    }

    defer file.Close()

    for _, value := range values {
	str := strconv.Itoa(value)
	file.WriteString(str + "\n")
    }

    return
}


func main() {
    flag.Parse()

    if infile != nil {
	fmt.Println("infile=", *infile, "outfile=", *outfile, "algorithm=", *algorithm)
    }

    values, err := readValues(*infile)
    if err != nil {
	fmt.Println(err)
    } else {
	fmt.Println("Read values: ", values)
    }

    // Sort the values
    if err != nil {
	fmt.Println(err)
    } else {
	switch *algorithm {
	    case "qsort":
		fmt.Println("Start sorting using algorith: ", *algorithm)
		qsort.QuickSort(values)
	    default:
		fmt.Println("Unsupported sort algorithm: ", *algorithm)
	}
    }
    fmt.Println("Sorting finished.")

    err = writeValues(values, *outfile)
    if err != nil {
	fmt.Println(err)
    } else {
	fmt.Println("Succeeded in writing sorted result to file ", *outfile, "\n")
    }
}
