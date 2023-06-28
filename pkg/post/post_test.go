package post_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/mizumoto-cn/dailyreport/pkg/post"
)

func TestPost(t *testing.T) {
	ctx := context.Background()

	// Register a mock poster factory
	post.Register("mock", post.PosterFactoryFunc(func(ctx context.Context, pc any) (post.Poster, error) {
		return &mockPoster{}, nil
	}), nil)

	// Test posting with a valid poster name
	err := post.Post(ctx, "mock")
	assert.NoError(t, err)

	// Test posting with an invalid poster name
	err = post.Post(ctx, "invalid")
	assert.Error(t, err)
}

type mockPoster struct{}

func (p *mockPoster) Post(ctx context.Context, token ...string) error {
	return nil
}

var _ post.Poster = (*mockPoster)(nil)

// END: yz3j7f5d9x1p
