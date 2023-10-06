package utils

const (
	DefaultClusters = 10
	DefaultFaux     = false
	DefaultDays     = 30

	MinClusters = 2
	MaxClusters = 198

	MinDays = 1
	MaxDays = 365

	LogErrorPrefix = "[ERR] "

	NameCount   = 53
	PlaceCount  = 98
	AnimalCount = 54
	ThingCount  = 52
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
)
