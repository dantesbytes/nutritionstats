package main

type ScoreType int

const (

	Food ScoreType = iota
	Beverage
	Water
	Cheese
)

type NutritoinalScore struct {

	Value int
	Positive int
	Negative int
	ScoreType ScoreType
}

type EnergyKJ float64
type SugarGram float64
type SaturatedFattyAcids float64
type SodiumMilligram float64
type FruitsPercent float64
type FiberGram float64
type ProtienGram float64


type NutritionalData struct {

	Energy EnergyKJ
	Sugars SugarGram
	SaturatedFattyAcids SaturatedFattyAcids
	Sodium SodiumMilligram
	Fruits FruitsPercent
	Fiber FiberGram
	Protien ProtienGram
	IsWater bool

}

var energyLevels = []float64{}
var sugarLevels = []float64{}
var saturatedFattyAcids = []float64{}
var sodiumLevels = []float64{}
var fiberLevels = []float64{}
var protienLevels = []float64{}

var energyLevelsBeverage = []float64{}
var sugarLevelsBeverage = []float64{}

func (e EnergyKJ)GetPoints(st ScoreType) int {


	
}

func (s SugarGram) GetPoints(st ScoreType) int {

}

func (sfa SaturatedFattyAcids) GetPoints(st ScoreType) int {

}



func (s SodiumMilligram) GetPoints(st ScoreType) int {

}

func (f FruitsPercent) GetPoints(st ScoreType) int {

}

func (f FiberGram) GetPoints(st ScoreType) int {


}

func (p ProtienGram) GetPoints(st ScoreType) int {

}

func EnergyFormKcal(kcal float64) EnergyKJ{

	return EnergyKJ(kcal* 4.184) 

}

func SodiumFromSalt(saltMg float64) SodiumMilligram {

	return SodiumMilligram(saltMg/2.5)

}



func GetNutritionalScore(n NutritionalData,st ScoreType) NutritoinalScore {

	
	value := 0
	positive := 0
	negative := 0

	if st != Water {
		fruitPoints := n.Fruits.GetPoints(st)
		fiberPoints := n.Fiber.GetPoints(st)

		negative = n.Energy.GetPoints(st) + n.Sugars.GetPoints(st) + n.SaturatedFattyAcids.GetPoints(st) + n.Sodium.GetPoints(st)
		positive = fruitPoints + fiberPoints + n.Protien.GetPoints(st)
	}

	return NutritoinalScore{
		Value: value,
		Positive: positive,
		Negative: negative,
		ScoreType: st,
	}

}
