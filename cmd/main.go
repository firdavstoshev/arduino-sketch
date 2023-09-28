package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/tarm/serial"
	"log"
)

var (
	Port *serial.Port
	Flag = false
)

func main() {
	// Открываем соединение с Serial Port (замените порт и скорость на свои)
	port, err := serial.OpenPort(&serial.Config{
		Name: "COM1", // Windows: "COM1"
		Baud: 9600,   // Скорость передачи
	})
	if err != nil {
		fmt.Println("Ошибка открытия порта:", err)
		return
	}
	Port = port
	defer func(port *serial.Port) {
		if err := port.Close(); err != nil {
			fmt.Println("Ошибка закрытия порта:", err)
		}
		fmt.Println("Порт успешно закрыт!")
	}(port)

	r := gin.Default()
	r.GET("/signal", SendSignal)
	log.Fatal(r.Run("localhost:8080"))
}

func SendSignal(c *gin.Context) {
	if Flag {
		if _, err := Port.Write([]byte("0")); err != nil {
			fmt.Println(err)
			return
		}
		Flag = false
		c.JSON(200, gin.H{"message": "Успешно выключен!"})
	} else {
		if _, err := Port.Write([]byte("1")); err != nil {
			fmt.Println(err)
			return
		}
		Flag = true
		c.JSON(200, gin.H{"message": "Успешно включен!"})
	}

}
