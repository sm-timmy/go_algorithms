package main

import "fmt"

func main() {
	//Один из вариантов сортировки
	arr := []int{1, 2, 3, 33, 12, 14, 461, 11, 1, 144}
	fmt.Println(arr)
	arr = quickSortMid(arr)
	fmt.Println(arr)

	//Один из вариантов поиска по массиву (слайсу)
	searchItem := 33
	res := binarySearchRecursive(arr, 0, len(arr)-1, searchItem)
	if res == -1 {
		fmt.Println("Not found!")
	} else {
		fmt.Printf("Search item %v found at index %v\n", searchItem, res)
	}

	//Поиск в глубину
	g := Graph{nodes: []*Node{
		{Name: "A",
			Children: []*Node{
				&Node{
					Name: "B",
					Children: []*Node{
						&Node{Name: "C",
							Children: nil,
						},
					},
				},
				&Node{
					Name:     "D",
					Children: nil,
				},
			},
		},
	}}
	fmt.Println(g.nodes[0].dfs("C"))

	//Поиск в ширину
	graph := [][]int{
		{0, 1, 1, 1, 0},
		{1, 0, 1, 0, 0},
		{1, 1, 0, 0, 0},
		{1, 0, 0, 0, 1},
		{0, 0, 0, 1, 0},
	}
	bfs(graph, 0)
}

func bubbleSort(arr []int) {
	for i := 0; i < len(arr)-1; i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i] > arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}

}

func selectionSort(arr []int) {
	for i := 0; i < len(arr)-1; i++ {
		minIdx := i
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[minIdx] {
				minIdx = j
			}
		}
		arr[i], arr[minIdx] = arr[minIdx], arr[i]
	}
}

func merge(l []int, r []int) []int {
	var merged []int
	for len(l) > 0 && len(r) > 0 {
		if l[0] < r[0] {
			merged = append(merged, l[0])
			l = l[1:]
		} else {
			merged = append(merged, r[0])
			r = r[1:]
		}
	}
	merged = append(merged, l...)
	merged = append(merged, r...)
	return merged
}

func mergeSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}
	mid := len(arr) / 2
	leftArr := mergeSort(arr[:mid])
	rightArr := mergeSort(arr[mid:])
	return merge(leftArr, rightArr)
}

func combSort(arr []int) []int {
	l := len(arr)
	factor := 1.247
	gap := int(float64(l) / factor)
	for gap > 0 {
		for i := 0; i < l && gap+i < l; i += gap {
			if arr[i] > arr[gap+i] {
				arr[i], arr[gap+i] = arr[gap+i], arr[i]
			}
		}
		gap--
	}
	return arr
}

func quickSort(arr []int) []int {
	var merged []int
	if len(arr) < 2 {
		return arr
	}
	pivot := arr[0]
	var left []int
	var right []int
	for i := 1; i < len(arr); i++ {
		if pivot > arr[i] {
			left = append(left, arr[i])
		} else {
			right = append(right, arr[i])
		}
	}
	merged = append(merged, quickSort(left)...)
	merged = append(merged, pivot)
	merged = append(merged, quickSort(right)...)
	return merged
}

func quickSortMid(arr []int) []int {
	var merged []int
	lenArr := len(arr)
	if lenArr < 2 {
		return arr
	}
	pivot := arr[lenArr/2]
	var left []int
	var right []int
	for i := 0; i < lenArr; i++ {
		if i == lenArr/2 {
			i++
		}
		if pivot > arr[i] {
			left = append(left, arr[i])
		} else {
			right = append(right, arr[i])
		}
	}
	merged = append(merged, quickSort(left)...)
	merged = append(merged, pivot)
	merged = append(merged, quickSort(right)...)
	return merged
}

func linearSearch(arr []int, searchItem int) int {
	for i := 0; i < len(arr); i++ {
		if arr[i] == searchItem {
			return i
		}
	}
	return -1
}

func binarySearch(arr []int, searchItem int) int {
	lenArr := len(arr)
	firstIdx := 0
	lastIdx := lenArr - 1
	for firstIdx != lastIdx {
		mid := (firstIdx + lastIdx) / 2
		if arr[mid] == searchItem {
			return mid
		} else if arr[mid] < searchItem {
			firstIdx = mid + 1
		} else {
			lastIdx = mid - 1
		}
	}
	return -1
}

func binarySearchRecursive(arr []int, firstIdx int, lastIdx int, searchItem int) int {
	if firstIdx <= lastIdx {
		mid := firstIdx + (lastIdx-firstIdx)/2
		if arr[mid] == searchItem {
			return mid
		} else if arr[mid] < searchItem {
			return binarySearchRecursive(arr, mid+1, lastIdx, searchItem)
		} else {
			return binarySearchRecursive(arr, firstIdx, mid-1, searchItem)
		}
	}
	return -1
}
