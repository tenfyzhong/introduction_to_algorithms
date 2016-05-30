/**
 * @file insert_sort_test.go
 * @brief
 * @author tenfyzhong
 * @email 364755805@qq.com
 * @created 2016-05-30 08:54:17
 */

package ch1

import "testing"

func TestInsertSortNormal(t *testing.T) {
	arr := []int{1, 2, 3}
	result := []int{1, 2, 3}
	sortarr := InsertSort(arr)
	if len(sortarr) != 3 {
		t.Error(sortarr)
	}

	for i := 0; i < len(sortarr); i = i + 1 {
		if result[i] != sortarr[i] {
			t.Error("sort error")
		}
	}
}
