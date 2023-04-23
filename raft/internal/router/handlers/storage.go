package handlers

import "github.com/gin-gonic/gin"

type StorageHandler struct{}

func NewStorageHandler() *StorageHandler {
	return new(StorageHandler)
}

func (s *StorageHandler) Set(c *gin.Context) {

}

func (s *StorageHandler) Get(c *gin.Context) {

}
