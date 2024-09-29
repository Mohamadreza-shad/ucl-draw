package draw

import (
	"context"
	"errors"
	"fmt"

	"github.com/Mohamadreza-shad/ucl-draw/client"
	"github.com/Mohamadreza-shad/ucl-draw/repository"
)

type Service struct {
	db   client.PgxInterface
	repo *repository.Queries
}

type Club struct {
	Id          int64
	Name        string
	Nationality string
	Seed        int32
}

type Match struct {
	HomeClubId int
	AwayClubId int
}

type MatchName struct {
	MatchId int64
	Home    string
	Away    string
}

func (s *Service) Draw(ctx context.Context) error {
	var matches []Match
	clubs, err := s.repo.GetAllClubs(ctx, s.db)
	if err != nil {
		return errors.New("something went wrong")
	}
	seeds := make(map[int32][]Club)
	for _, c := range clubs {
		seeds[c.Seed] = append(
			seeds[c.Seed],
			Club{
				Id:          c.ID,
				Name:        c.Name,
				Nationality: c.Nationality,
				Seed:        c.Seed,
			})
	}

	for _, c := range clubs {
		clubMatches := 0
		homeGames := 0
		awayGames := 0
		opponents := make(map[int]bool)

		for seedLevel := 1; seedLevel <= 4; seedLevel++ {
			for _, op := range seeds[int32(seedLevel)] {
				if c.ID == op.Id || c.Nationality == op.Nationality || opponents[int(op.Id)] {
					continue
				}
				if clubMatches >= 8 || homeGames == 4 || awayGames == 4 {
					break
				}
				if homeGames < 4 {
					matches = append(matches, Match{HomeClubId: int(c.ID), AwayClubId: int(op.Id)})
					homeGames++
				} else {
					matches = append(matches, Match{HomeClubId: int(op.Id), AwayClubId: int(c.ID)})
					awayGames++
				}
				clubMatches++
				opponents[int(op.Id)] = true
			}
		}
		if clubMatches != 8 {
			return fmt.Errorf("unable to generate 8 matches for club: %s", c.Name)
		}
	}
	for _, m := range matches {
		_, err := s.repo.CreateMatch(
			ctx,
			s.db,
			repository.CreateMatchParams{
				HostID:  int32(m.HomeClubId),
				GuestID: int32(m.AwayClubId),
			})
		if err != nil {
			return errors.New("something went wrong")
		}
	}
	return nil
}

func (s *Service) DrawResult(ctx context.Context) ([]MatchName, error) {
	matches, err := s.repo.GetMatches(ctx, s.db)
	if err != nil {
		return nil, errors.New("something went wrong")
	}
	matchNames := make([]MatchName, len(matches))
	for i, m := range matches {
		matchNames[i] = MatchName{
			Home:    m.HostName,
			Away:    m.AwayName,
			MatchId: m.ID,
		}
	}
	return matchNames, nil
}

func NewService(
	db client.PgxInterface,
	repo *repository.Queries,
) *Service {
	return &Service{
		db:   db,
		repo: repo,
	}
}
