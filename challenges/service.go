package challenges

import (
	"errors"
	"student-app/internal/models"
	"time"
)

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) CreateChallenge(req *CreateChallengeRequest, userID int) (*models.Challenge, error) {
	var startDate, endDate time.Time
	var err error

	if req.StartDate != "" {
		startDate, err = time.Parse("2006-01-02", req.StartDate)
		if err != nil {
			return nil, errors.New("неверный формат даты начала")
		}
	} else {
		startDate = time.Now()
	}

	if req.EndDate != "" {
		endDate, err = time.Parse("2006-01-02", req.EndDate)
		if err != nil {
			return nil, errors.New("неверный формат даты окончания")
		}
	} else {
		endDate = startDate.AddDate(0, 0, 7)
	}

	challenge := &models.Challenge{
		Title:       req.Title,
		Description: req.Description,
		Type:        req.Type,
		StartDate:   startDate,
		EndDate:     endDate,
		PrizeXP:     req.PrizeXP,
		PrizeCoins:  req.PrizeCoins,
		CreatedBy:   userID,
	}

	if challenge.Type == "" {
		challenge.Type = "tasks"
	}

	if err := s.repo.CreateChallenge(challenge); err != nil {
		return nil, err
	}
	return challenge, nil
}

func (s *Service) UpdateChallenge(id int, req *UpdateChallengeRequest) error {
	updates := make(map[string]interface{})

	if req.Title != "" {
		updates["title"] = req.Title
	}
	if req.Description != "" {
		updates["description"] = req.Description
	}
	if req.Type != "" {
		updates["type"] = req.Type
	}
	if req.StartDate != "" {
		updates["start_date"] = req.StartDate
	}
	if req.EndDate != "" {
		updates["end_date"] = req.EndDate
	}
	if req.PrizeXP > 0 {
		updates["prize_xp"] = req.PrizeXP
	}
	if req.PrizeCoins > 0 {
		updates["prize_coins"] = req.PrizeCoins
	}
	if req.Status != "" {
		updates["status"] = req.Status
	}

	return s.repo.UpdateChallenge(id, updates)
}

func (s *Service) DeleteChallenge(id int) error {
	return s.repo.DeleteChallenge(id)
}

func (s *Service) GetChallengeByID(id int, userID int) (*ChallengeResponse, error) {
	ch, err := s.repo.FindChallengeByID(id)
	if err != nil {
		return nil, err
	}

	participantsCount, _ := s.repo.GetParticipantsCount(id)
	myScore, _ := s.repo.GetParticipantScore(id, userID)
	myRank, _ := s.repo.GetParticipantRank(id, userID)

	return &ChallengeResponse{
		ID:                ch.ID,
		Title:             ch.Title,
		Description:       ch.Description,
		Type:              ch.Type,
		StartDate:         ch.StartDate,
		EndDate:           ch.EndDate,
		PrizeXP:           ch.PrizeXP,
		PrizeCoins:        ch.PrizeCoins,
		ParticipantsCount: participantsCount,
		MyScore:           myScore,
		MyRank:            myRank,
		CreatedBy:         ch.CreatedBy,
		CreatedAt:         ch.CreatedAt,
	}, nil
}

func (s *Service) GetAllChallenges(userID int, query *GetChallengesQuery) ([]ChallengeResponse, error) {
	limit := query.Limit
	if limit == 0 {
		limit = 20
	}

	challenges, err := s.repo.GetAllChallenges(query.Type, query.Status, limit, query.Offset)
	if err != nil {
		return nil, err
	}

	var response []ChallengeResponse
	for _, ch := range challenges {
		participantsCount, _ := s.repo.GetParticipantsCount(ch.ID)
		myScore, _ := s.repo.GetParticipantScore(ch.ID, userID)

		response = append(response, ChallengeResponse{
			ID:                ch.ID,
			Title:             ch.Title,
			Description:       ch.Description,
			Type:              ch.Type,
			StartDate:         ch.StartDate,
			EndDate:           ch.EndDate,
			PrizeXP:           ch.PrizeXP,
			PrizeCoins:        ch.PrizeCoins,
			ParticipantsCount: participantsCount,
			MyScore:           myScore,
			CreatedBy:         ch.CreatedBy,
			CreatedAt:         ch.CreatedAt,
		})
	}
	return response, nil
}

func (s *Service) JoinChallenge(challengeID, userID int) error {
	ch, err := s.repo.FindChallengeByID(challengeID)
	if err != nil {
		return errors.New("вызов не найден")
	}

	now := time.Now()
	if now.Before(ch.StartDate) {
		return errors.New("вызов еще не начался")
	}
	if now.After(ch.EndDate) {
		return errors.New("вызов уже завершен")
	}

	isParticipant, err := s.repo.IsParticipant(challengeID, userID)
	if err != nil {
		return err
	}
	if isParticipant {
		return errors.New("вы уже участвуете в вызове")
	}

	return s.repo.JoinChallenge(challengeID, userID)
}

func (s *Service) LeaveChallenge(challengeID, userID int) error {
	return s.repo.LeaveChallenge(challengeID, userID)
}

func (s *Service) UpdateScore(challengeID, userID, score int, adminID int) error {
	// Проверяем права (только создатель вызова или админ)
	ch, err := s.repo.FindChallengeByID(challengeID)
	if err != nil {
		return err
	}
	if ch.CreatedBy != adminID {
		return errors.New("только организатор может изменять оценки")
	}

	return s.repo.UpdateScore(challengeID, userID, score)
}

func (s *Service) GetLeaderboard(challengeID int, userID int) (*LeaderboardResponse, error) {
	participants, err := s.repo.GetLeaderboard(challengeID)
	if err != nil {
		return nil, err
	}

	ch, err := s.repo.FindChallengeByID(challengeID)
	if err != nil {
		return nil, err
	}

	var response []ParticipantResponse
	for i, p := range participants {
		// Получаем имя пользователя
		var userName string
		s.repo.db.QueryRow("SELECT name FROM users WHERE id = $1", p.UserID).Scan(&userName)

		response = append(response, ParticipantResponse{
			UserID:   p.UserID,
			UserName: userName,
			Score:    p.Score,
			Rank:     i + 1,
			JoinedAt: p.JoinedAt,
		})
	}

	return &LeaderboardResponse{
		ChallengeID:    challengeID,
		ChallengeTitle: ch.Title,
		Participants:   response,
		TotalCount:     len(response),
	}, nil
}

func (s *Service) AwardPrizes(challengeID int, adminID int) error {
	ch, err := s.repo.FindChallengeByID(challengeID)
	if err != nil {
		return err
	}
	if ch.CreatedBy != adminID {
		return errors.New("только организатор может награждать")
	}

	return s.repo.AwardPrizes(challengeID)
}
