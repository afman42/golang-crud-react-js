package post

type Service interface {
	GetPosts() ([]Post, error)
	GetPostByID(input GetPostDetailInput) (Post, error)
	CreatePost(input CreatePostInput) (Post, error)
	UpdatePost(inputID GetPostDetailInput, inputData CreatePostInput) (Post, error)
	DeletePostByID(input GetPostDeleteInput) (Post, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) GetPosts() ([]Post, error) {
	posts, err := s.repository.FindAll()
	if err != nil {
		return posts, err
	}

	return posts, nil
}

func (s *service) GetPostByID(input GetPostDetailInput) (Post, error) {
	post, err := s.repository.FindByID(input.ID)

	if err != nil {
		return post, err
	}

	return post, nil
}

func (s *service) CreatePost(input CreatePostInput) (Post, error) {
	post := Post{}
	post.Title = input.Title
	post.Description = input.Description

	newPost, err := s.repository.Save(post)
	if err != nil {
		return newPost, err
	}

	return newPost, nil
}

func (s *service) UpdatePost(inputID GetPostDetailInput, inputData CreatePostInput) (Post, error) {
	post, err := s.repository.FindByID(inputID.ID)
	if err != nil {
		return post, err
	}

	post.Title = inputData.Title
	post.Description = inputData.Description

	updatedPost, err := s.repository.Update(post)
	if err != nil {
		return updatedPost, err
	}

	return updatedPost, nil
}

func (s *service) DeletePostByID(input GetPostDeleteInput) (Post, error) {
	post, err := s.repository.Delete(input.ID)

	if err != nil {
		return post, err
	}

	return post, nil
}
