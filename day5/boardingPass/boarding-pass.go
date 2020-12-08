package boardingPass

import (
	"sort"
	"strings"
)

func CalculateSeat(data []byte) int {
	values := strings.Split(string(data), "\n")

	var seatIds []int

	for _, boardingPass := range values {
		rows := make([]int, 128)
		for i := 0; i < len(rows); i++ {
			rows[i] = i
		}

		seats := make([]int, 8)
		for i := 0; i < len(seats); i++ {
			seats[i] = i
		}

		for i, _ := range []byte(boardingPass) {
			if string(boardingPass[i]) == "F" {
				rows = rows[:len(rows)/2]
			} else if string(boardingPass[i]) == "B" {
				rows = rows[len(rows)/2:]
			} else if string(boardingPass[i]) == "L" {
				seats = seats[:len(seats)/2]
			} else if string(boardingPass[i]) == "R" {
				seats = seats[len(seats)/2:]
			}
		}

		seatId := rows[0] * 8 + seats[0]

		seatIds = append(seatIds, seatId)
	}

	sort.Ints(seatIds)

	return seatIds[len(seatIds)-1]
}