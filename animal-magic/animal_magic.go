package chance

import "math/rand"

// RollADie returns a random int d with 1 <= d <= 20.
func RollADie() int {

	num := 0
	for num <= 0 {
		num = rand.Intn(21)
	}

	return num
}

// GenerateWandEnergy returns a random float64 f with 0.0 <= f < 12.0.
func GenerateWandEnergy() float64 {
	return 12 * rand.Float64()
}

// ShuffleAnimals returns a slice with all eight animal strings in random order.
func ShuffleAnimals() []string {

	animalSlice := []string{"ant", "beaver", "cat", "dog", "elephant", "fox", "giraffe", "hedgehog"}

	rand.Shuffle(len(animalSlice), func(i, j int) {
		animalSlice[i], animalSlice[j] = animalSlice[j], animalSlice[i]
	})

	return animalSlice
}
