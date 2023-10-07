package mimik_core

import (
	"fmt"
	"mimik/internal/data"
	"mimik/internal/utils"
)

func CreateNamespaceFiles(flags data.LaunchFlags) {
	fmt.Println("Creating Namespace files")

	var selectedData data.SelectedContent = utils.ExtractRequiredLocations(flags)

	if selectedData.Validated == true {

	}
}
