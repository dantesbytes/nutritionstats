package main

import (

	"fmt"
)

func main() {

	ns := GetNutritionalScore(NutritionalData{

		Energy: EnergyFromKcal(),
		Sugars: SugarGram(),
		SaturatedFattyAcids: SaturatedFattyAcids(),
		Sodium: SodiumMilligram(),
		Fruits: FruitsPercent(),
		Fiber: FiberGram(),
		Protien: ProtienGram(),



	}, Food)

	fmt.Printf("Nutritoinal Score: %d\n", ns.Value)
	fmt.Printf("NutriScore: %s\n", ns.GetNutritionalScore())
}