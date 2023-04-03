package server

import v1 "github.com/tiptok/gz-blog-microsevices/api/protobuf/blog/v1"

var prefix = "/" + v1.BlogService_ServiceDesc.ServiceName + "/"

var AuthMethods = map[string]bool{
	prefix + "SignUp":               false, // 不需要验证
	prefix + "SignIn":               false,
	prefix + "CreatePost":           true, // 需要验证
	prefix + "UpdatePost":           true,
	prefix + "GetPost":              false,
	prefix + "ListPosts":            false,
	prefix + "DeletePost":           true,
	prefix + "CreateComment":        true,
	prefix + "UpdateComment":        true,
	prefix + "DeleteComment":        true,
	prefix + "ListCommentsByPostID": false,
}

var httpPrefix = "/api.rest.blog.v1" + "/"

var AuthHttpMethods = map[string]bool{
	prefix + "sign-up":              false, // 不需要验证
	prefix + "sign-in":              false,
	prefix + "posts":                true, // 需要验证
	prefix + "posts/{id}":           true,
	prefix + "GetPost":              false,
	prefix + "ListPosts":            false,
	prefix + "DeletePost":           true,
	prefix + "CreateComment":        true,
	prefix + "UpdateComment":        true,
	prefix + "DeleteComment":        true,
	prefix + "ListCommentsByPostID": false,
}
