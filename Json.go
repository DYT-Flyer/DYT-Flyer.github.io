package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"

	"fyne.io/fyne/v2/canvas"
)

var dir, err = os.Getwd()

var tools []string = []string{"Adobe Firefly", "DeepDream Generator", "Artbreeder Mixer", "DeepAI", "Starry AI", "PicsArt AI Image Generator", "Stability AI", "Midjourney"}
var simple []int = []int{35, 23, 33, 22, 36, 42, 48, 41, 7, 30, 40, 24, 26, 1, 44, 8, 28, 49, 12, 5, 18, 4, 43, 27, 29, 17, 11, 6, 39, 47, 16, 14, 34, 10, 38, 25, 31, 2, 50, 37, 32, 45, 13, 21, 19, 20, 15, 46, 9, 3}
var complex []int = []int{21, 32, 12, 37, 2, 45, 7, 9, 5, 41, 8, 16, 31, 1, 40, 34, 38, 49, 27, 36, 6, 50, 47, 28, 17, 46, 11, 13, 23, 39, 19, 48, 4, 3, 15, 33, 20, 14, 44, 29, 24, 26, 30, 18, 35, 10, 25, 43, 22, 42}
var instructions []string = []string{
	"Rank Similarity to Prompt (1-8): Consider how well each image relates to the accompanying text. Assign a rating from 1 to 8, with 1 being the best (most similar) and 8 being the worst (least similar).",
	"Rank Aesthetics (1-8): Evaluate the visual appeal and aesthetics of each image, regardless of text. Assign a rating from 1 to 8, with 1 being the best (how beautiful the image is) and 8 being the worst (least aesthetically pleasing)",
	"Rank Realism (1-8): Determine the level of realism or believability in each image. Assign a rating from 1 to 8, with 1 being the best (most realistic) and 8 being the worst (least realistic).",
	"Rank Visual Quality (1-8): Assess the overall image quality, including factors like resolution, sharpness, and clarity. Assign a rating from 1 to 8, with 1 being the best (highest quality) and 8 being the worst lowest quality)."}

var userID string
var userIDFile string = dir + "/userID.txt"
var complexFolder string = dir + "/Dataset-Images/Complex/"
var simpleFolder string = dir + "/Dataset-Images/Simple/"

var fileIndex int
var userIndex int = 0

var simpleComplex = 0
var instruction int = 0

var imgs []*canvas.Image = make([]*canvas.Image, 8)
var permutation []int = []int{0, 1, 2, 3, 4, 5, 6, 7}
var userChoices []int = make([]int, 8)
var rows int

// List of functions:
// userIDExists()
// appendToFile()
// createString()
// saveResults()

// appendToFile() -> userID.txt
// saveResults() -> createString() -> saveResults() -> results.txt

// Checks to see if userID exists in file
func userIDExists() (bool, error) {
	f, err := os.Open(userIDFile)
	if err != nil {
		return false, err
	}
	defer f.Close()
	rows = 0

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		rows++
		if strings.TrimSpace(scanner.Text()) == userID {
			return true, nil
		}
	}
	if err := scanner.Err(); err != nil {
		return false, err
	}
	return false, nil
}

// Add userID to end of file
func appendToFile() error {
	// Open the file in append mode
	file, err := os.OpenFile(userIDFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	// Write the data to the file
	if _, err := file.Write([]byte(userID + "\n")); err != nil {
		if err != nil {
			return err
		}
	}

	return nil
}

func createString() string {
	// "Thomas Sherk; 2; Simple; Tree in Africa; Similarity; [1,2,3,4,5,6,7,8]; [8,2,1,3,4,5,6,7]; "
	// userID; fileIndex+userIndex; simpleComplex; simplePrompts; instruction; rankings; permutation
	var output string
	if simpleComplex == 0 {
		output = strconv.Itoa(rows) + ";" + userID + ";" + strconv.Itoa(fileIndex+userIndex) + ";" + strconv.Itoa(simpleComplex) + ";" + simplePrompts[simple[fileIndex+userIndex-1]-1] + ";" + strconv.Itoa(instruction) + ";" + "[" + strconv.Itoa(userChoices[0]) + "," + strconv.Itoa(userChoices[1]) + "," + strconv.Itoa(userChoices[2]) + "," + strconv.Itoa(userChoices[3]) + "," + strconv.Itoa(userChoices[4]) + "," + strconv.Itoa(userChoices[5]) + "," + strconv.Itoa(userChoices[6]) + "," + strconv.Itoa(userChoices[7]) + "]" + ";" + "[" + strconv.Itoa(permutation[0]) + "," + strconv.Itoa(permutation[1]) + "," + strconv.Itoa(permutation[2]) + "," + strconv.Itoa(permutation[3]) + "," + strconv.Itoa(permutation[4]) + "," + strconv.Itoa(permutation[5]) + "," + strconv.Itoa(permutation[6]) + "," + strconv.Itoa(permutation[7]) + "]; "
	} else {
		output = strconv.Itoa(rows) + ";" + userID + ";" + strconv.Itoa(fileIndex+userIndex) + ";" + strconv.Itoa(simpleComplex) + ";" + complexPrompts[complex[fileIndex+userIndex-1]-1] + ";" + strconv.Itoa(instruction) + ";" + "[" + strconv.Itoa(userChoices[0]) + "," + strconv.Itoa(userChoices[1]) + "," + strconv.Itoa(userChoices[2]) + "," + strconv.Itoa(userChoices[3]) + "," + strconv.Itoa(userChoices[4]) + "," + strconv.Itoa(userChoices[5]) + "," + strconv.Itoa(userChoices[6]) + "," + strconv.Itoa(userChoices[7]) + "]" + ";" + "[" + strconv.Itoa(permutation[0]) + "," + strconv.Itoa(permutation[1]) + "," + strconv.Itoa(permutation[2]) + "," + strconv.Itoa(permutation[3]) + "," + strconv.Itoa(permutation[4]) + "," + strconv.Itoa(permutation[5]) + "," + strconv.Itoa(permutation[6]) + "," + strconv.Itoa(permutation[7]) + "]; "
	}
	return output
}

func saveResults() {
	// Permutation = [8,2,1,3,4,5,6,7]
	// UserChoices = [1,2,3,4,5,6,7,8]
	// Actual Ranking = [3,2,4,5,6,7,8,1]
	// userID, File Index, Simple/Complex, Prompt, Instruction, Rankings, Permutation
	output := createString()

	file, err := os.OpenFile("results.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error opening or creating file:", err)
		return
	}
	defer file.Close()

	// Write the string to the file
	_, err = io.WriteString(file, output+"\n")
	if err != nil {
		fmt.Println("Error writing to file:", err)
	}
}
