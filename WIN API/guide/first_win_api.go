package main

import (
	"fmt" // "Системные вызовы". Основной инструмент для общения с ОС на низком уровне

	"golang.org/x/sys/windows"
)

var (
	showMessageBox bool = true
)

func main() {

	// ====================
	// |    MessageBox    |
	// ====================
	if showMessageBox {
		text, _ := windows.UTF16PtrFromString("TEST")
		caption, _ := windows.UTF16PtrFromString("HELLO!")

		retCode, err := windows.MessageBox(
			0,                // значение 0, означает что у нашего окна нет родителя, он независим
			text,             // Текст
			caption,          // Заголовок
			windows.MB_YESNO, // Используем готовую константу вместо "магического числа" 0x4
		)

		if retCode == 0 {
			fmt.Println("Error MessageBoxW:", err.Error())
		}

		fmt.Printf("Вы нажали кнопку с кодом: %d\n", retCode)
	}
}
