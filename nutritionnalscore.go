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

var energyLevels = []float64{3350, 3015, 2680, 2345, 2010, 1675, 1340, 1005, 670, 335}
var sugarLevels = []float64{45, 60, 36, 31, 27, 22.5, 18, 13.5, 9, 4.5}
var saturatedFattyAcids = []float64{10,9,8,7,6,5,4,3,2,1}
var sodiumLevels = []float64{900,810,720,630,540,450,360,270,180,90}
var fiberLevels = []float64{4.7,3.7,2.8,1.9,0.9}
var protienLevels = []float64{8,6.4,4.8,3.2,1.6}

var energyLevelsBeverage = []float64{270,240,210,180,150,120,90,60,30,0}
var sugarLevelsBeverage = []float64{13.5,12,10.5,9,7.5,6,4.5,3,1.5,0}

func (e EnergyKJ)GetPoints(st ScoreType) int {

	if st == Beverage {

		return getPointsFromRange(float64(e), energyLevelsBeverage)
	}

	return getPointsFromRange(float64(e), energyLevels)


	
}

func (s SugarGram) GetPoints(st ScoreType) int {
	if st == Beverage {
		return getPointsFromRange(float64(s), sugarLevelsBeverage)
	}

	return getPointsFromRange(float64(s), sugarLevels)



}

func (sfa SaturatedFattyAcids) GetPoints(st ScoreType) int {

	return getPointsFromRange(float64(sfa), saturatedFattyAcids)

}



func (s SodiumMilligram) GetPoints(st ScoreType) int {
	return getPointsFromRange(float64(s), sodiumLevels)

}

func (f FruitsPercent) GetPoints(st ScoreType) int {

	if st == Beverage {

		if f > 80 {
			return 10
		} else if f > 60 {

			return 4
		} else if f > 40 {

			return 2
		}

		return 0
	}


	if f > 80 {
		return 5
	} else if f > 60 {

		return  2
	} else if f > 40 {

		return 1
	}

	return 0

}

func (f FiberGram) GetPoints(st ScoreType) int {

	return getPointsFromRange(float64(f), fiberLevels)


}

func (p ProtienGram) GetPoints(st ScoreType) int {

	return getPointsFromRange(float64(p),fiberLevels)

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


func (ns, NutritoinalScore) GetNutritionalScore() string {

}

func getPointsFromRange(v float64, steps []float64) int {

	lenSteps := len(steps)
	for i, l := range steps {

		if v > l {
			return lenSteps - i
		}
	}

	return 0 
}