package main

import (
	"time"

	"github.com/ps/entity"
)

var (
	nextCommentId  = 100
	nextBlogPostId = 100
)

var authors = map[string]*entity.Author{
	"jcleese": &entity.Author{
		ID:        1,
		FirstName: "John",
		LastName:  "Cleese",
		Username:  "jcleese",
	},
	"tjones": &entity.Author{
		ID:        2,
		FirstName: "Terry",
		LastName:  "Jones",
		Username:  "tjones",
	},
	"eidle": &entity.Author{
		ID:        3,
		FirstName: "Eric",
		LastName:  "Idle",
		Username:  "eidle",
	},
	"tgilliam": &entity.Author{
		ID:        4,
		FirstName: "Terry",
		LastName:  "Gilliam",
		Username:  "tgilliam",
	},
	"gchapman": &entity.Author{
		ID:        5,
		FirstName: "Graham",
		LastName:  "Chapman",
		Username:  "gchapman",
	},
	"mpalin": &entity.Author{
		ID:        6,
		FirstName: "Michael",
		LastName:  "Palin",
		Username:  "mpalin",
	},
}

var blogPosts = []entity.BlogPost{
	entity.BlogPost{
		entity.ContentItem{
			ID:          1,
			Subject:     "From Killer Rabbits to African Swallows - How to Find an Unusual Pet",
			Body:        "",
			Author:      authors["jcleese"],
			Comments:    []entity.Comment{},
			CreatedDate: makeTime(2014, time.January, 14),
			PublishDate: makeTime(2014, time.January, 16),
			IsPublished: true,
		}},
	entity.BlogPost{
		entity.ContentItem{
			ID:          2,
			Subject:     "Ducks and Witches - Why Physics Matters",
			Body:        "",
			Author:      authors["tjones"],
			Comments:    []entity.Comment{},
			CreatedDate: makeTime(2014, time.April, 2),
			PublishDate: makeTime(2014, time.April, 8),
			IsPublished: true,
		}},
	entity.BlogPost{
		entity.ContentItem{
			ID:          3,
			Subject:     "Why I Think Sir Robin is the Bravest Man I Know",
			Body:        "",
			Author:      authors["eidle"],
			Comments:    []entity.Comment{},
			CreatedDate: makeTime(2015, time.February, 18),
			PublishDate: makeTime(2015, time.February, 27),
			IsPublished: true,
		}},
	entity.BlogPost{
		entity.ContentItem{
			ID:          4,
			Author:      authors["tgilliam"],
			Comments:    []entity.Comment{},
			CreatedDate: makeTime(2015, time.March, 14),
			PublishDate: makeTime(2015, time.March, 16),
			IsPublished: true,
		}},
	entity.BlogPost{
		entity.ContentItem{
			ID:          5,
			Subject:     "Why I Took the Limb Dismemberment Class",
			Body:        "",
			Author:      authors["mpalin"],
			Comments:    []entity.Comment{},
			CreatedDate: makeTime(2015, time.April, 5),
			PublishDate: makeTime(2015, time.April, 8),
			IsPublished: true,
		}},
	entity.BlogPost{
		entity.ContentItem{
			ID:          6,
			Subject:     "Swamp Castle - A Study in Modern Architecure Gone Wrong",
			Body:        "",
			Author:      authors["jcleese"],
			Comments:    []entity.Comment{},
			CreatedDate: makeTime(2015, time.May, 7),
			PublishDate: makeTime(2015, time.May, 12),
			IsPublished: true,
		}},
	entity.BlogPost{
		entity.ContentItem{
			ID:      7,
			Subject: "What Its Like to be the Only Survivor of 'Hiding in a Field'",
			Body: `It seemed like such an easy gig. I just had to stand (or crouch)
			where they told me in the middle of a field and wait. Then the explosions started.
			I've been a practicing field-stander for 8 years and I've never seen anything like it, one
			minute a fellow field-stander was peacefully going about his job and then - BOOM, gone.
			Apparently, we had stepped into some kind of anti-guerilla training that the government
			was doing, but no one bothered to tell us that they'd be using live explosives. It was a
			rough day to say the least. I don't sleep much anymore and the have a hard time holding down a job...`,
			Author:      authors["tjones"],
			Comments:    []entity.Comment{},
			CreatedDate: makeTime(2015, time.June, 16),
			PublishDate: makeTime(2015, time.June, 19),
			IsPublished: true,
		}},
	entity.BlogPost{
		entity.ContentItem{
			ID:      8,
			Subject: "Finding my Inner Silly Walk",
			Body: `I first learned about the Ministry of Silly Walks when I was
			6 years old. I was fascinated by the graceful means that the elegant ladies
			and gentlemen of the ministry used as their personal means of locomotion.
			I immediately set my young mind to task of perfecting my own style with hopes
			of eventually joining the elite at the ministry. I was never interested in sports or
			video games like other kids; for me, it was all about my walk. As I grew older
			I joined the Silly Walk Club at school and that's when I started to realize that
			I was born with a rare disease that started when I entered my teen years. By
			the time I was 15, I couldn't so much as a forward aerial-half turn. My walk was
			perfectly normal, and my dreams were shattered.`,
			Author:      authors["eidle"],
			Comments:    []entity.Comment{},
			CreatedDate: makeTime(2015, time.August, 1),
			PublishDate: makeTime(2015, time.August, 16),
			IsPublished: true,
		}},
	entity.BlogPost{
		entity.ContentItem{
			ID:      9,
			Subject: "Buying a Dead Parrot Saved My Marriage",
			Body: `We started out so happy. My wife and I married young, but we
			knew that we were right for each other from the beginning. As I started
			to build my career, I started to have to work later hours and travel and
			it didn't take long for us to feel the strain. Things were headed for a bad end
			until I got Polly. Many people fixate on the fact that our parrot is dead, and I
			guess I can understand that, but I can't describe the calming influence that she
			has on our little home. She just sits there quietly nailed to her perch and
			casts her gaze across our small apartment, never troubled by a thing. She reminds
			both of us that there really is nothing to worry about. We've even started to talk
			about beginning a family. We're confident that we can do it, and we think that Polly
			will be great with children.`,
			Author:      authors["tgilliam"],
			Comments:    []entity.Comment{},
			CreatedDate: makeTime(2015, time.October, 27),
			PublishDate: makeTime(2015, time.October, 31),
			IsPublished: true,
		}},
	entity.BlogPost{
		entity.ContentItem{
			ID:      10,
			Subject: "A Brief History of Spam",
			Body: `Thousands of years before there were lumber jacks, men
			were still men and needed to prove themselves as such. Mostly they did
			this by chasing large animals around with pointed sticks until the poor
			creatures gave up the ghost and lept off of a cliff. The over-excited
			hunters would then pluck all of the useful bits off of the poor beast
			and then sit around for days gorging themselves on their prize. Over time,
			mankind grew and began to civilize itself by farming. On the whole, this
			was a good thing and let to such innovations as animal domestication,
			writing, and an overall improvement in the human condition. Unfortunately for
			those amonst mankind with powerful hunting urges, the game was chased off, pointed
			sticks were increasingly discouraged and farms generally did better on flat ground where
			there were few, if any conveniently placed cliffs for prey to leap from. In order
			to satisfy their lust for the raw pleasure of rough, barely palatable meat, one of the
			greatest hunter of the farming age invented Spam, and life has never been quite the same.`,
			Author:      authors["mpalin"],
			Comments:    []entity.Comment{},
			CreatedDate: makeTime(2015, time.November, 7),
			PublishDate: makeTime(2015, time.December, 8),
			IsPublished: true,
		}},
}

func makeTime(year int, month time.Month, day int) *time.Time {
	result := time.Date(year, month, day, 0, 0, 0, 0, time.UTC)

	return &result
}
