package data

type LaunchFlags struct {
	Clusters  uint16
	RunMimik  bool
	Days      int
	DataPath  string
	Validated bool // Indicates if the flags were successfully validated
}

type DeploymentMetadata struct {
	NumDeployments uint8
	Metadata       []map[string]uint8
}

type SelectedContent struct {
	ContinentCountryLocationMap map[string]map[string]map[string]map[string]DeploymentMetadata
	Validated                   bool // Indicates if the map is populated
}
