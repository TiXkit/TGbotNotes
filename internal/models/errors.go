package models

import "errors"

var (
	EmailNotFound         = errors.New("Данного email нет в базе ")
	InvalidEmail          = errors.New("Данный email имеет некорректный вид ")
	RangeOverflow         = errors.New("Превышен максимальный диапазон значений при поиске email ")
	ErrorFromLocalStorage = errors.New("Произошла ошибка при попытке обращения к локальному файлу ")

	// опущенный функцинал, на осмыслении
	InvalidPathStorage = errors.New("По данному пути не был найден файл формата JSON ")
)
