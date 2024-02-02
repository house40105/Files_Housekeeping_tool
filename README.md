# Files_Housekeeping_tool
The Housekeeping Tool is a command-line utility written in Go, specifically crafted for efficient file maintenance tasks based on user-defined directory and count parameters. Developed with precision, it intelligently organizes files by Group ID (GID) and column (Type), selectively preserving the latest files within each category. Integrated seamlessly with 'crontab,' this tool enables automated and periodic file cleaning, ensuring a well-organized and clutter-free directory.


### Usage
1. Clone the repository:
   ```sh
   git clone https://github.com/house40105/Files_Housekeeping_tool.git
   ```
2. Navigate to the project directory:
   ```sh
   cd File_housekeeping_tool
   ```
3. Run the tool:
   ```sh
   go run main.go
   ```

### Configuration File (`modelHouseKeeping.ini`)
The INI file (`modelHouseKeeping.ini`) serves as the configuration file for the tool. It specifies the file directory and the keep count parameter.
```
[modelHouseKeeping]
modelDirectory = /path/to/models
keepCount = 5
```
* **modelDirectory:** The directory where model files are located.
* **keepCount:** The number of most recent files to retain for each GID and column.

### Code Structure
`main`
* Loads configuration parameters from the INI file.
* Extracts model directory and keep count from the configuration.
* Calls the `houseKeeping` function to perform file housekeeping.

`houseKeeping`
* Retrieves a list of files in the specified model directory.
* Categorizes files into a map using GID as the key and column as a secondary key.
* Iterates over each GID and column, sorts files, and retains only the specified count.
* Deletes older files beyond the specified count.

### File Naming Convention
The tool assumes a specific file naming convention: 
`file_<gid>_<column>_<date>.txt`. The `gid` is extracted as the second part, and the column is extracted as the third part of the filename. The date part is excluded for categorization.

### Customization
Modify the `houseKeeping` function based on specific naming conventions or categorization requirements.
Adjust the configuration parameters in `modelHouseKeeping.ini` to suit the desired model directory and keep count.
