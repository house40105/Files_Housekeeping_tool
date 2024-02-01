package main

import (
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
	"time"

	"gopkg.in/ini.v1"
)

const cfgPath = "modelHouseKeeping.ini"

// const cfgPath = "/home/miot/modelHouseKeeping.ini"

func main() {
	// Load configuration from the specified INI file path
	inidata, err := ini.Load(cfgPath)
	if err != nil {
		fmt.Printf("Fail to read ini file: %v", err)
		os.Exit(1)
	}
	// Extract export directory and number of files to generate from the configuration
	directory := inidata.Section("generateTestModelFiles").Key("exportDirectory").String()
	numFiles := inidata.Section("generateTestModelFiles").Key("numFiles").MustInt(1)

	// Generate test files with the specified parameters
	err = generateTestFiles(directory, numFiles)
	if err != nil {
		fmt.Println("Error:", err)
	}
}

// generateTestModelFiles generates test files with random properties and content
func generateTestFiles(directory string, numFiles int) error {
	// Initialize the random number generator with the current time as the seed
	source := rand.NewSource(time.Now().UnixNano())
	random := rand.New(source)

	for i := 1; i <= numFiles; i++ {
		// Generate random values for GID, category, and date
		gid := random.Intn(3) + 91 // Randomly generate GID
		category := getRandomCategory()
		date := getRandomDate()

		// Create a file name based on the generated properties
		fileName := fmt.Sprintf("file_%d_%s_%s.txt", gid, category, date)
		filePath := filepath.Join(directory, fileName)

		// Generate content for the file
		content := fmt.Sprintf("Content for %s", fileName)
		// Write the content to the file
		err := os.WriteFile(filePath, []byte(content), 0644)
		if err != nil {
			return err
		}

		fmt.Printf("Generated test file: %s\n", fileName)
	}

	return nil
}

// getRandomCategory returns a random category from a predefined list
func getRandomCategory() string {
	categories := []string{"tpye1", "type2"}
	return categories[rand.Intn(len(categories))]
}

// getRandomDate generates a random date within a specified range and returns it in a formatted string
func getRandomDate() string {
	// Define the start and end dates for the random date generation
	start := time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC).Unix()
	end := time.Date(2023, 12, 31, 23, 59, 59, 0, time.UTC).Unix()
	source := rand.NewSource(time.Now().UnixNano())
	random := rand.New(source)

	randomTime := time.Unix(random.Int63n(end-start)+start, 0)

	return randomTime.Format("20060102")
}
