package main

import (
	"fmt"
	"taskReceiver1/model"
)

//Задача 7: Обнуление структуры
//Создайте структуру Settings с полями Volume, Brightness и Mode.
//Напишите функцию ResetSettings, которая принимает указатель на Settings
//и сбрасывает все поля в их нулевые значения.

func main() {
	sett := &model.Settings{Volume: 100, Brightness: 99, Mode: "ON"}

	model.ResetSettings(sett)
	fmt.Printf("Volume: %d, Brightness: %d, Mode: %s\n", sett.Volume, sett.Brightness, sett.Mode)
}
