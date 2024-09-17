package route

import (
	"net/http"

	"github.com/TimeGladiator/TimeGladiator/internal/controller"
)

func RegisterRoutes() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/v1/healthcheck", controller.HealthcheckHandler)
	return mux
}
