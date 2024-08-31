package helperts

import (
	"blog-rest/internal/models"
	"sort"
	"time"
)

func ParseTime(timeStr string) time.Time {
	parsedTime, _ := time.Parse(time.RFC3339, timeStr)
	return parsedTime
}

func SortByPosts(users []models.User, posts []models.Post, order string) []models.User {
	postCount := make(map[int]int)

	for _, post := range posts {
		postCount[post.UserID]++
	}

	if order == "asc" {
		sort.Slice(users, func(i, j int) bool {
			return postCount[users[i].ID] < postCount[users[j].ID]
		})
	} else {
		sort.Slice(users, func(i, j int) bool {
			return postCount[users[i].ID] > postCount[users[j].ID]
		})
	}

	return users
}
