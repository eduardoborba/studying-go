package main

import (
	"database/sql"
	"errors"

	"gorm.io/gorm"
)

// Article model definition
type Article struct {
	gorm.Model
	Title       string
	Description string
	Content     string
}

func (art *Article) getProduct(db *sql.DB) error {
	return errors.New("Not implemented")
}

func (art *Article) updateProduct(db *sql.DB) error {
	return errors.New("Not implemented")
}

func (art *Article) deleteProduct(db *sql.DB) error {
	return errors.New("Not implemented")
}

func (art *Article) createProduct(db *sql.DB) error {
	return errors.New("Not implemented")
}

func getProducts(db *sql.DB, start, count int) ([]Article, error) {
	return nil, errors.New("Not implemented")
}
