package main

import "fmt"

func asteroidCollision(asteroids []int) []int {
	if len(asteroids) == 0 {
		return asteroids
	}
	// var stack []int
	stack := []int{asteroids[0]}
	var dontAddToStack bool
	for i := 1; i < len(asteroids); i++ {
		for len(stack) > 0 {
			isExplodable, isStackExploded, isCurrentExploded := isExplodable(stack[len(stack)-1], asteroids[i])
			dontAddToStack = isCurrentExploded
			if isExplodable {
				if isStackExploded && !isCurrentExploded {
					stack = stack[0 : len(stack)-1]
				} else if isCurrentExploded && !isStackExploded {
					break
				} else if isCurrentExploded && isStackExploded {
					stack = stack[0 : len(stack)-1]
					dontAddToStack = true
					break
				}
			} else {
				break
			}
		}
		if dontAddToStack == false {
			stack = append(stack, asteroids[i])
		}
		dontAddToStack = false
	}

	return stack
}

func AbsInt(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// stack => left
// curr => right
func isExplodable(leftValue, rightValue int) (bool, bool, bool) { // _, isStackExploded, isCurrExploded

	var leftDirection, rightDirection bool // true=right, false:left
	if leftValue > 0 {
		leftDirection = true
	} else {
		leftDirection = false
	}

	if rightValue > 0 {
		rightDirection = true
	} else {
		rightDirection = false
	}

	if leftDirection == true && rightDirection == false {
		if AbsInt(leftValue) > AbsInt(rightValue) {
			return true, false, true
		} else if AbsInt(leftValue) < AbsInt(rightValue) {
			return true, true, false
		} else {
			return true, true, true
		}
	} else {
		return false, false, false
	}
}

// stack => left
// curr => right
func main() {
	// fmt.Println(isExplodable(-2, -2))
	// fmt.Println(isExplodable(5, -6))
	// fmt.Println(isExplodable(3, -6))
	// fmt.Println(isExplodable(-6, 2))
	// fmt.Println(isExplodable(2, -1))
	// fmt.Println(isExplodable(2, 4))
	// fmt.Println(isExplodable(-6, 4))
	fmt.Println(asteroidCollision([]int{5, 10, -5}))
	fmt.Println(asteroidCollision([]int{8, -8}))
	fmt.Println(asteroidCollision([]int{10, 2, -5}))
	fmt.Println(asteroidCollision([]int{3, 5, -6, 2, -1, 4}))
	fmt.Println(asteroidCollision([]int{-2, -2, 1, -1}))
	fmt.Println(asteroidCollision([]int{1, -1, -2, -2}))
}
