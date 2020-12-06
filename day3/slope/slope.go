package slope

import (
	"fmt"
	"strings"
)

type slope struct {
	geometry map[int][]byte
	currentPositionX int
	currentPositionY int
	treesHit int
}

func (s *slope) start(right int, down int) {
	for {
		fmt.Println("pos: ", s.currentPositionX, s.currentPositionY)

		// Are we there yet?
		if len(s.geometry[s.currentPositionX]) == 0 {
			break
		}

		// Check if out of bounds
		if len(s.geometry[s.currentPositionX]) <= s.currentPositionY {
			s.currentPositionY -= len(s.geometry[s.currentPositionX])
		}

		if s.didHitTree() {
			s.treesHit++
		}

		s.currentPositionX += down
		s.currentPositionY += right
	}
}

func (s slope) didHitTree() bool {
	return string(s.geometry[s.currentPositionX][s.currentPositionY]) == "#"
}

func initialiseSlope(data []byte) slope {
	values := strings.Split(string(data), "\n")
	s := slope{geometry: map[int][]byte{}, treesHit: 0, currentPositionX: 0, currentPositionY: 0}

	for i, value := range values {
		if len(value) == 0 {
			break
		}

		s.geometry[i] = []byte(value)
	}

	return s
}

func (s *slope) reset() {
	s.treesHit = 0
	s.currentPositionX = 0
	s.currentPositionY = 0
}

func Tobbogan(data []byte) int {
	s := initialiseSlope(data)

	var slopes = make(map[int][]int)
	slopes[0] = []int{1,1}
	slopes[1] = []int{3,1}
	slopes[2] = []int{5,1}
	slopes[3] = []int{7,1}
	slopes[4] = []int{1,2}

	fmt.Println(s.geometry)
	var result [5]int
	for i, slope := range slopes {
		s.start(slope[0], slope[1])
		result[i] = s.treesHit
		s.reset()
	}


	return result[0] * result[1] * result[2] * result[3] * result[4]
}
