package util

import "errors"

func Search(element int, array []int) (int, error) {
	leftPosition, rightPosition := 0, len(array)

	for leftPosition < rightPosition {
		currentPosition := (leftPosition + (rightPosition - 1)) / 2
		if array[currentPosition] == element {
			return currentPosition, nil
		}

		if array[currentPosition] < element {
			leftPosition = currentPosition + 1
		} else {
			rightPosition = currentPosition
		}
	}
	return 0, errors.New("element is absent")
}
