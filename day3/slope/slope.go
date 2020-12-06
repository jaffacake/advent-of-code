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

func (s *slope) start() {
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

		s.currentPositionX++
		s.currentPositionY += 3
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

func Tobbogan(data []byte) int {
	s := initialiseSlope(data)

	fmt.Println(s.geometry)
	s.start()

	return s.treesHit
}
