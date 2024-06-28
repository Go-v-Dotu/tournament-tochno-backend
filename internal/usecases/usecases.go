package usecases

import (
	"github.com/ThreeDotsLabs/watermill/components/cqrs"
	"tournament_participation_service/internal/domain"
	"tournament_participation_service/internal/usecases/commands"
	"tournament_participation_service/internal/usecases/queries"
)

type UseCases struct {
	Commands Commands
	Queries  Queries
}

type Commands struct {
	CreateHostHandler       *commands.CreateHostHandler
	CreatePlayerHandler     *commands.CreatePlayerHandler
	CreateTournamentHandler *commands.CreateTournamentHandler
}

type Queries struct {
	AllTournamentHandler           *queries.AllTournamentHandler
	EnrolledPlayersHandler         *queries.EnrolledPlayersHandler
	ParticipatedTournamentsHandler *queries.ParticipatedTournamentsHandler
	TournamentByIDHandler          *queries.TournamentByIDHandler
}

func NewUseCases(
	eventBus *cqrs.EventBus,
	hostRepo domain.HostRepository,
	playerRepo domain.PlayerRepository,
	tournamentRepo domain.TournamentRepository,
	tournamentQueryService queries.TournamentQueryService,
	playerQueryService queries.PlayerQueryService,
) *UseCases {
	return &UseCases{
		Commands: Commands{
			CreateHostHandler:       commands.NewCreateHostHandler(hostRepo),
			CreatePlayerHandler:     commands.NewCreatePlayerHandler(playerRepo),
			CreateTournamentHandler: commands.NewCreateTournamentHandler(tournamentRepo),
		},
		Queries: Queries{
			AllTournamentHandler:           queries.NewAllTournamentHandler(tournamentQueryService),
			EnrolledPlayersHandler:         queries.NewEnrolledPlayersHandler(playerRepo, tournamentRepo, playerQueryService),
			ParticipatedTournamentsHandler: queries.NewParticipatedTournamentsHandler(playerRepo, tournamentQueryService),
			TournamentByIDHandler:          queries.NewTournamentByIDHandler(playerRepo, tournamentRepo, tournamentQueryService),
		},
	}
}
