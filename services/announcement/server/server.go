package server

import (
	"context"

	"github.com/KISS-Keep-It-Simple-Stupid/TrekDestinyBackend/services/announcement/db"
	"github.com/KISS-Keep-It-Simple-Stupid/TrekDestinyBackend/services/announcement/pb"
)

type Repository struct {
	pb.UnimplementedAnnouncementServer
	DB db.Repository
}

func New(db db.Repository) *Repository {
	return &Repository{
		DB: db,
	}
}

func (s *Repository) CreateCard(context.Context, *pb.CreateCardRequest) (*pb.CreateCardResponse, error) {
	return nil, nil
}
