package main

import "testing"

func TestCanPlaceFlowers(t *testing.T) {

	checkResult := func(t testing.TB, got, want bool) {
		t.Helper()
		if got != want {
			t.Errorf("expected %t but got %t", want, got)
		}
	}

	t.Run("Case 1", func(t *testing.T) {
		flowerbed := []int{1, 0, 0, 0, 1}
		n := 1

		got := canPlaceFlowers(flowerbed, n)
		want := true

		checkResult(t, got, want)
	})

	t.Run("Case 2", func(t *testing.T) {
		flowerbed := []int{1, 0, 0, 0, 1}
		n := 2

		got := canPlaceFlowers(flowerbed, n)
		want := false

		checkResult(t, got, want)
	})
}
