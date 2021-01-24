package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"strconv"
	"testing"
)

var a App

func TestMain(m *testing.M) {
	a.Initialize(
		os.Getenv("APP_DB_USERNAME"),
		os.Getenv("APP_DB_PASSWORD"),
		os.Getenv("APP_DB_NAME"))

	ensureTableExists()
	code := m.Run()
	clearTable()
	os.Exit(code)
}

func ensureTableExists() {
	if err := a.DB.AutoMigrate(); err != nil {
		log.Fatal(err)
	}
}

func clearTable() {
	a.DB.Exec("DELETE FROM articles")
	a.DB.Exec("ALTER SEQUENCE articles_id_seq RESTART WITH 1")
}

func executeRequest(req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	a.Router.ServeHTTP(rr, req)

	return rr
}

func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}
}

func addArticles(count int) {
	if count < 1 {
		count = 1
	}

	for i := 0; i < count; i++ {
		article := Article{Title: "Title" + strconv.Itoa(i), Description: "Description" + strconv.Itoa(i), Content: "Content" + strconv.Itoa(i)}
		a.DB.Create(&article)
	}
}

func TestEmptyTable(t *testing.T) {
	clearTable()

	req, _ := http.NewRequest("GET", "/ds", nil)
	response := executeRequest(req)

	checkResponseCode(t, http.StatusOK, response.Code)

	if body := response.Body.String(); body != "[]" {
		t.Errorf("Expected an empty array. Got %s", body)
	}
}

func TestGetNonExistentd(t *testing.T) {
	clearTable()

	req, _ := http.NewRequest("GET", "/ds/11", nil)
	response := executeRequest(req)

	checkResponseCode(t, http.StatusNotFound, response.Code)

	var m map[string]string
	json.Unmarshal(response.Body.Bytes(), &m)
	if m["error"] != "Article not found" {
		t.Errorf("Expected the 'error' key of the response to be set to 'Article not found'. Got '%s'", m["error"])
	}
}

func TestCreateArticle(t *testing.T) {
	clearTable()

	var jsonStr = []byte(`{"title":"Some title", "description", "Some Desc", "content": "Some Content"}`)
	req, _ := http.NewRequest("POST", "/articles", bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")

	response := executeRequest(req)
	checkResponseCode(t, http.StatusCreated, response.Code)

	var m map[string]interface{}
	json.Unmarshal(response.Body.Bytes(), &m)

	if m["title"] != "Some Title" {
		t.Errorf("Expected article title to be 'Some Title'. Got '%v'", m["title"])
	}

	if m["description"] != "Some Desc" {
		t.Errorf("Expected article desc to be 'Some Desc'. Got '%v'", m["description"])
	}

	if m["content"] != "Some Content" {
		t.Errorf("Expected article content to be 'Some Content'. Got '%v'", m["content"])
	}

	if m["id"] != 1.0 {
		t.Errorf("Expected article ID to be '1'. Got '%v'", m["id"])
	}
}

func TestGetArticle(t *testing.T) {
	clearTable()
	addArticles(1)

	req, _ := http.NewRequest("GET", "/articles/1", nil)
	response := executeRequest(req)

	checkResponseCode(t, http.StatusOK, response.Code)
}

func TestUpdateArticle(t *testing.T) {
	clearTable()
	addArticles(1)

	req, _ := http.NewRequest("GET", "/articles/1", nil)
	response := executeRequest(req)
	var originalArticle map[string]interface{}
	json.Unmarshal(response.Body.Bytes(), &originalArticle)

	var jsonStr = []byte(`{"title":"Updated title", "description", "Updated Description", "content": "Updated Content"}`)
	req, _ = http.NewRequest("PUT", "/articles/1", bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")

	response = executeRequest(req)

	checkResponseCode(t, http.StatusOK, response.Code)

	var m map[string]interface{}
	json.Unmarshal(response.Body.Bytes(), &m)

	if m["id"] != originalArticle["id"] {
		t.Errorf("Expected the id to remain the same (%v). Got %v", originalArticle["id"], m["id"])
	}

	if m["title"] != "Updated Title" {
		t.Errorf("Expected article title to be 'Updated Title'. Got '%v'", m["title"])
	}

	if m["description"] != "Updated Description" {
		t.Errorf("Expected article desc to be 'Updated Description'. Got '%v'", m["description"])
	}

	if m["content"] != "Updated Content" {
		t.Errorf("Expected article content to be 'Updated Content'. Got '%v'", m["content"])
	}
}

func TestDeleteArticle(t *testing.T) {
	clearTable()
	addArticles(1)

	req, _ := http.NewRequest("GET", "/articles/1", nil)
	response := executeRequest(req)
	checkResponseCode(t, http.StatusOK, response.Code)

	req, _ = http.NewRequest("DELETE", "/articles/1", nil)
	response = executeRequest(req)

	checkResponseCode(t, http.StatusOK, response.Code)

	req, _ = http.NewRequest("GET", "/articles/1", nil)
	response = executeRequest(req)
	checkResponseCode(t, http.StatusNotFound, response.Code)
}
