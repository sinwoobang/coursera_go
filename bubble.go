package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Swap(nums []int, i int) {
	for j := i + 1; j < len(nums); j++ {
		if (nums[i] > nums[j]) {
			tmp := nums[i];
			nums[i] = nums[j];
			nums[j] = tmp;
		}
	}
}


func BubbleSort(nums []int) {
	for i := 0; i < len(nums); i++ {
		Swap(nums, i);
	}
}


func main() {
	fmt.Println("Enter numbers: ")
	scanner := bufio.NewScanner(os.Stdin);
	scanner.Scan();
	text := strings.Fields(scanner.Text());
	
	var nums []int = make([]int, 0);
	for _, num_txt := range text {
		i, _ := strconv.Atoi(num_txt)
		nums = append(nums, i);
	}
	fmt.Printf("Raw data: %v\n", nums);

	BubbleSort(nums);
	fmt.Printf("Sorted data: %v\n", nums);
}