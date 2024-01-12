package lasagna

// TODO: define the 'PreparationTime()' function

func PreparationTime(layers []string, value int) int {
	numberOfLayers := len(layers)

	if value == 0 {
		return numberOfLayers * 2
	}

	return numberOfLayers * value
}

// TODO: define the 'Quantities()' function
func Quantities(slice []string) (int, float64) {
	wNoodles := 0
	wSauce := 0.0

	for i := 0; i < len(slice); i++ {
		if slice[i] == "sauce" {
			wSauce += 0.2
		} else if slice[i] == "noodles" {
			wNoodles += 50
		}
	}

	return wNoodles, wSauce
}

// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(friendsList []string, myList []string) {
	myList[len(myList)-1] = friendsList[len(friendsList)-1]
}

// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(slice []float64, people int) []float64 {

	var sliceNew []float64
	for i := 0; i < len(slice); i++ {
		sliceNew = append(sliceNew, slice[i]*(float64)(people/2))
	}

	return sliceNew
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.

