package service

import (
	"context"
	"fmt"
	"github.com/gookit/slog"
	"sync"
)

func (s *Service) PreloadCache(ctx context.Context) error {
	preloadFns := []func(context.Context) error{
		s.PreloadOrdersCache,
	}

	errChan := make(chan error, len(preloadFns))
	var wg sync.WaitGroup

	for _, fn := range preloadFns {
		wg.Add(1)
		go func(f func(context.Context) error) {
			defer wg.Done()
			if err := f(ctx); err != nil {
				errChan <- err
			}
		}(fn)
	}

	wg.Wait()
	close(errChan)

	var errs []error
	for err := range errChan {
		errs = append(errs, err)
	}

	if len(errs) > 0 {
		return fmt.Errorf("cache preload errors: %v", errs)
	}
	slog.Info("Orders cache successfully preloaded")

	return nil
}

func (s *Service) PreloadOrdersCache(ctx context.Context) error {
	data, err := s.repo.GetAllOrders(ctx)
	if err != nil {
		slog.Errorf("failed to load orders for cache: %v", err)
		return err
	}

	for _, order := range data {
		cacheKey := fmt.Sprintf("order:%s", order.OrderUID.String())

		if err := s.redis.SetJSON(cacheKey, order, 0); err != nil {
			slog.Errorf("failed to set orders cache: %v", err)
			return err
		}
	}

	return nil
}
