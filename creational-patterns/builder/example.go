package builder

import "fmt"

func BuilderExampleMain() {
	qb := NewQueryBuilder("Posts")

	// Get 10 most liked posts since the beginning of 2024
	qb.Select("user", "likes", "media", "publishedAt").Where("publishedAt >= '2024-01-01'").OrderBy("likes").Limit(10)

	query := qb.Execute()

	fmt.Println(query)
}
