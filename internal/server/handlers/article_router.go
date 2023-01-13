package handlers

import (
	"encoding/json"
	"fmt"
	"go-postgres/pkg/article"
	"go-postgres/pkg/response"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
)

type ArticleRouter struct {
	Repository article.Repository
}

/*
	CREATE ARTICLE VIDEOGAME
*/

func (art *ArticleRouter) CreateHandler(w http.ResponseWriter, r *http.Request) {
	var articleRepo article.ArticleGame

	err := json.NewDecoder(r.Body).Decode(&articleRepo)

	if err != nil {
		response.HTTPError(w, r, http.StatusBadRequest, err.Error())
	}

	defer r.Body.Close()

	ctx := r.Context()
	err = art.Repository.Create(ctx, &articleRepo)

	if err != nil {
		response.HTTPError(w, r, http.StatusBadRequest, err.Error())
	}

	w.Header().Add("Location", fmt.Sprintf("%s%d", r.URL.String(), articleRepo.ID))
	response.JSON(w, r, http.StatusCreated, response.Map{"Article VideoGame": articleRepo})

}

/*
	LIST Articles

*/

func (art *ArticleRouter) GetAllHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	articles, err := art.Repository.GetAll(ctx)
	if err != nil {
		response.HTTPError(w, r, http.StatusNotFound, err.Error())
		return
	}

	response.JSON(w, r, http.StatusOK, articles)
}

/*
	DELETE ARTICLES
*/

func (art *ArticleRouter) DeleteHandler(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		response.HTTPError(w, r, http.StatusBadRequest, err.Error())
		return
	}

	ctx := r.Context()
	err = art.Repository.Delete(ctx, uint(id))
	if err != nil {
		response.HTTPError(w, r, http.StatusNotFound, err.Error())
		return
	}

	response.JSON(w, r, http.StatusOK, response.Map{"msg": "Delete Article!"})
}

/*
	GET BY USER
*/

func (art *ArticleRouter) GetByUserHandler(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "idUser")
	ctx := r.Context()

	//Convertir str a integer
	id_user, err := strconv.Atoi(idStr)
	if err != nil {
		response.HTTPError(w, r, http.StatusNotFound, err.Error())
		return
	}

	//uint significa entero sin signo, es decir siempre un entero positivo
	articles, err := art.Repository.GetByUser(ctx, uint(id_user))
	if err != nil {
		response.HTTPError(w, r, http.StatusNotFound, err.Error())
		return
	}

	response.JSON(w, r, http.StatusOK, response.Map{"ArticlesVideoGames": articles})

}

/*
	GET ONE ARTICLES
*/

func (art *ArticleRouter) GetOneHandler(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		response.HTTPError(w, r, http.StatusBadRequest, err.Error())
		return
	}

	ctx := r.Context()
	article, err := art.Repository.GetOne(ctx, uint(id))
	if err != nil {
		response.HTTPError(w, r, http.StatusNotFound, err.Error())
		return
	}

	response.JSON(w, r, http.StatusOK, response.Map{"ArticleVideoGame": article})
}

/*
	UPDATE ARTICLES
*/

func (art *ArticleRouter) UpdateHandler(w http.ResponseWriter, r *http.Request) {

	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		response.HTTPError(w, r, http.StatusBadRequest, err.Error())
		return
	}

	var articleRepo article.ArticleGame
	err = json.NewDecoder(r.Body).Decode(&articleRepo)

	if err != nil {
		response.HTTPError(w, r, http.StatusBadRequest, err.Error())
		//GOlang regresa los parametros de la funcion por defecto en el return, creo.
		return
	}

	defer r.Body.Close()

	ctx := r.Context()
	article, err := art.Repository.Update(ctx, uint(id), articleRepo)
	if err != nil {
		response.HTTPError(w, r, http.StatusNotFound, err.Error())
		return
	}

	//response.JSON(w, r, http.StatusOK, nil)
	response.JSON(w, r, http.StatusOK, response.Map{"Article VideoGame": article})

}

//ROUTERS

func (art *ArticleRouter) Routes() http.Handler {
	r := chi.NewRouter()

	r.Get("/", art.GetAllHandler)  //LIST
	r.Post("/", art.CreateHandler) //CREATE

	r.Get("/{id}/", art.GetOneHandler) //DETAIL
	r.Put("/{id}/", art.UpdateHandler) //UPDATE

	r.Delete("/{id}/", art.DeleteHandler)          //DELETE
	r.Get("/user/{idUser}/", art.GetByUserHandler) //LIST ARTICLES BY ID USER
	return r
}
