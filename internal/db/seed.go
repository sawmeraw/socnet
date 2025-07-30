package db

// import (
// 	"context"
// 	"fmt"
// 	"log"
// 	"math/rand"

// 	"github.com/sawmeraw/gogo/internal/store"
// )

// var usernames = []string{
// 	"sunnyPixel", "astroWhale", "codeBandit", "mistyMount", "chocoWolf",
// 	"electricMango", "radishNinja", "aquaRider", "byteDreamer", "frostEcho",
// 	"moonTaco", "lavaSloth", "iceCricket", "zipFalcon", "cloudTofu",
// 	"happyCactus", "slimOtter", "ghostLobster", "plasmaBug", "turboKoala",
// }

// var emails = []string{
// 	"sunnyPixel99@example.com", "astro.whale@example.org", "codeBandit42@example.net",
// 	"mistyMount88@example.com", "choco_wolf123@example.org",
// 	"electricMango77@example.net", "radishNinja@example.com", "aquaRider22@example.org",
// 	"dreamer.byte@example.net", "frostEcho91@example.com",
// 	"moonTaco@example.org", "lavaSloth35@example.net", "iceCricket08@example.com",
// 	"zip.falcon@example.org", "cloudTofu@example.net",
// 	"happy.cactus@example.com", "slimOtter14@example.org", "ghost.lobster@example.net",
// 	"plasmaBug@example.com", "turboKoala99@example.org",
// }

// var titles = []string{
// 	"Exploring the Edge of the Universe", "How to Bake the Perfect Sourdough",
// 	"Why Otters Are the Cutest Animals", "Top 10 VSCode Extensions in 2025",
// 	"A Deep Dive into Quantum Computing", "The Joy of Early Morning Runs",
// 	"Mastering the Art of Minimalism", "The Rise of Electric Dirt Bikes",
// 	"When AI Meets Classical Music", "Rainy Days and Productivity Boosts",
// 	"How I Built My Backyard Observatory", "Coding a Game in Go — Step by Step",
// 	"Tokyo’s Hidden Food Streets", "The Psychology of Color in Design",
// 	"From Zero to Kubernetes Hero", "The Best Aussie Beaches No One Talks About",
// 	"Designing with Accessibility in Mind", "Understanding Crypto Tax Laws",
// 	"A Vegan’s Guide to BBQ Season", "The Magic of Soundscapes in Film",
// }

// var contents = []string{
// 	"Ever wondered what lies beyond our galaxy? Let’s take a journey into deep space theories.",
// 	"A step-by-step guide to making crusty, tangy sourdough using only three ingredients.",
// 	"Otters hold hands while sleeping — and we’re here to prove they’re nature’s cuddle champions.",
// 	"Boost your dev workflow with these killer VSCode plugins and power-user shortcuts.",
// 	"Quantum computers promise exponential speed. But how do they actually work?",
// 	"Running at sunrise isn't just healthy, it's therapeutic — here’s my 30-day morning run diary.",
// 	"Decluttering isn’t just a trend. It’s a lifestyle shift toward intentional living.",
// 	"Why rugged electric bikes are the next big thing for adrenaline junkies and eco-warriors.",
// 	"I built an AI model that generates Bach-style fugues. Here’s what happened.",
// 	"Rainy days often get a bad rap — here’s how I doubled my productivity indoors.",
// 	"This post chronicles my DIY observatory build. Spoiler: The roof retracts!",
// 	"Want to make your own game using Go? Here's how I did it in under a week.",
// 	"Beyond sushi — a walk through Tokyo’s street eats you’ll wish you knew sooner.",
// 	"Colors aren’t just pretty — they steer user emotions and UX decisions. Here's how.",
// 	"New to Kubernetes? I was too. This is my roadmap from confused to confident.",
// 	"Hidden beach gems in South Australia you won’t believe exist — until you visit.",
// 	"Designing apps isn’t just about looks — it’s about usability for everyone.",
// 	"Got crypto gains? Here’s what Aussie law says about reporting and deducting.",
// 	"BBQ season is back — here's how vegans can party with smoky plant-based goodness.",
// 	"Sound design transforms film — let me show you the power of auditory storytelling.",
// }

