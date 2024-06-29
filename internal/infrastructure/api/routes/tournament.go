package routes

import (
	"net/http"

	"tournament_participation_service/internal/infrastructure/api/controllers"
	"tournament_participation_service/internal/infrastructure/api/views"
	"tournament_participation_service/internal/usecases"

	"github.com/labstack/echo/v4"
)

type tournamentHandler struct {
	uc *usecases.UseCases
}

// ParticipatedTournaments godoc
//
//	@Summary		Participated Tournaments
//	@Description	get all tournaments participated by authorized user
//	@Tags			tournaments
//	@Accept			json
//	@Produce		json
//	@Param			Authorization	header		string	true	"Authorization info"
//	@Success		200				{object}	views.ParticipatedTournamentsResponse
//	@Router			/user/tournaments [get]
func (h *tournamentHandler) ParticipatedTournaments(c echo.Context) error {
	ctx := c.Request().Context()

	userID := c.Request().Header.Get("Authorization")
	if userID == "" {
		return c.JSON(http.StatusUnauthorized, views.ErrorResponse{})
	}

	tournaments, err := h.uc.Queries.ParticipatedTournamentsHandler.Execute(ctx, userID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, views.ErrorResponse{})
	}

	tournamentViews := make([]*views.TournamentPreview, 0, len(tournaments))
	for _, tournament := range tournaments {
		tournamentViews = append(tournamentViews, &views.TournamentPreview{
			ID:    tournament.ID,
			Title: tournament.Title,
			Host: &views.Host{
				ID:       tournament.Host.ID,
				UserID:   tournament.Host.UserID,
				Username: tournament.Host.Username,
			},
			Date:         tournament.Date,
			TotalPlayers: tournament.TotalPlayers,
		})
	}

	resp := &views.ParticipatedTournamentsResponse{Tournaments: tournamentViews}

	return c.JSON(http.StatusOK, resp)
}

// GetTournaments godoc
//
//	@Summary		Get Tournaments
//	@Description	get all tournaments
//	@Tags			tournaments
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	views.GetTournamentsResponse
//	@Router			/tournaments [get]
func (h *tournamentHandler) GetTournaments(c echo.Context) error {
	ctx := c.Request().Context()

	tournaments, err := h.uc.Queries.AllTournamentHandler.Execute(ctx)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, views.ErrorResponse{})
	}

	tournamentViews := make([]*views.TournamentPreview, 0, len(tournaments))
	for _, tournament := range tournaments {
		tournamentViews = append(tournamentViews, &views.TournamentPreview{
			ID:    tournament.ID,
			Title: tournament.Title,
			Host: &views.Host{
				ID:       tournament.Host.ID,
				UserID:   tournament.Host.UserID,
				Username: tournament.Host.Username,
			},
			Date:         tournament.Date,
			TotalPlayers: tournament.TotalPlayers,
		})
	}

	resp := &views.GetTournamentsResponse{Tournaments: tournamentViews}

	return c.JSON(http.StatusOK, resp)
}

// GetTournament godoc
//
//	@Summary		Get Tournament
//	@Description	get tournament
//	@Tags			tournaments
//	@Accept			json
//	@Produce		json
//	@Param			Authorization	header		string	true	"Authorization info"
//	@Param			id				path		string	true	"ID of the tournament"
//	@Success		200				{object}	views.GetTournamentResponse
//	@Router			/tournaments/{id} [get]
func (h *tournamentHandler) GetTournament(c echo.Context) error {
	ctx := c.Request().Context()

	userID := c.Request().Header.Get("Authorization")

	var req controllers.GetTournamentRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, views.ErrorResponse{})
	}

	tournament, err := h.uc.Queries.TournamentByIDHandler.Execute(ctx, userID, req.ID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, views.ErrorResponse{})
	}

	tournamentView := &views.Tournament{
		ID:    tournament.ID,
		Title: tournament.Title,
		Host: &views.Host{
			ID:       tournament.Host.ID,
			UserID:   tournament.Host.UserID,
			Username: tournament.Host.Username,
		},
		Date:           tournament.Date,
		TotalPlayers:   tournament.TotalPlayers,
		PlayerEnrolled: tournament.PlayerEnrolled,
	}

	resp := &views.GetTournamentResponse{Tournament: tournamentView}

	return c.JSON(http.StatusOK, resp)
}

// GetPlayers godoc
//
//	@Summary		Get Players
//	@Description	get players for tournament
//	@Tags			tournaments,players
//	@Accept			json
//	@Produce		json
//	@Param			Authorization	header		string	true	"Authorization info"
//	@Param			id				path		string	true	"ID of the tournament"
//	@Success		200				{object}	views.GetPlayersResponse
//	@Router			/tournaments/{id}/players [get]
func (h *tournamentHandler) GetPlayers(c echo.Context) error {
	ctx := c.Request().Context()

	userID := c.Request().Header.Get("Authorization")
	if userID == "" {
		return c.JSON(http.StatusUnauthorized, views.ErrorResponse{})
	}

	var req controllers.GetPlayersRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, views.ErrorResponse{})
	}

	enrolledPlayers, err := h.uc.Queries.EnrolledPlayersHandler.Execute(ctx, userID, req.TournamentID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, views.ErrorResponse{})
	}

	playerViews := make([]*views.Player, 0, len(enrolledPlayers))
	for _, enrolledPlayer := range enrolledPlayers {
		playerViews = append(playerViews, &views.Player{
			ID:       enrolledPlayer.ID,
			UserID:   enrolledPlayer.UserID,
			Username: enrolledPlayer.Username,
			Dropped:  enrolledPlayer.Dropped,
		})
	}

	resp := &views.GetPlayersResponse{Players: playerViews}

	return c.JSON(http.StatusOK, resp)
}

func makeTournamentRoutes(e *echo.Group, uc *usecases.UseCases) {
	h := tournamentHandler{uc: uc}

	{
		e := e.Group("/user/tournaments")
		e.GET("", h.ParticipatedTournaments)
	}

	{
		e := e.Group("/tournaments")
		e.GET("", h.GetTournaments)

		{
			e := e.Group("/:id")
			e.GET("", h.GetTournament)

			{
				e := e.Group("/players")
				e.GET("", h.GetPlayers)
			}
		}
	}
}
