// Code generated by goctl. DO NOT EDIT.
package types

type SignUpRequest struct {
	Username string `json:"username,optional"`
	Email    string `json:"email,optional"`
	Password string `json:"password"`
}

type SignUpResponse struct {
	Token string `json:"token"`
}

type SignInRequest struct {
	Username string `json:"username,optional"`
	Email    string `json:"email,optional"`
	Password string `json:"password"`
}

type SignInResponse struct {
	Token string `json:"token"`
}

type User struct {
	Id       int64  `json:"id"`
	Username string `json:"username"`
	Avatar   string `json:"avatar"`
}

type Post struct {
	Id            int64  `json:"id,optional"`
	Title         string `json:"title,optional"`
	Content       string `json:"content,optional"`
	UserId        int64  `json:"user_id,optional"`
	CommentsCount int64  `json:"comments_count,optional"`
	CreatedAt     int64  `json:"created_at,optional"`
	UpdatedAt     int64  `json:"updated_at,optional"`
	User          User   `json:"user,optional"`
}

type Comment struct {
	Id        int64  `json:"id,optional"`
	Content   string `json:"content,optional"`
	PostId    int64  `json:"post_id,optional"`
	UserId    int64  `json:"user_id,optional"`
	CreatedAt int64  `json:"created_at,optional"`
	UpdatedAt int64  `json:"updated_at,optional"`
	Post      Post   `json:"post,optional"`
	User      User   `json:"user,optional"`
}

type CreatePostRequest struct {
	Post Post `json:"post"`
}

type CreatePostResponse struct {
	Post Post `json:"post"`
}

type DeletePostRequest struct {
	Id int64 `path:"id"`
}

type DeletePostResponse struct {
	Success bool `json:"success"`
}

type DeleteCommentRequest struct {
	Id int64 `path:"id"`
}

type DeleteCommentResponse struct {
	Success bool `json:"success"`
}

type GetPostRequest struct {
	Id int64 `path:"id"`
}

type GetPostResponse struct {
	Post Post `json:"post"`
}

type UpdatePostRequest struct {
	Post Post `json:"post"`
}

type UpdatePostResponse struct {
	Success bool `json:"success"`
}

type ListPostsRequest struct {
	Offset uint32 `form:"offset,optional"`
	Limit  uint32 `form:"limit"`
}

type ListPostsResponse struct {
	Post  []Post `json:"posts"`
	Total uint64 `json:"total"`
}

type ListCommentsByPostIDRequest struct {
	Offset uint32 `form:"offset,optional"`
	Limit  uint32 `form:"limit"`
	PostId uint64 `form:"postId"`
}

type ListCommentsByPostIDResponse struct {
	Comment []Comment `json:"comments"`
	Total   uint64    `json:"total"`
}

type CreateCommentRequest struct {
	Comment Comment `json:"comment"`
}

type CreateCommentResponse struct {
	Comment Comment `json:"comment"`
}

type UpdateCommentRequest struct {
	Comment Comment `json:"comment"`
}

type UpdateCommentResponse struct {
	Success bool `json:"success"`
}
