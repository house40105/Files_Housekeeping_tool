## generateTestFiles
The `generateTestFiles.go` is a Go program designed to generate test files with random properties and content based on a configuration file. This document provides comprehensive information on how to use, customize, and understand the functionality of the tool.

### Overview
The Model Housekeeping Tool is a utility that simplifies the process of generating diverse test files for scenarios requiring random GID (Group ID), category (Type), and date attributes. It utilizes a configuration file (`modelHouseKeeping.ini`) to specify parameters such as the export directory and the number of test files to generate.

### Getting Started
**Prerequisites**  
* Go installed on your machine.
* An INI configuration file (`modelHouseKeeping.ini`).

**Installation**
1. Clone the repository:
   ```sh
   git clone https://github.com/house40105/Files_Housekeeping_tool.git
   ```
2. Navigate to the project directory:
   ```sh
   cd Files_Housekeeping_tool/generate_test_files
   ```
3. Run the tool:
   ```sh
   go run generateTestFiles.go
   ```

### Configuration File (`modelHouseKeeping.ini`)
The INI file serves as the tool's configuration file, allowing users to customize the export directory and the number of test files to generate.
```
[generateTestModelFiles]
exportDirectory = /path/to/export
numFiles = 10
```
* exportDirectory: The directory where test files will be exported.
* numFiles: The number of test files to generate.

### Usage
1. Configure the `modelHouseKeeping.ini` file with the desired parameters.
2. Run the tool using the go run main.go command.

### Code Structure
`main`
* Loads configuration parameters from the INI file.
* Extracts export directory and the number of files to generate.
* Calls the generateTestFiles function to generate test files.

`generateTestFiles`
* Generates test files with random GID (gid), category (type), and date (format: ).
* Writes the generated content to each file.
* Prints a message for each generated file.

`getRandomCategory`
* Returns a random category (type) from a predefined list.

`getRandomDate`
* Generates a random date within a specified range and returns it in a formatted string.

### Customization
* Modify the `getRandomCategory` function to include additional categories if needed.
* Adjust the date range in the `getRandomDate` function based on specific testing requirements.
