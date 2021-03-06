package dvdrental

type Film struct {
	Title       string
	Description string
	Actors      []Actor
}

type Actor struct {
	FirstName string
	LastName  string
}

var FILMS = [...]Film{
	{
		Title:       "ACADEMY DINOSAUR",
		Description: "A Epic Drama of a Feminist And a Mad Scientist who must Battle a Teacher in The Canadian Rockies",
		Actors: []Actor{
			{FirstName: "CHRISTIAN", LastName: "GABLE"},
			{FirstName: "JOHNNY", LastName: "CAGE"},
			{FirstName: "LUCILLE", LastName: "TRACY"},
			{FirstName: "MARY", LastName: "KEITEL"},
			{FirstName: "MENA", LastName: "TEMPLE"},
			{FirstName: "OPRAH", LastName: "KILMER"},
			{FirstName: "PENELOPE", LastName: "GUINESS"},
			{FirstName: "ROCK", LastName: "DUKAKIS"},
			{FirstName: "SANDRA", LastName: "PECK"},
			{FirstName: "WARREN", LastName: "NOLTE"},
		},
	},
	{
		Title:       "ACE GOLDFINGER",
		Description: "A Astounding Epistle of a Database Administrator And a Explorer who must Find a Car in Ancient China",
		Actors: []Actor{
			{FirstName: "BOB", LastName: "FAWCETT"},
			{FirstName: "CHRIS", LastName: "DEPP"},
			{FirstName: "MINNIE", LastName: "ZELLWEGER"},
			{FirstName: "SEAN", LastName: "GUINESS"},
		},
	},
	{
		Title:       "ADAPTATION HOLES",
		Description: "A Astounding Reflection of a Lumberjack And a Car who must Sink a Lumberjack in A Baloon Factory",
		Actors: []Actor{
			{FirstName: "BOB", LastName: "FAWCETT"},
			{FirstName: "CAMERON", LastName: "STREEP"},
			{FirstName: "JULIANNE", LastName: "DENCH"},
			{FirstName: "NICK", LastName: "WAHLBERG"},
			{FirstName: "RAY", LastName: "JOHANSSON"},
		},
	},
	{
		Title:       "AFFAIR PREJUDICE",
		Description: "A Fanciful Documentary of a Frisbee And a Lumberjack who must Chase a Monkey in A Shark Tank",
		Actors: []Actor{
			{FirstName: "FAY", LastName: "WINSLET"},
			{FirstName: "JODIE", LastName: "DEGENERES"},
			{FirstName: "KENNETH", LastName: "PESCI"},
			{FirstName: "OPRAH", LastName: "KILMER"},
			{FirstName: "SCARLETT", LastName: "DAMON"},
		},
	},
	{
		Title:       "AFRICAN EGG",
		Description: "A Fast-Paced Documentary of a Pastry Chef And a Dentist who must Pursue a Forensic Psychologist in The Gulf of Mexico",
		Actors: []Actor{
			{FirstName: "DUSTIN", LastName: "TAUTOU"},
			{FirstName: "GARY", LastName: "PHOENIX"},
			{FirstName: "MATTHEW", LastName: "CARREY"},
			{FirstName: "MATTHEW", LastName: "LEIGH"},
			{FirstName: "THORA", LastName: "TEMPLE"},
		},
	},
	{
		Title:       "AGENT TRUMAN",
		Description: "A Intrepid Panorama of a Robot And a Boy who must Escape a Sumo Wrestler in Ancient China",
		Actors: []Actor{
			{FirstName: "JAYNE", LastName: "NESON"},
			{FirstName: "KENNETH", LastName: "HOFFMAN"},
			{FirstName: "KIRSTEN", LastName: "PALTROW"},
			{FirstName: "MORGAN", LastName: "WILLIAMS"},
			{FirstName: "REESE", LastName: "WEST"},
			{FirstName: "SANDRA", LastName: "KILMER"},
			{FirstName: "WARREN", LastName: "NOLTE"},
		},
	},
	{
		Title:       "AIRPLANE SIERRA",
		Description: "A Touching Saga of a Hunter And a Butler who must Discover a Butler in A Jet Boat",
		Actors: []Actor{
			{FirstName: "JIM", LastName: "MOSTEL"},
			{FirstName: "MENA", LastName: "HOPPER"},
			{FirstName: "MICHAEL", LastName: "BOLGER"},
			{FirstName: "OPRAH", LastName: "KILMER"},
			{FirstName: "RICHARD", LastName: "PENN"},
		},
	},
	{
		Title:       "AIRPORT POLLOCK",
		Description: "A Epic Tale of a Moose And a Girl who must Confront a Monkey in Ancient India",
		Actors: []Actor{
			{FirstName: "FAY", LastName: "KILMER"},
			{FirstName: "GENE", LastName: "WILLIS"},
			{FirstName: "LUCILLE", LastName: "DEE"},
			{FirstName: "SUSAN", LastName: "DAVIS"},
		},
	},
	{
		Title:       "ALABAMA DEVIL",
		Description: "A Thoughtful Panorama of a Database Administrator And a Mad Scientist who must Outgun a Mad Scientist in A Jet Boat",
		Actors: []Actor{
			{FirstName: "CHRISTIAN", LastName: "GABLE"},
			{FirstName: "ELVIS", LastName: "MARX"},
			{FirstName: "GRETA", LastName: "KEITEL"},
			{FirstName: "MENA", LastName: "TEMPLE"},
			{FirstName: "MERYL", LastName: "ALLEN"},
			{FirstName: "RIP", LastName: "CRAWFORD"},
			{FirstName: "RIP", LastName: "WINSLET"},
			{FirstName: "WARREN", LastName: "NOLTE"},
			{FirstName: "WILLIAM", LastName: "HACKMAN"},
		},
	},
	{
		Title:       "ALADDIN CALENDAR",
		Description: "A Action-Packed Tale of a Man And a Lumberjack who must Reach a Feminist in Ancient China",
		Actors: []Actor{
			{FirstName: "ALEC", LastName: "WAYNE"},
			{FirstName: "GRETA", LastName: "MALDEN"},
			{FirstName: "JADA", LastName: "RYDER"},
			{FirstName: "JUDY", LastName: "DEAN"},
			{FirstName: "RAY", LastName: "JOHANSSON"},
			{FirstName: "RENEE", LastName: "TRACY"},
			{FirstName: "ROCK", LastName: "DUKAKIS"},
			{FirstName: "VAL", LastName: "BOLGER"},
		},
	},
}

