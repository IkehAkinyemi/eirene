package models

type Store interface {
	CreatePost()
	GetPost()
	AddComment()
	UpdatePost()
}