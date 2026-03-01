package main

func moveZeroes(nums []int) {
	insertPosition := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[insertPosition] = nums[i]
			insertPosition++
		}
	}

	for i := insertPosition; i < len(nums); i++ {
		nums[i] = 0
	}
}

func main() {
	moveZeroes([]int{0, 1, 0, 3, 12})
}
