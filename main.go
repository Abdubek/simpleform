package app

import (
	"app/app/usecase"
	"app/ports/http"
)

func main() {
	usecase := usecase.NewUseCase()
	http.NewServer(":8080", usecase).Run()
}
