package mimik_core

import (
	"fmt"
	"mimik/internal/data"
	"mimik/internal/utils"
	"time"
)

func CreateNamespaceFiles(flags data.LaunchFlags) {
	fmt.Println("Creating Namespace files")

	var selectedData data.SelectedContent = utils.ExtractRequiredLocations(flags)

	if selectedData.Validated == true {

	}

	// Get the current time
	currentTime := time.Now()

	// Convert the current time to milliseconds
	currentTimeMillis := currentTime.UnixNano() / int64(time.Millisecond)

	fmt.Println("Current time in milliseconds:", currentTimeMillis)
}
