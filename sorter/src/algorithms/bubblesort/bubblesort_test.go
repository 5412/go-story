package bubblesort

import "testing"

func TestBubbleSort1(t *testing.T) {
    values := []int{5, 6, 7, 4, 3, 2, 1}
    BubbleSort(values)
    if values[0] != 1 || values[1] != 2 || values[2] != 3 {
        t.Error("BubbleSort() failed. Got", values, "Expected 1 2 3 4 5 6 7")
    }
}
func TestBubbleSort2(t *testing.T) {
    values := []int{5, 6, 7, 4, 3, 2, 1}
    BubbleSort(values)
    if values[0] != 1 || values[1] != 2 || values[2] != 3 {
        t.Error("BubbleSort() failed. Got", values, "Expected 1 2 3 4 5 6 7")
    }
}
func TestBubbleSort3(t *testing.T) {
    values := []int{5, 6, 7, 4, 3, 2, 1}
    BubbleSort(values)
    if values[0] != 1 || values[1] != 2 || values[2] != 3 {
        t.Error("BubbleSort() failed. Got", values, "Expected 1 2 3 4 5 6 7")
    }
}
