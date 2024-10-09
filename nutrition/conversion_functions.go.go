package nutrition


func EnergyFromKcal(kcal float64) EnergyKJ {

	return EnergyKJ(kcal * 4.184)
}


func SodiumFromSalt(saltMg float64) SodiumMilligram {

	return SodiumMilligram(saltMg/2.5)

}
