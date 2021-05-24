package main

func isAlive(cors [][]bool, row int, col int) bool {
	return cors[row][col]
}

func getAliveNeighborCount(cors [][]bool, row int, col int, colNum int, rowNum int) int {
	count := 0

	if row-1 >= 0 && isAlive(cors, row-1, col) {
		count++
	}

	if row+1 < rowNum && isAlive(cors, row+1, col) {
		count++
	}

	if row-1 >= 0 && col-1 >= 0 && isAlive(cors, row-1, col-1) {
		count++
	}

	if col-1 >= 0 && isAlive(cors, row, col-1) {
		count++
	}

	if row+1 < rowNum && col-1 >= 0 && isAlive(cors, row+1, col-1) {
		count++
	}

	if row-1 >= 0 && col+1 < colNum && isAlive(cors, row-1, col+1) {
		count++
	}

	if col+1 < colNum && isAlive(cors, row, col+1) {
		count++
	}

	if row+1 < rowNum && col+1 < colNum && isAlive(cors, row+1, col+1) {
		count++
	}

	return count
}

func calculateCooridnates(cor [][]bool, colNum int, rowNum int) {
	temp := make([][]bool, len(cor))
	for i := range cor {
		temp[i] = make([]bool, len(cor[i]))
		copy(temp[i], cor[i])
	}
	for row := 0; row < rowNum; row++ {
		for col := 0; col < colNum; col++ {
			aliveCount := getAliveNeighborCount(cor, row, col, colNum, rowNum)
			if isAlive(cor, row, col) && (aliveCount < 2 || aliveCount > 3) {
				temp[row][col] = false
			} else if !isAlive(cor, row, col) && aliveCount == 3 {
				temp[row][col] = true
			}
		}
	}

	for i := range temp {
		copy(cor[i], temp[i])
	}
}

func gameStep(cor [][]bool, colNum int, rowNum int) string {
	var gameBoard string
	for row := 0; row < rowNum; row++ {
		for col := 0; col < colNum; col++ {
			if isAlive(cor, row, col) {
				gameBoard += "\u25A0 "
			} else {
				gameBoard += "\u25A1 "
			}
		}
		gameBoard += "\n"
	}
	calculateCooridnates(cor, colNum, rowNum)
	return gameBoard
}