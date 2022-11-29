package blogposts

import "testing/fstest"

type Post struct {
	Title, Description, Body string
	Tags                     []string
}

func NewPostFromFS(fileSystem fstest.MapFS) []Post {
	return []Post{{}, {}}
}
