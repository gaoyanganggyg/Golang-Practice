package bubblesort

func BubbleSort(data []int) []int {
	s_flag := true
	for i:=0; i<len(data)-1; i++{
		for j:=0; j<len(data)-i-1; j++{
			if data[j] > data[j+1]{
				data[j], data[j+1] = data[j+1], data[j]
				s_flag = false
			}
		}
		if s_flag{
			return data
		}
	}
	return data
}
