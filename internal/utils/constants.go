package utils

const (
	DefaultClusters = 10
	DefaultFaux     = false
	DefaultDays     = 30

	MinClusters = 1
	MaxClusters = 180

	MinDays = 1
	MaxDays = 365

	LogErrorPrefix = "[ERR] "

	NameCount   = 53
	PlaceCount  = 98
	AnimalCount = 54
	ThingCount  = 52

	DefaultDirName = "mimik-data"

	MinDeployments = 5
	MaxDeployments = 15
	MinContainers  = 1
	MaxContainers  = 2
	MinReplicas    = 1
	MaxReplicas    = 5

	LowModeMinCPU        float64 = 0.001
	LowModeMaxCPU        float64 = 0.3
	LowModeMinMemory     uint64  = 75000000
	LowModemaxMemory     uint64  = 150000000
	LowModeCPURequest    float64 = 0.5
	LowModeCPULimit      float64 = 1
	LowModeMemoryRequest uint64  = 150000000
	LowModeMemoryLimit   uint64  = 150000000

	AvgModeMinCPU        float64 = 0.3
	AvgModeMaxCPU        float64 = 1.5
	AvgModeMinMemory     uint64  = 100000000
	AvgModemaxMemory     uint64  = 300000000
	AvgModeCPURequest    float64 = 1
	AvgModeCPULimit      float64 = 2
	AvgModeMemoryRequest uint64  = 100000000
	AvgModeMemoryLimit   uint64  = 300000000

	HighModeMinCPU        float64 = 0.8
	HighModeMaxCPU        float64 = 3
	HighModeMinMemory     uint64  = 300000000
	HighModemaxMemory     uint64  = 1500000000
	HighModeCPURequest    float64 = 2
	HighModeCPULimit      float64 = 4
	HighModeMemoryRequest uint64  = 1500000000
	HighModeMemoryLimit   uint64  = 1500000000

	VeryHighModeMinCPU        float64 = 1.5
	VeryHighModeMaxCPU        float64 = 6
	VeryHighModeMinMemory     uint64  = 1500000000
	VeryHighModemaxMemory     uint64  = 4000000000
	VeryHighModeCPURequest    float64 = 4
	VeryHighModeCPULimit      float64 = 4
	VeryHighModeMemoryRequest uint64  = 4000000000
	VeryHighModeMemoryLimit   uint64  = 4000000000

	Query_CPU_Request    string = "kube_pod_container_resource_requests_cpu_cores"
	Query_CPU_Limit      string = "kube_pod_container_resource_limits_cpu_cores"
	Query_Memory_Request string = "kube_pod_container_resource_requests_memory_bytes"
	Query_Memory_limit   string = "kube_pod_container_resource_limits_memory_bytes"
	Query_CPU_Usage      string = "container_cpu_usage_seconds_total"
	Query_Memory_Usage   string = "container_memory_usage_bytes"
)

var (
	// Names list
	Names = [NameCount]string{
		"abigail", "alexander", "alice", "amelia", "andrew", "anthony", "aria", "ava", "benjamin", "brooklyn",
		"caleb", "carter", "charles", "charlotte", "chloe", "christopher", "claire", "daniel", "elijah", "elizabeth",
		"ella", "emma", "ethan", "eva", "grace", "hannah", "harper", "henry", "isabella", "jacob",
		"jackson", "james", "john", "joseph", "julian", "liam", "lily", "lucas", "lucy", "mason",
		"matthew", "mia", "michael", "nathaniel", "nicholas", "noah", "oliver", "olivia", "samuel", "sarah",
		"sophia", "william", "zoe",
	}

	// Places list
	Places = [PlaceCount]string{
		"abilene", "albuquerque", "alexandria", "amarillo", "anchorage", "ann-arbor", "arlington", "atlanta", "austin", "baltimore",
		"billings", "boston", "boulder", "bridgeport", "broken-arrow", "burbank", "cambridge", "carlsbad", "carrollton", "cary",
		"cedar-rapids", "centennial", "charleston", "charlotte", "chicago", "clearwater", "clovis", "college-station", "colorado-springs", "columbia",
		"columbus", "coral-springs", "corona", "corpus-christi", "dallas", "davenport", "denton", "denver", "detroit", "downey",
		"duluth", "edison", "edmond", "el-monte", "el-paso", "elgin", "evansville", "fairfield", "fargo", "fayetteville",
		"flint", "fontana", "fort-collins", "fort-lauderdale", "fort-worth", "frisco", "fullerton", "gainesville", "garden-grove", "garland",
		"gilbert", "glendale", "green-bay", "greensboro", "houston", "indianapolis", "jacksonville", "kansas-city", "las-vegas", "long-beach",
		"los-angeles", "louisville", "memphis", "miami", "milwaukee", "minneapolis", "nashville", "new-orleans", "new-york", "oakland",
		"oklahoma-city", "omaha", "philadelphia", "phoenix", "pittsburgh", "portland", "raleigh", "sacramento", "san-antonio", "san-diego",
		"san-francisco", "san-jose", "seattle", "tucson", "tulsa", "virginia-beach", "washington", "wichita",
	}

	// Animals list (100 unique animals)
	Animals = [AnimalCount]string{
		"ant", "bat", "bear", "bee", "bird", "camel", "cat", "cheetah", "chicken", "chimpanzee",
		"cow", "crocodile", "deer", "dog", "dolphin", "duck", "elephant", "fish", "fox", "frog",
		"giraffe", "goat", "gorilla", "hippopotamus", "horse", "kangaroo", "koala", "lion", "lizard", "monkey",
		"octopus", "otter", "owl", "panda", "parrot", "penguin", "pig", "rabbit", "raccoon", "rhinoceros",
		"scorpion", "seal", "shark", "sheep", "snail", "snake", "spider", "squirrel", "tiger", "turtle",
		"walrus", "whale", "wolf", "zebra",
	}

	// Things list
	Things = [ThingCount]string{
		"apple", "ball", "book", "bottle", "box", "car", "chair", "computer", "cup", "desk",
		"door", "flower", "guitar", "hat", "house", "key", "lamp", "phone", "picture", "shoes",
		"table", "tree", "umbrella", "watch", "window", "pen", "pencil", "shirt", "shoe", "sock",
		"bag", "camera", "clock", "hat", "headphones", "jacket", "magazine", "mirror", "newspaper", "pants",
		"pillow", "ring", "scissors", "soap", "spoon", "table", "toothbrush", "toothpaste", "towel", "wallet",
		"water-bottle", "wine-glass",
	}

	// Modes list
	ResourceUsageMode = [4]string{
		"low",
		"avg",
		"high",
		"very_high",
	}

	// Resource Queries list
	Queries = map[string]string{
		Query_CPU_Request:    "gauge",
		Query_CPU_Limit:      "gauge",
		Query_Memory_Request: "gauge",
		Query_Memory_limit:   "gauge",
		Query_CPU_Usage:      "counter",
		Query_Memory_Usage:   "gauge",
	}
)
