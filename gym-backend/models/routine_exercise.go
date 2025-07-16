package models

import "gorm.io/gorm"

type RoutineExercise struct {
	gorm.Model
	RoutineID  uint
	ExerciseID uint
	Sets       int
	Reps       int
	RestSecs   int
}