package main

func main() {

	gelen := make([]int, 0, 10)
	gelen = append(gelen, 6, 4, 1, 3)

	BobbleSort(gelen)

}

func BobbleSort(gelen []int) {

	for i := 0; i < len(gelen); i++ {

		for j := 0; j < len(gelen)-1-i; j++ {
			if gelen[j] > gelen[j+1] {
				temp := gelen[j]
				gelen[j] = gelen[j+1]
				gelen[j+1] = temp
			}
		}
	}
	for k := 0; k < len(gelen); k++ {
		print(gelen[k])
		print("\n")
	}
}
