package main

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"gopkg.in/ini.v1"
)

const cfgPath = "modelHouseKeeping.ini"

func main() {
	// Load configuration from the specified INI file path
	inidata, err := ini.Load(cfgPath)
	if err != nil {
		fmt.Printf("Fail to read ini file: %v", err)
		os.Exit(1)
	}
	// Extract model directory and keep count from the configuration
	directory := inidata.Section("modelHouseKeeping").Key("modelDirectory").String()
	keepCount := inidata.Section("modelHouseKeeping").Key("keepCount").MustInt(1)
	fmt.Printf("Keep Count:%d\nDirectory: %s\n", keepCount, directory)

	// Perform housekeeping based on the extracted parameters
	fmt.Println("Start!")
	if keepCount >= 0 {
		err := houseKeeping(directory, keepCount)
		if err != nil {
			fmt.Println("Error:", err)
		}
	} else {
		fmt.Printf("Invalid keepCount: %d\n", keepCount)
	}
	fmt.Println("End!")
}

// houseKeeping performs file housekeeping based on specified directory and keep count
func houseKeeping(directory string, keepCount int) error {
	// Retrieve a list of files in the specified directory
	files, err := os.ReadDir(directory)
	if err != nil {
		return err
	}

	// Initialize a map to categorize files by GID and column
	fileMap := make(map[string]map[string][]string)

	// Categorize files into the map using GID as the key and column as a secondary key
	for _, file := range files {
		if !file.IsDir() {
			fileName := file.Name()
			parts := strings.Split(fileName, "_")

			if len(parts) == 4 {
				gid := parts[1]
				column := parts[2] // Category, excluding the date part
				if _, ok := fileMap[gid]; !ok {
					fileMap[gid] = make(map[string][]string)
				}
				fileMap[gid][column] = append(fileMap[gid][column], fileName)
			}
		}
	}

	// Iterate over each GID and column, sort files, and keep only the specified count
	for gid, columns := range fileMap {
		for column, files := range columns {
			fmt.Println("---------------------------------")
			fmt.Printf("Gid: %s,\tColumn: %s,\tFile count: %d\nFiles: %v\n", gid, column, len(files), files)

			if len(files) > keepCount {
				sort.Strings(files)

				// Delete older files beyond the specified count
				deleteCount := len(files) - keepCount
				for _, file := range files[:deleteCount] {
					filePath := filepath.Join(directory, file)

					// Check if the file exists before attempting deletion
					if _, err := os.Stat(filePath); err == nil {
						if err := os.Remove(filePath); err != nil {
							fmt.Printf("Error deleting file %s: %v\n", file, err)
						} else {
							fmt.Printf("Deleted: %s\n", file)
						}
					}
				}

				fmt.Printf("Keeping latest %d files for gid %s and column %s: %v\n", keepCount, gid, column, files[deleteCount:])
			} else {
				fmt.Printf("Not enough files for gid %s and column %s. Skipping...\n", gid, column)
			}
		}
	}

	return nil
}
