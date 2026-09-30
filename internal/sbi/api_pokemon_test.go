package sbi

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestGetPokemon(t *testing.T) {
	gin.SetMode(gin.TestMode)

	router := gin.New()

	s := &Server{}
	pokemonGroup := router.Group("/pokemon")
	applyRoutes(pokemonGroup, s.getPokemonRoute())

	req, _ := http.NewRequest(http.MethodGet, "/pokemon/Pikachu", nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", w.Code)
	}

	expected := `{"name":"Pikachu","type":"Electric"}`
	if strings.TrimSpace(w.Body.String()) != expected {
		t.Fatalf("expected %s, got %s", expected, w.Body.String())
	}
}

func TestCreatePokemon(t *testing.T) {
	gin.SetMode(gin.TestMode)

	router := gin.New()

	s := &Server{}
	pokemonGroup := router.Group("/pokemon")
	applyRoutes(pokemonGroup, s.getPokemonRoute())

	body := `{"name":"Eevee","type":"Normal"}`

	req, _ := http.NewRequest(
		http.MethodPost,
		"/pokemon/",
		strings.NewReader(body),
	)
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	if w.Code != http.StatusCreated {
		t.Fatalf("expected status 201, got %d", w.Code)
	}

	expected := `{"name":"Eevee","type":"Normal"}`
	if strings.TrimSpace(w.Body.String()) != expected {
		t.Fatalf("expected %s, got %s", expected, w.Body.String())
	}
}
