package sorts

import (
	"ranking/pods"
	"strconv"
)

func countSort(arr []pods.Pod, n int, exp int)  {
	output := make([]pods.Pod, n)
	count := make([]int, 10)

	// contiamo quante volte appare una certa
	// cifra
	for i := 0; i < n; i++ {

		val, _ := strconv.Atoi(arr[i].Vis)

		count[(val / exp) % 10]++
	}

	// count[i] contiene adesso la vera
	// posizione di i in output (l'ultima posizione)
	for i := 1; i < 10; i++ {
		count[i] += count[i - 1]
	}

	//build del output
	for i := n - 1; i >= 0; i-- {

		val, _ := strconv.Atoi(arr[i].Vis)

		output[count[(val / exp) % 10] - 1] = arr[i]
		count[(val / exp) % 10]--
	}

	for i := 0; i < n; i++ {
		arr[i] = output[i]
	}
}

// trova il massimo dal array arr con lunghezza n
func maxFromArray(arr []pods.Pod, n int) int {

	M := -1

	for i := 0; i < n; i++ {

		val, _ := strconv.Atoi(arr[i].Vis)

		if M < val {
			M = val
		}
	}

	return M
}

func RadixSort(arr []pods.Pod, n int) {
	// troviamo il numero massimo dal array
	// per sapere il numero massimo di cifre
	m := maxFromArray(arr, n)

	// facciamo countSort per ogni cifra
	// exp aiutera a trovare la cifra
	// (numero / 10^i) % 10 essendo la cifra
	// sulla posizione i
	for exp := 1; m / exp > 0; exp *= 10 {
		countSort(arr, n, exp)
	}
}

func MergeSort(arr1, arr2 []pods.Pod) []pods.Pod {
	i, j := 0, 0
	var sortedArr []pods.Pod

	for i < len(arr1) && j < len(arr2) {

		curI, _ := strconv.Atoi(arr1[i].Vis)
		curJ, _ := strconv.Atoi(arr2[j].Vis)

		if curI < curJ {
			sortedArr = append(sortedArr, arr1[i])
			i++
		} else {
			sortedArr = append(sortedArr, arr2[j])
			j++
		}
	}

	for i < len(arr1) {
		sortedArr = append(sortedArr, arr1[i])
		i++
	}

	for j < len(arr2) {
		sortedArr = append(sortedArr, arr2[j])
		j++
	}

	return sortedArr
}
