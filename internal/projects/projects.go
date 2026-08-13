package projects

import "time"

type Project struct {
	ID          string    `db:"id"`
	Name        string    `db:"name"`
	Description string    `db:"description"`
	CreatedAt   time.Time `db:"created_at"`
	UpdatedAt   time.Time `db:"updated_at"`
}

type CreateProjectRequest struct {
	Name        string `json:"name" validate:"required"`
	Description string `json:"description" validate:"required"`
}

type CreateProjectResponse struct {
	Project Project `json:"project"`
}

type UpdateProjectRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}
