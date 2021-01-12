package sort

func MergerArray(array []int, left int, right int, tmp []int) {
	if left >= right {
		return
	} else {
		index := 0
		mid := left + (right-left)>>1
		MergerArray(array, left, mid, tmp)
		MergerArray(array, mid+1, right, tmp)

		begin1, end1 := left, mid
		begin2, end2 := mid+1, right
		for begin1 <= end1 && begin2 <= end2 {
			if array[begin1] < array[begin2] {
				tmp[index] = array[begin1]
				begin1++
			} else {
				tmp[index] = array[begin2]
				begin2++
			}
			index++
		}
		for begin1 <= end1 {
			tmp[index] = array[begin1]
			index++
			begin1++
		}
		for begin2 <= end2 {
			tmp[index] = array[begin2]
			index++
			begin2++
		}
		for i := 0; i < index; i++ {
			array[left+i] = tmp[i]
		}
	}
}

func MergerSort(array []int) {
	tmp := make([]int, len(array))
	MergerArray(array, 0, len(array)-1, tmp)
}

