package data

type LaunchFlags struct {
	Clusters  uint16
	RunMimik  bool
	Days      int
	Validated bool // Indicates if the flags were successfully validated
}

type SelectedContent struct {
	ContinentCountryLocationMap map[string]map[string][]string
	Validated                   bool // Indicates if the map is populated
}
