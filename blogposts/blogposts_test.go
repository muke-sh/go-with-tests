package blogposts_test

import (
	"go-with-tests/blogposts"
	"testing"
	"testing/fstest"
)

func TestNewBlogPosts(t *testing.T) {
	fs := fstest.MapFS{
		"hello world.md":  {Data: []byte("hi")},
		"hello world2.md": {Data: []byte("hola")},
	}

	posts := blogposts.NewPostFromFS(fs)

	if len(posts) != len(fs) {
		t.Errorf("got %d posts, but wanted %d posts", len(posts), len(fs))
	}
}
