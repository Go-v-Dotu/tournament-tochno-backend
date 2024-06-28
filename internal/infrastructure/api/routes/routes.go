package routes

import (
	"tournament_participation_service/internal/application"

	"github.com/labstack/echo/v4"
)

func Make(e *echo.Group, app *application.App) {
	makeTournamentRoutes(e, app.UseCases)
}
