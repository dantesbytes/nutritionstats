package main

type ScoreType int

const (

	Food ScoreType = iota
	Beverage
	Water
	Cheese
)

type NutritoinalScore struct {

	value int
	Positive int
	Negative int
	ScoreType ScoreType
}