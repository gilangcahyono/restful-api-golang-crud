package src

import (
	"encoding/json"
	"strconv"

	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

type badRes struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
}

type okRes struct {
	Status  bool      `json:"status"`
	Message string    `json:"message,omitempty"`
	Data    []Product `json:"data"`
}

func GetProduct(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	products, err := Get()

	if err != nil {
		res := badRes{
			Status:  false,
			Message: "Internal Server Error!",
		}
		json.NewEncoder(w).Encode(res)
	} else {
		if len(products) > 0 {
			res := okRes{
				Status: true,
				Data:   products,
			}
			json.NewEncoder(w).Encode(res)
		} else {
			res := badRes{
				Status:  false,
				Message: "Data tidak ditemukan!",
			}
			json.NewEncoder(w).Encode(res)
		}
	}
}

func GetProductById(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		res := badRes{
			Status:  false,
			Message: "Internal Server Error!",
		}
		json.NewEncoder(w).Encode(res)
	} else {
		product, err := GetById(int64(id))
		if err != nil {
			res := badRes{
				Status:  false,
				Message: "Data tidak ditemukan!",
			}
			json.NewEncoder(w).Encode(res)
		} else {
			json.NewEncoder(w).Encode(product)
		}
	}
}

func CreateProduct(w http.ResponseWriter, r *http.Request) {

	var product Product
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		res := badRes{
			Status:  false,
			Message: "Request data invalid!",
		}
		json.NewEncoder(w).Encode(res)
	} else {

		insertID := Create(product)

		if insertID == -1 {
			res := badRes{
				Status:  false,
				Message: "Internal server error",
			}
			json.NewEncoder(w).Encode(res)
		} else {
			res := badRes{
				Status:  true,
				Message: "Data berhasil ditambahkan",
			}
			json.NewEncoder(w).Encode(res)
		}
	}
}

func UpdateProduct(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])

	if err != nil {
		res := badRes{
			Status:  false,
			Message: "Internal Server Error!",
		}
		json.NewEncoder(w).Encode(res)
	} else {
		var product Product
		err = json.NewDecoder(r.Body).Decode(&product)
		if err != nil {
			res := badRes{
				Status:  false,
				Message: "Request data invalid!",
			}
			json.NewEncoder(w).Encode(res)
		} else {
			updatedRows := Update(int64(id), product)
			if updatedRows > 0 {
				res := badRes{
					Status:  true,
					Message: "Data berhasil di update",
				}
				json.NewEncoder(w).Encode(res)
			} else {
				res := badRes{
					Status:  false,
					Message: "Data tidak ditemukan, Update gagal!",
				}
				json.NewEncoder(w).Encode(res)
			}
		}
	}
}

func DeleteProduct(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		res := badRes{
			Status:  false,
			Message: "Internal Server Error!",
		}
		json.NewEncoder(w).Encode(res)
	} else {

		deletedRows := Delete(int64(id))
		if deletedRows > 0 {
			res := badRes{
				Status:  true,
				Message: "Data berhasil dihapus",
			}
			json.NewEncoder(w).Encode(res)
		} else {
			res := badRes{
				Status:  false,
				Message: "Data tidak ditemukan, Hapus gagal!",
			}
			json.NewEncoder(w).Encode(res)
		}
	}
}
