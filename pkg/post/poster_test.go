package post_test

import (
	"context"
	"errors"
	"testing"

	. "github.com/mizumoto-cn/dailyreport/pkg/post"
)

// type mockPoster struct{}

// func (p *mockPoster) Post(ctx context.Context, token ...string) error {
// 	return nil
// }

type mockPosterFactory struct{}

func (f *mockPosterFactory) CreatePoster(ctx context.Context, pc any) (Poster, error) {
	return &mockPoster{}, nil
}

func TestPosterFactoryMap_CreatePoster(t *testing.T) {
	factoryMap := PosterFactoryMap{
		FactoryMap: make(map[string]PosterFactory),
		ConfMap:    make(map[string]any),
	}
	factoryMap.Register("mock", &mockPosterFactory{}, nil)

	t.Run("known poster", func(t *testing.T) {
		poster, err := factoryMap.CreatePoster(context.Background(), "mock")
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
		if _, ok := poster.(*mockPoster); !ok {
			t.Errorf("expected a mockPoster, but got %T", poster)
		}
	})

	t.Run("unknown poster", func(t *testing.T) {
		_, err := factoryMap.CreatePoster(context.Background(), "unknown")
		if !errors.Is(err, ErrUnknownPoster) {
			t.Errorf("expected ErrUnknownPoster, but got %v", err)
		}
	})
}
