package games

import (
	"errors"
	"student-app/internal/models"
)

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{repo: repo}
}

// ========== Управление играми (админ) ==========
func (s *Service) CreateGame(req *CreateGameRequest) (*models.MiniGame, error) {
	game := &models.MiniGame{
		Title:       req.Title,
		Description: req.Description,
		Icon:        req.Icon,
		XPReward:    req.XPReward,
		CoinReward:  req.CoinReward,
		GameData:    req.GameData,
	}

	if game.Icon == "" {
		game.Icon = "🎮"
	}
	if game.XPReward == 0 {
		game.XPReward = 10
	}

	if err := s.repo.CreateGame(game); err != nil {
		return nil, err
	}
	return game, nil
}

func (s *Service) UpdateGame(id int, req *UpdateGameRequest) error {
	updates := make(map[string]interface{})

	if req.Title != "" {
		updates["title"] = req.Title
	}
	if req.Description != "" {
		updates["description"] = req.Description
	}
	if req.Icon != "" {
		updates["icon"] = req.Icon
	}
	if req.XPReward > 0 {
		updates["xp_reward"] = req.XPReward
	}
	if req.CoinReward > 0 {
		updates["coin_reward"] = req.CoinReward
	}
	if req.GameData != "" {
		updates["game_data"] = req.GameData
	}
	updates["is_active"] = req.IsActive

	return s.repo.UpdateGame(id, updates)
}

func (s *Service) DeleteGame(id int) error {
	return s.repo.DeleteGame(id)
}

func (s *Service) GetGameByID(id int) (*GameResponse, error) {
	game, err := s.repo.FindGameByID(id)
	if err != nil {
		return nil, err
	}

	playsCount, avgScore, _ := s.repo.GetGameStats(id)

	return &GameResponse{
		ID:          game.ID,
		Title:       game.Title,
		Description: game.Description,
		Icon:        game.Icon,
		XPReward:    game.XPReward,
		CoinReward:  game.CoinReward,
		GameData:    game.GameData,
		PlaysCount:  playsCount,
		AvgScore:    avgScore,
		CreatedAt:   game.CreatedAt,
	}, nil
}

func (s *Service) GetAllGames(query *GetGamesQuery) ([]GameResponse, error) {
	limit := query.Limit
	if limit == 0 {
		limit = 50
	}

	games, err := s.repo.GetAllGames(query.IsActive, limit, query.Offset)
	if err != nil {
		return nil, err
	}

	var response []GameResponse
	for _, game := range games {
		playsCount, avgScore, _ := s.repo.GetGameStats(game.ID)

		response = append(response, GameResponse{
			ID:          game.ID,
			Title:       game.Title,
			Description: game.Description,
			Icon:        game.Icon,
			XPReward:    game.XPReward,
			CoinReward:  game.CoinReward,
			PlaysCount:  playsCount,
			AvgScore:    avgScore,
			CreatedAt:   game.CreatedAt,
		})
	}
	return response, nil
}

// ========== Игровая механика для пользователей ==========
func (s *Service) SubmitGameResult(userID int, req *SubmitGameResultRequest) (*GameResultResponse, error) {
	// Получаем информацию об игре
	game, err := s.repo.FindGameByID(req.GameID)
	if err != nil {
		return nil, errors.New("игра не найдена")
	}

	// Сохраняем результат
	result := &models.GameResult{
		UserID:      userID,
		GameID:      req.GameID,
		Score:       req.Score,
		XPEarned:    game.XPReward,
		CoinsEarned: game.CoinReward,
	}

	if err := s.repo.SaveGameResult(result); err != nil {
		return nil, err
	}

	// Начисляем награду
	if game.XPReward > 0 {
		s.repo.AddXP(userID, game.XPReward)
	}
	if game.CoinReward > 0 {
		s.repo.AddCoins(userID, game.CoinReward)
	}

	// Получаем лучший результат пользователя
	//bestScore, _ := s.repo.GetUserBestScore(userID, req.GameID)

	return &GameResultResponse{
		ID:          result.ID,
		GameID:      result.GameID,
		GameTitle:   game.Title,
		Score:       result.Score,
		XPEarned:    result.XPEarned,
		CoinsEarned: result.CoinsEarned,
		PlayedAt:    result.PlayedAt,
	}, nil
}

func (s *Service) GetUserResults(userID, gameID int, limit, offset int) ([]GameResultResponse, error) {
	results, err := s.repo.GetUserGameResults(userID, gameID, limit, offset)
	if err != nil {
		return nil, err
	}

	var response []GameResultResponse
	for _, r := range results {
		// Получаем название игры
		game, _ := s.repo.FindGameByID(r.GameID)
		gameTitle := ""
		if game != nil {
			gameTitle = game.Title
		}

		response = append(response, GameResultResponse{
			ID:          r.ID,
			GameID:      r.GameID,
			GameTitle:   gameTitle,
			Score:       r.Score,
			XPEarned:    r.XPEarned,
			CoinsEarned: r.CoinsEarned,
			PlayedAt:    r.PlayedAt,
		})
	}
	return response, nil
}

func (s *Service) GetLeaderboard(gameID int, limit int) ([]GameLeaderboardEntry, error) {
	if limit == 0 {
		limit = 10
	}
	return s.repo.GetGameLeaderboard(gameID, limit)
}

func (s *Service) GetUserBestScore(userID, gameID int) (int, error) {
	return s.repo.GetUserBestScore(userID, gameID)
}
