package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

func readInput(colNum *int, rowNum *int, iterations *int) [][]bool {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	rowCount := 0
	for scanner.Scan() {
		rowCount++
	}
	*rowNum = rowCount - 1

	cor := make([][]bool, rowCount)
	file.Seek(0, 0)
	scanner2 := bufio.NewScanner(file)
	row := 0
	for scanner2.Scan() {
		if row != rowCount-1 {
			nodes := strings.Split(scanner2.Text(), " ")
			*colNum = len(nodes)
			cor[row] = make([]bool, len(nodes))
			for col := 0; col < len(nodes); col++ {
				cor[row][col] = nodes[col] == "x"
			}
			row++
		} else {
			*iterations, err = strconv.Atoi(scanner2.Text())
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return cor
}

func main() {
	colNum, rowNum, iterations := 0, 0, 0
	cor := readInput(&colNum, &rowNum, &iterations)

	ticker := time.NewTicker(500 * time.Millisecond)
	done := make(chan bool)
	start := true

	go func() {
		for {
			select {
			case <-done:
				return
			case <-ticker.C:
				if !start {
					fmt.Printf("\r\x1b[" + strconv.Itoa(rowNum) + "A")
				} else {
					start = false
				}

				fmt.Printf(gameStep(cor, colNum, rowNum))
			}
		}
	}()

	time.Sleep(time.Duration(iterations*500) * time.Millisecond)
	ticker.Stop()
	done <- true
}
