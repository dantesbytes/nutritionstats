package main

import (

	"fmt"
)

func main() {

	ns := GetNutritionalScore(NutritionalData{

		Energy: EnergyFromKcal(0),
		Sugars: SugarGram(9),
		SaturatedFattyAcids: SaturatedFattyAcids(9),
		Sodium: SodiumMilligram(900),
		Fruits: FruitsPercent(70),
		Fiber: FiberGram(80),
		Protien: ProtienGram(20),



	}, Food)

	fmt.Printf("Nutritoinal Score: %d\n", ns.Value)
	fmt.Printf("NutriScore: %s\n", ns.GetNutritionalScore(8))
}