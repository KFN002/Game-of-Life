package handler

import (
	"context"
	"encoding/json"
	"life_game/internal/service"
	"net/http"
)

type Decorator func(http.Handler) http.Handler

type LifeStates struct {
	service.LifeService
}

func New(ctx context.Context,
	lifeService service.LifeService,
) (http.Handler, error) {
	serveMux := http.NewServeMux()
	lifeState := LifeStates{
		LifeService: lifeService,
	}
	serveMux.HandleFunc("/nextstate", lifeState.nextState)
	return serveMux, nil
}

func Decorate(next http.Handler, ds ...Decorator) http.Handler {
	decorated := next
	for d := len(ds) - 1; d >= 0; d-- {
		decorated = ds[d](decorated)
	}
	return decorated
}

func (ls *LifeStates) nextState(w http.ResponseWriter, r *http.Request) {
	worldState := ls.LifeService.NewState()

	err := json.NewEncoder(w).Encode(worldState.Cells)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
