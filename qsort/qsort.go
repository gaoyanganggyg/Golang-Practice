package qsort

func Qsort(data []int, left, right int) {
	tmp := data[left]
	p := left
	i, j := left, right
	for i <= j {
		for j >= p && data[j] > tmp {
			j--
		}
		for j >= p {
			data[j] = data[p]
			p = j
		}
		if data[i] <= tmp && i <=p {
			i++
		}
		if i <= p {
			data[p] = data[i]
			p = i
		}
		data[p] = tmp
		if p - left > 1 {
			Qsort(data, left, p-1)
		}
		if right - p > 1 {
			Qsort(data, p+1, right)
		}
	}

}
