package projects

import (
	"context"
)

type ColumnService struct {
	columnRepo ColumnRepository
}

func NewColumnService(repo ColumnRepository) *ColumnService {
	return &ColumnService{columnRepo: repo}
}

func (s *ColumnService) CreateColumn(ctx context.Context, c *Column) (Column, error) {
	return s.columnRepo.Create(ctx, c)
}
