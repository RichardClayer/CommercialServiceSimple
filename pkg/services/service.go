package services

import "database/sql"

type Service struct {
    db *sql.DB
}

func NewService(db *sql.DB) *Service {
    return &Service{
        db: db,
    }
}

func (s *Service) connect() {

}