func getFilms(n ...int) []Film {
	films := make([]Film, 0, len(n))
	for _, i := range n {
		films = append(films, FILMS[i])
	}
	return films
}

func getFilmPtrs(n ...int) []*Film{
	films := make([]*Film, 0, len(n))
	for _, i := range n {
		film := FILMS[i]
		films = append(films, &film)
	}
	return films
}

func setA() []Film {
	films := getFilms(0, 1, 2, 3, 4, 5, 6, 7, 8, 9)
	films[4].Description = "A Documentary of a Pastry Chef And a Dentist who must Pursue a Forensic Psychologist in The Gulf of Mexico"
	films[5].Actors = []Actor{
		{FirstName: "KIRSTEN", LastName: "PALTROW"},
		{FirstName: "MORGAN", LastName: "WILLIAMS"},
		{FirstName: "REESE", LastName: "WEST"},
		{FirstName: "LUCILLE", LastName: "DEE"},
		{FirstName: "SUSAN", LastName: "DAVIS"},
		{FirstName: "SANDRA", LastName: "KILMER"},
		{FirstName: "WARREN", LastName: "NOLTE"},
	}
	return films
}

func setB() []Film {
	films := getFilms(0, 1, 2, 3, 4, 5, 6, 7, 8, 9)
	return films
}
