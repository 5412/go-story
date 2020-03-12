package qsort

import "testing"

func TestQuickSort1(t *testing.T) {
    values := []int{1, 2, 4, 3, 5}
    QuickSort(values)

    if values[2] != 3 || values[3] != 4 {
        t.Error("QuickSort() failed. Got", values, "Expected 1 2 3 4 5")
    }
}
