package main

var demoOrgs = []string{
	"prequels",
	"originals",
}

type demoUser struct {
	id        string
	email     string
	firstName string
	lastName  string
	auth0ID   string
}

var demoUsers = []demoUser{
	demoUser{
		firstName: "qui-gon",
		lastName:  "jinn",
	},
	demoUser{
		firstName: "obi-wan",
		lastName:  "kenobi",
	},
	demoUser{
		firstName: "anakin",
		lastName:  "skywalker",
	},
	demoUser{
		firstName: "padme",
		lastName:  "amidala",
	},
	demoUser{
		firstName: "C",
		lastName:  "3P0",
	},
	demoUser{
		firstName: "R2",
		lastName:  "D2",
	},
	demoUser{
		firstName: "sheev",
		lastName:  "palpatine",
	},
	demoUser{
		firstName: "darth",
		lastName:  "maul",
	},
	demoUser{
		firstName: "count",
		lastName:  "dooku",
	},
	demoUser{
		firstName: "luke",
		lastName:  "skywalker",
	},
	demoUser{
		firstName: "leia",
		lastName:  "organa",
	},
	demoUser{
		firstName: "han",
		lastName:  "solo",
	},
	demoUser{
		firstName: "chewbacca",
		lastName:  "",
	},
}

var titles = []string{
	"The Phantom Menace",
	"Attack of the Clones",
	"Revenge of the Sith",
	"A New Hope",
	"The Empire Strikes Back",
	"Return of the Jedi",
}

var bodies = map[int][]string{
	0: []string{
		"Turmoil has engulfed the Galactic Republic. The taxation of trade routes to outlying star systems is in dispute.",
		"Hoping to resolve the matter with a blockade of deadly battleships, the greedy Trade Federation has stopped all shipping to the small planet of Naboo.",
		"While the congress of the Republic endlessly debates this alarming chain of events, the Supreme Chancellor has secretly dispatched two Jedi Knights, the guardians of peace and justice in the galaxy, to settle the conflict....",
	},
	1: []string{
		"There is unrest in the Galactic Senate. Several thousand solar systems have declared their intentions to leave the Republic.",
		"This Separatist movement, under the leadership of the mysterious Count Dooku, has made it difficult for the limited number of Jedi Knights to maintain peace and order in the galaxy.",
		"Senator Amidala, the former Queen of Naboo, is returning to the Galactic Senate to vote on the critical issue of creating an ARMY OF THE REPUBLIC to assist the overwhelmed Jedi....",
	},
	2: []string{
		"War! The Republic is crumbling under attacks by the ruthless Sith Lord, Count Dooku. There are heroes on both sides. Evil is everywhere.",
		"In a stunning move, the fiendish droid leader, General Grievous, has swept into the Republic capital and kidnapped Chancellor Palpatine, leader of the Galactic Senate.",
		"As the Separatist Droid Army attempts to flee the besieged capital with their valuable hostage, two Jedi Knights lead a desperate mission to rescue the captive Chancellor....",
	},
	3: []string{
		"It is a period of civil war. Rebel spaceships, striking from a hidden base, have won their first victory against the evil Galactic Empire.",
		"During the battle, Rebel spies managed to steal secret plans to the Empire's ultimate weapon, the DEATH STAR, an armored space station with enough power to destroy an entire planet.",
		"Pursued by the Empire's sinister agents, Princess Leia races home aboard her starship, custodian of the stolen plans that can save her people and restore freedom to the galaxy....",
	},
	4: []string{
		"It is a dark time for the Rebellion. Although the Death Star has been destroyed, Imperial troops have driven the Rebel forces from their hidden base and pursued them across the galaxy.",
		"Evading the dreaded Imperial Starfleet, a group of freedom fighters led by Luke Skywalker have established a new secret base on the remote ice world of Hoth.",
		"The evil lord Darth Vader, obsessed with finding young Skywalker, has dispatched thousands of remote probes into the far reaches of space....",
	},
	5: []string{
		"Luke Skywalker has returned to his home planet of Tatooine in an attempt to rescue his friend Han Solo from the clutches of the vile gangster Jabba the Hutt.",
		"Little does Luke know that the GALACTIC EMPIRE has secretly begun construction on a new armored space station even more powerful than the first dreaded Death Star.",
		"When completed, this ultimate weapon will spell certain doom for the small band of rebels struggling to restore freedom to the galaxy...",
	},
}

var quotes = []string{
	"Oh, I have a bad feeling about this.",
	"Hello there!",
	"I know.",
	"You will be a Jedi. I promise.",
	"Begun the clone war has.",
	"May the Force be with you.",
	"Henceforth you shall be known as...Darth Vader.",
	"Did you hear that? They've shut down the main reactor!",
	"Use the Force, Luke.",
	"Tell Jabba I've got his money.",
	"Echo 3 to Echo 7 - Han ol' buddy, do you read me?",
}

var sexes = []string{
	"MALE",
	"FEMALE",
}

var races = []string{
	"AMERICAN_INDIAN_OR_ALASKA_NATIVE",
	"ASIAN",
	"BLACK_OR_AFRICAN_AMERICAN",
	"HISPANIC_OR_LATINO",
	"WHITE",
}

var specimenTypes = []string{
	"BLOOD",
}

var containerTypes = []string{
	"VIAL",
}

var specimenStatuses = []string{
	"DESTROYED",
	"EXHAUSTED",
	"IN_INVENTORY",
	"IN_TRANSIT",
	"LOST",
	"RESERVED",
	"TRANSFERRED",
}

var bloodTypes = []string{
	"O_NEG",
	"O_POS",
	"A_NEG",
	"A_POS",
	"B_NEG",
	"B_POS",
	"AB_NEG",
	"AB_POS",
}
