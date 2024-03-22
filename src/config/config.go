package config

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type ConfigReader struct {
	data           []string
	gridSize       int
	numShip        int
	player1Ship    []string
	player2Ship    []string
	numMissile     int
	player1Missile []string
	player2Missile []string
}

func NewConfigReader() *ConfigReader {
	c := ConfigReader{}

	return &c
}

func (c *ConfigReader) ReadFile(filepath string) {
	fmt.Printf("\n\nReading a file in Go...\n")

	data, err := os.ReadFile(filepath)
	if err != nil {
		log.Panicf("failed reading data from file: %s", err)
	}
	fmt.Printf("\nFile Name: %s", filepath)
	fmt.Printf("\nSize: %d bytes", len(data))
	fmt.Printf("\nData: %s\n\n", data)

	c.data = strings.Split(string(data), "\n")

	// process data
	gridSizeint, _ := strconv.Atoi(c.data[0])
	numShipint, _ := strconv.Atoi(c.data[1])
	c.gridSize = gridSizeint
	c.numShip = numShipint

	c.player1Ship = strings.Split(c.data[2], ":")
	c.player2Ship = strings.Split(c.data[3], ":")
	numMissile, _ := strconv.Atoi(c.data[4])
	c.numMissile = numMissile
	c.player1Missile = strings.Split(c.data[5], ":")
	c.player2Missile = strings.Split(c.data[6], ":")
}

func (c ConfigReader) GetGridSize() int {
	return c.gridSize
}

func (c ConfigReader) GetNumShip() int {
	return c.numShip
}

func (c ConfigReader) GetNumMissile() int {
	return c.numMissile
}

func (c ConfigReader) GetPlayer1Ship() []string {
	return c.player1Ship
}

func (c ConfigReader) GetPlayer2Ship() []string {
	return c.player2Ship
}

func (c ConfigReader) GetPlayer1Missile() []string {
	return c.player1Missile
}

func (c ConfigReader) GetPlayer2Missile() []string {
	return c.player2Missile
}
