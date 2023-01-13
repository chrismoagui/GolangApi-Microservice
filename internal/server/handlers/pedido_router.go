package handlers

import (
	"encoding/json"
	"fmt"
	"go-postgres/pkg/pedido"
	"go-postgres/pkg/response"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
)

type PedidoRouter struct {
	Repository pedido.Repository
}

func (a *PedidoRouter) CreateHandler(w http.ResponseWriter, r *http.Request) {
	var pedidoRepo pedido.PedidoGame

	err := json.NewDecoder(r.Body).Decode(&pedidoRepo)

	if err != nil {
		response.HTTPError(w, r, http.StatusBadRequest, err.Error())
	}

	defer r.Body.Close()

	ctx := r.Context()
	err = a.Repository.Create(ctx, &pedidoRepo)

	if err != nil {
		response.HTTPError(w, r, http.StatusBadRequest, err.Error())
	}

	w.Header().Add("Location", fmt.Sprintf("%s%d", r.URL.String(), pedidoRepo.ID))
	response.JSON(w, r, http.StatusCreated, response.Map{"pedidos VideoGame": pedidoRepo})

}

func (a *PedidoRouter) GetAllHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	pedidos, err := a.Repository.GetAll(ctx)
	if err != nil {
		response.HTTPError(w, r, http.StatusNotFound, err.Error())
		return
	}

	response.JSON(w, r, http.StatusOK, pedidos)
}

func (art *PedidoRouter) DeleteHandler(w http.ResponseWriter, r *http.Request) {
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

	response.JSON(w, r, http.StatusOK, response.Map{"msg": "Delete pedido!"})
}

/*
	GET BY USER
*/

func (a *PedidoRouter) GetByUserHandler(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "idUser")
	ctx := r.Context()

	//Convertir str a integer
	id_user, err := strconv.Atoi(idStr)
	if err != nil {
		response.HTTPError(w, r, http.StatusNotFound, err.Error())
		return
	}

	//uint significa e  ntero sin signo, es decir siempre un entero positivo
	pedidos, err := a.Repository.GetByUser(ctx, uint(id_user))
	if err != nil {
		response.HTTPError(w, r, http.StatusNotFound, err.Error())
		return
	}

	response.JSON(w, r, http.StatusOK, response.Map{"ArticlesVideoGames": pedidos})

}

/*
	GET ONE ARTICLES
*/

func (a *PedidoRouter) GetOneHandler(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		response.HTTPError(w, r, http.StatusBadRequest, err.Error())
		return
	}

	ctx := r.Context()
	pedido, err := a.Repository.GetOne(ctx, uint(id))
	if err != nil {
		response.HTTPError(w, r, http.StatusNotFound, err.Error())
		return
	}

	response.JSON(w, r, http.StatusOK, response.Map{"PedidoVideoGame": pedido})
}

/*
	UPDATE ARTICLES
*/

func (a *PedidoRouter) UpdateHandler(w http.ResponseWriter, r *http.Request) {

	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		response.HTTPError(w, r, http.StatusBadRequest, err.Error())
		return
	}

	var pedidoRepo pedido.PedidoGame
	err = json.NewDecoder(r.Body).Decode(&pedidoRepo)

	if err != nil {
		response.HTTPError(w, r, http.StatusBadRequest, err.Error())
		//GOlang regresa los parametros de la funcion por defecto en el return, creo.
		return
	}

	defer r.Body.Close()

	ctx := r.Context()
	pedido, err := a.Repository.Update(ctx, uint(id), pedidoRepo)
	if err != nil {
		response.HTTPError(w, r, http.StatusNotFound, err.Error())
		return
	}

	//response.JSON(w, r, http.StatusOK, nil)
	response.JSON(w, r, http.StatusOK, response.Map{"pedido VideoGame": pedido})

}

//ROUTERS

func (a *PedidoRouter) Routes() http.Handler {
	r := chi.NewRouter()

	r.Get("/", a.GetAllHandler)  //LIST
	r.Post("/", a.CreateHandler) //CREATE

	r.Get("/{id}/", a.GetOneHandler) //DETAIL
	r.Put("/{id}/", a.UpdateHandler) //UPDATE

	r.Delete("/{id}/", a.DeleteHandler)          //DELETE
	r.Get("/user/{idUser}/", a.GetByUserHandler) //LIST ARTICLES BY ID USER
	return r
}
