package main

import (
	"bytes"
	"encoding/binary"
	"log"
	"math/big"
	"os"
	"github.com/fatih/color"
)

// IntToHex converts an int64 to a byte array
func IntToHex(num int64) []byte {
	buff := new(bytes.Buffer)
	err := binary.Write(buff, binary.BigEndian, num)
	if err != nil {
		log.Panic(err)
	}

	return buff.Bytes()
}

// ReverseBytes reverses a byte array
func ReverseBytes(data []byte) {
	for i, j := 0, len(data)-1; i < j; i, j = i+1, j-1 {
		data[i], data[j] = data[j], data[i]
	}
}

// Equal tells whether a and b contain the same elements.
// A nil argument is equivalent to an empty slice.
func Equal(a, b []byte) bool {
    if len(a) != len(b) {
        return false
    }
    for i, v := range a {
        if v != b[i] {
            return false
        }
    }
    return true
}

func targetToDifficultyFloat64(target *big.Int) float64 {
	maxTarget := targetFromTargetBits(0)
	return float64(new(big.Int).Div(maxTarget, target).Uint64())
}

func invertArray(s []float64) []float64 {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
        s[i], s[j] = s[j], s[i]
    }
    return s
}

func targetToDifficulty(target *big.Int) *big.Int {
	maxTarget := targetFromTargetBits(0)
	return new(big.Int).Div(maxTarget, target)
}

func difficultyToTarget(difficulty *big.Int) *big.Int {
	maxTarget := targetFromTargetBits(0)
	return new(big.Int).Div(maxTarget, difficulty)
}

func targetFromTargetBits(targetBits int) *big.Int {
	target := big.NewInt(1)
	target.Lsh(target, uint(256 - targetBits))
	return target
}

func bigFloatToBigInt (f *big.Float) *big.Int {
	a := new(big.Int) 
	z, _ := f.Int(a)
	return z
}

func bigIntToBigFloat (i *big.Int) *big.Float {
	return new(big.Float).SetInt(i)
}

func printGreen(text string) {
	color.Green(text)
}

func printRed(text string) {
	color.Red(text)
}

func printBlue(text string) {
	color.Blue(text)
}

func printYellow(text string) {
	color.Yellow(text)
}

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func WriteToFile(filename string, text string) {
	f, _ := os.Create(filename)

	f, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0600)
	if err != nil {
	    panic(err)
	}

	defer f.Close()

	if _, err = f.WriteString(text); err != nil {
	    panic(err)
	}
}
