package main

import (
    "fmt"
    "github.com/dantesbytes/nutritionstats/nutrition"
)

func main() {
    ns := nutrition.GetNutritionalScore(nutrition.NutritionalData{
        Energy:              nutrition.EnergyFromKcal(90),
        Sugars:              nutrition.SugarGram(9),
        SaturatedFattyAcids: nutrition.SaturatedFattyAcids(9),
        Sodium:              nutrition.SodiumMilligram(900),
        Fruits:              nutrition.FruitsPercent(70),
        Fiber:               nutrition.FiberGram(80),
        //Protien:             nutrition.ProtienGram(20),
    }, nutrition.Food)

    fmt.Printf("Nutritional Score: %d\n", ns.Value)
    fmt.Printf("NutriScore: %s\n", ns.GetNutriScore())
}
