package controllers

import "net/http"

func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Creating a new user"))
}

func GetAllUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Get all users"))
}

func GetUserById(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Get user by id"))
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Updating user"))
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Deleting user"))
}