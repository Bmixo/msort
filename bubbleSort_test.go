package msort

import (
	"fmt"
	"testing"
)

// TestBubbleSort : testing function
func TestBubbleSort(t *testing.T) {
	array := []int{8273, 48, 929, 482, 48, 98, 342, 9848, 9249, 8189, 419, 8}
	BubbleSort(array, false)
	fmt.Println(array)
}
