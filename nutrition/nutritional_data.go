package nutrition

type ScoreType int

const (

	Food ScoreType = iota
	Beverage 
	Water 
	Cheese
)



type EnergyKJ float64
type SugarGram float64
type SaturatedFattyAcids float64
type SodiumMilligram float64
type FruitsPercent float64
type FiberGram float64
type ProteinGram float64


type NutritionalData struct {

	Energy EnergyKJ
	Sugars SugarGram
	SaturatedFattyAcids SaturatedFattyAcids
	Sodium SodiumMilligram
	Fruits FruitsPercent
	Fiber FiberGram
	Protein ProteinGram
	IsWater bool

}
