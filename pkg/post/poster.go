package post

import "context"

type Poster interface {
	Post(ctx context.Context, token ...string) error
}

type PosterFactory interface {
	CreatePoster(context.Context, any) (Poster, error)
}

type PosterFactoryFunc func(ctx context.Context, pc any) (Poster, error)

func (f PosterFactoryFunc) CreatePoster(ctx context.Context, pc any) (Poster, error) {
	return f(ctx, pc)
}

type PosterFactoryMap struct {
	FactoryMap map[string]PosterFactory
	ConfMap    map[string]any
}

func (m PosterFactoryMap) Register(name string, factory PosterFactory, pc any) {
	m.FactoryMap[name] = factory
	m.ConfMap[name] = pc
}

func (m PosterFactoryMap) CreatePoster(ctx context.Context, name string) (Poster, error) {
	factory, ok := m.FactoryMap[name]
	if !ok {
		return nil, ErrUnknownPoster
	}
	return factory.CreatePoster(ctx, m.ConfMap[name])
}
