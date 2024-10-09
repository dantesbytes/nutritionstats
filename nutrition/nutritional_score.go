package nutrition

type NutritionalScore struct {
	Value     int
	Positive  int
	Negative  int
	ScoreType ScoreType
}

// overal grade of your nutrition score 

var scoreToLetter = []string{"A","B","C","D","E"}

func GetNutritionalScore(n NutritionalData, st ScoreType) NutritionalScore {

	
	value := 0
	positive := 0
	negative := 0

	if st != Water {
		fruitPoints := n.Fruits.GetPoints(st)
		fiberPoints := n.Fiber.GetPoints(st)

		negative = n.Energy.GetPoints(st) + n.Sugars.GetPoints(st) + n.SaturatedFattyAcids.GetPoints(st) + n.Sodium.GetPoints(st)
		positive = fruitPoints + fiberPoints + n.Protein.GetPoints(st)

		if st == Cheese{
			value = negative - positive

		} else {
			if negative >= 11 && fruitPoints < 5 {
				value = negative - positive - fruitPoints
			} else {

				value = negative - positive
			}
		}
	}

	return NutritionalScore {
		Value: value,
		Positive: positive,
		Negative: negative,
		ScoreType: st,
	}

}

func (ns NutritionalScore) GetNutriScore() string {

	if ns.ScoreType == Food {
		return scoreToLetter[getPointsFromRange(float64(ns.Value), []float64{18,19,2,-1})]
	}

	


	if ns.ScoreType == Water {
		return scoreToLetter[0]
	}

	return scoreToLetter[getPointsFromRange(float64(ns.Value), []float64{9, 5,2,-1})]

}
