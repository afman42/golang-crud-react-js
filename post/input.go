package post

type GetPostDetailInput struct {
	ID int `uri:"id" binding:"required"`
}

type GetPostDeleteInput struct {
	ID int `uri:"id" binding:"required"`
}

type CreatePostInput struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
}

type FormUpdateCampaignInput struct {
	ID          int
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
	Error       error
}
