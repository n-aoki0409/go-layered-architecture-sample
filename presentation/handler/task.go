package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/n-aoki0409/go-layered-architecture-sample/usecase"
)

type TaskHandler interface {
	Post(c *gin.Context)
	Get(c *gin.Context)
	Put(c *gin.Context)
	Delete(c *gin.Context)
}

type taskHandler struct {
	taskUsecase usecase.TaskUsecase
}

func NewTaskHandler(taskUsecase usecase.TaskUsecase) TaskHandler {
	return &taskHandler{taskUsecase: taskUsecase}
}

type requestTask struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

type responseTask struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

func (th *taskHandler) Post(c *gin.Context) {
	var req requestTask
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
	}

	createdTask, err := th.taskUsecase.Create(req.Title, req.Content)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
	}

	res := responseTask{
		ID:      createdTask.ID,
		Title:   createdTask.Title,
		Content: createdTask.Content,
	}

	c.JSON(http.StatusCreated, res)
}

func (th *taskHandler) Get(c *gin.Context) {
	id, err := strconv.Atoi((c.Param("id")))
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
	}

	foundTask, err := th.taskUsecase.FindByID(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
	}

	res := responseTask{
		ID:      foundTask.ID,
		Title:   foundTask.Title,
		Content: foundTask.Content,
	}

	c.JSON(http.StatusOK, res)
}

func (th *taskHandler) Put(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
	}

	var req requestTask
	if err := c.Bind(&req); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
	}

	updatedTask, err := th.taskUsecase.Update(id, req.Title, req.Content)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
	}

	res := responseTask{
		ID:      updatedTask.ID,
		Title:   updatedTask.Title,
		Content: updatedTask.Content,
	}

	c.JSON(http.StatusOK, res)
}

func (th *taskHandler) Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
	}

	err = th.taskUsecase.Delete(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
	}

	c.JSON(http.StatusNoContent, gin.H{"message": "delete succeed"})
}
