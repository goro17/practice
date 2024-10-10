package main

func canPlaceFlowers(flowerbed []int, n int) bool {

	// count variable to calculate max flowers we can plant
	count := 0

	if flowerbed[0] == 0 && len(flowerbed) == 1 {
		count++
	}

	// Case 1: you can plant a flower on the first position if there
	// are no flowers in the first 2 positions (e.g., 001...)
	if len(flowerbed) > 1 && flowerbed[0] == 0 && flowerbed[1] == 0 {
		count++
		flowerbed[0] = 1
	}

	// Case 2: you can plant a flower in the middle if there are 3
	// consecutive zeroes (e.g., ...010...)
	for i := 1; i < len(flowerbed)-1; i++ {
		if flowerbed[i] == 0 && flowerbed[i-1] == 0 && flowerbed[i+1] == 0 {
			flowerbed[i] = 1
			count++
		}
	}

	// Case 3: you can plant a flower on the last position if there
	// are no flowers in the last 2 positions (e.g., ...100)
	if len(flowerbed) > 2 && flowerbed[len(flowerbed)-2] == 0 && flowerbed[len(flowerbed)-1] == 0 {
		count++
	}

	return count >= n
}
