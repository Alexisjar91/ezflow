package projects

import (
	"context"
)

type ProjectService struct {
	projectRepo ProjectRepository
}

func NewProjectService(repo ProjectRepository) *ProjectService {
	return &ProjectService{projectRepo: repo}
}

func (s *ProjectService) CreateProject(ctx context.Context, p *Project) (Project, error) {
	return s.projectRepo.Create(ctx, p)
}
