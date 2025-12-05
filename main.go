package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func init() {
	// Mute Gin console logs agar tidak mengganggu output CLI
	gin.SetMode(gin.ReleaseMode)
}

func main() {

	// Jalankan API di background
	go func() {
		r := SetupRouter()
		_ = r.Run(":8080") // mute server log
	}()

	// Jika ada file input
	if len(os.Args) > 1 {
		runFromFile(os.Args[1])
		return
	}

	// CLI Mode
	runInteractive()
}

func runInteractive() {
	reader := bufio.NewScanner(os.Stdin)
	fmt.Print(">> ")
	for reader.Scan() {
		processCLI(reader.Text())
		fmt.Print(">> ")
	}
}

func runFromFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Cannot open file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}
		processCLI(line)
	}
}

func processCLI(line string) {
	args := strings.Split(line, " ")
	cmd := strings.ToLower(args[0]) // FIX: Status → status

	switch cmd {

	case "create_parking_lot":
		size, _ := strconv.Atoi(args[1])
		NewParkingLot(size)

	case "park":
		reg := args[1]
		msg, _ := Park(reg)
		fmt.Println(msg)

	case "leave":
		reg := args[1]
		hours, _ := strconv.Atoi(args[2])
		msg, _ := Leave(reg, hours)
		fmt.Println(msg)

	case "status":
		fmt.Print(Status())

	default:
		fmt.Println("Invalid command")
	}
}
