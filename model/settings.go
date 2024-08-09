package model

//Задача 7: Обнуление структуры
//Создайте структуру Settings с полями Volume, Brightness и Mode.
//Напишите функцию ResetSettings, которая принимает указатель на Settings
//и сбрасывает все поля в их нулевые значения.

type Settings struct {
	Volume     int
	Brightness int
	Mode       string
}

func ResetSettings(s *Settings) {
	s.Volume = 0
	s.Brightness = 0
	s.Mode = "Off"
}
