package post

type PostFormatter struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

func FormatPost(post Post) PostFormatter {
	postFormatter := PostFormatter{}
	postFormatter.ID = post.ID
	postFormatter.Title = post.Title
	postFormatter.Description = post.Description

	return postFormatter
}
func FormatPosts(posts []Post) []PostFormatter {
	postsFormatter := []PostFormatter{}

	for _, post := range posts {
		postFormatter := FormatPost(post)
		postsFormatter = append(postsFormatter, postFormatter)
	}

	return postsFormatter
}
