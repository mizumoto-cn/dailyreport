package post

import (
	"context"
)

var (
	posters PosterFactoryMap = PosterFactoryMap{
		FactoryMap: make(map[string]PosterFactory),
		ConfMap:    make(map[string]any),
	}
)

func DefaultPosterFactoryMap() *PosterFactoryMap {
	return &posters
}

func Register(name string, factory PosterFactory, pc any) *PosterFactoryMap {
	posters.Register(name, factory, pc)
	return &posters
}

func CreatePoster(ctx context.Context, name string) (Poster, error) {
	return posters.CreatePoster(ctx, name)
}

func Post(ctx context.Context, name string, token ...string) error {
	p, err := posters.CreatePoster(ctx, name)
	if err != nil {
		return err
	}
	return p.Post(ctx, token...)
}