// var tags = [][]string{
// 	{"space", "astronomy", "science"}, {"baking", "food", "homemade"},
// 	{"animals", "otters", "cute"}, {"coding", "VSCode", "developer"},
// 	{"quantum", "tech", "innovation"}, {"running", "fitness", "mindfulness"},
// 	{"minimalism", "lifestyle", "decluttering"}, {"ebikes", "adventure", "tech"},
// 	{"AI", "music", "creativity"}, {"productivity", "motivation", "lifestyle"},
// 	{"DIY", "astronomy", "build"}, {"gameDev", "golang", "tutorial"},
// 	{"Tokyo", "food", "travel"}, {"design", "psychology", "UX"},
// 	{"Kubernetes", "devOps", "cloud"}, {"beach", "travel", "australia"},
// 	{"accessibility", "design", "UX"}, {"crypto", "finance", "australia"},
// 	{"vegan", "BBQ", "recipes"}, {"film", "sound", "creativity"},
// }

// var commentsText = []string{
// 	"Great post!",
// 	"Really insightful, thanks!",
// 	"I learned something new today.",
// 	"Couldn’t agree more.",
// 	"This is so helpful!",
// 	"Wow, well said.",
// 	"Love the perspective here.",
// 	"Such a good read!",
// 	"Interesting take on the topic.",
// 	"Appreciate you sharing this.",
// 	"Nice breakdown.",
// 	"This gave me something to think about.",
// 	"Thanks for explaining it so clearly.",
// 	"Brilliantly put.",
// 	"I'm bookmarking this!",
// 	"Exactly what I needed.",
// 	"So true!",
// 	"Absolutely!",
// 	"Well articulated.",
// 	"Keep it coming!",
// }

// func Seed(store store.Storage) {
// 	ctx := context.Background()

// 	users := generateUsers(100)

// 	for _, user := range users {
// 		if err := store.Users.Create(ctx, user); err != nil {
// 			log.Println("Error creating user: ", err)
// 			return
// 		}
// 	}

// 	posts := generatePosts(100, users)

// 	for _, post := range posts {
// 		if err := store.Posts.Create(ctx, post); err != nil {
// 			log.Println("Error creating post: ", err)
// 			return
// 		}
// 	}

// 	comments := generateComments(100, users, posts)
// 	for _, comment := range comments {

// 		if err := store.Comments.Create(ctx, comment); err != nil {
// 			log.Println("Error creating comment: ", err)
// 			return
// 		}
// 	}

// 	log.Println("databaes seeding complete.")
// }

// func generateComments(n int, users []*store.User, posts []*store.Post) []*store.Comment {
// 	comments := make([]*store.Comment, n)

// 	for i := 0; i < n; i++ {
// 		user := users[rand.Intn(len(users))]
// 		post := posts[rand.Intn(len(posts))]
// 		comment := commentsText[i%len(commentsText)] + fmt.Sprintf("%d", i)
// 		comments[i] = &store.Comment{
// 			UserID:  user.ID,
// 			PostID:  post.ID,
// 			Content: comment,
// 		}
// 	}

// 	return comments
// }

// func generatePosts(n int, users []*store.User) []*store.Post {
// 	posts := make([]*store.Post, n)

// 	for i := 0; i < n; i++ {
// 		user := users[rand.Intn(len(users))]

// 		posts[i] = &store.Post{
// 			UserID:  user.ID,
// 			Title:   titles[rand.Intn(len(titles))],
// 			Content: contents[rand.Intn(len(contents))],
// 			Tags:    tags[rand.Intn(len(tags))],
// 		}
// 	}

// 	return posts
// }

// func generateUsers(n int) []*store.User {

// 	users := make([]*store.User, n)

// 	for i := 0; i < n; i++ {
// 		users[i] = &store.User{
// 			Username: usernames[i%len(usernames)] + fmt.Sprintf("%d", i),
// 			Email:    emails[i%len(emails)] + fmt.Sprintf("%d", i),
// 			// Password: "123123",
// 		}
// 	}

// 	return users
// }
