package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Leave struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Status string `json:"status"`
}

var leaves []Leave

func main() {
	// Inisialisasi data awal
	leaves = append(leaves, Leave{ID: "1", Name: "John Doe", Status: "Pending"})
	leaves = append(leaves, Leave{ID: "2", Name: "Jane Smith", Status: "Approved"})

	// Mendaftarkan endpoint-endpoint yang tersedia
	http.HandleFunc("/leaves", getLeaves)
	http.HandleFunc("/leaves/create", createLeave)
	http.HandleFunc("/leaves/update", updateLeave)

	// Menjalankan server pada port 8080
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func getLeaves(w http.ResponseWriter, r *http.Request) {
	// Mengembalikan data cuti dalam bentuk JSON
	json.NewEncoder(w).Encode(leaves)
}

func createLeave(w http.ResponseWriter, r *http.Request) {
	// Mengambil data dari request body
	var newLeave Leave
	json.NewDecoder(r.Body).Decode(&newLeave)

	// Menambahkan data cuti baru ke dalam slice
	leaves = append(leaves, newLeave)

	// Mengembalikan status 201 Created
	w.WriteHeader(http.StatusCreated)
}

func updateLeave(w http.ResponseWriter, r *http.Request) {
	// Mengambil data dari request body
	var updatedLeave Leave
	json.NewDecoder(r.Body).Decode(&updatedLeave)

	// Mencari data cuti berdasarkan ID
	for i, leave := range leaves {
		if leave.ID == updatedLeave.ID {
			// Mengupdate status cuti
			leaves[i].Status = updatedLeave.Status
			break
		}
	}

	// Mengembalikan status 204 No Content
	w.WriteHeader(http.StatusNoContent)
}

func deleteLeave(w http.ResponseWriter, r *http.Request) {
	// Mengambil data dari request body
	var deletedLeave Leave
	json.NewDecoder(r.Body).Decode(&deletedLeave)

	// Mencari data cuti berdasarkan ID dalam database
	for i, leave := range leaves {
		if leave.ID == deletedLeave.ID {
			// Menghapus data cuti dari database
			leaves = append(leaves[:i], leaves[i+1:]...)
			break
		}
	}

	// Mengembalikan status 204 No Content
	w.WriteHeader(http.StatusNoContent)
}
