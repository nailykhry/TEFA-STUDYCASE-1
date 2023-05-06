package repository

import (
	"TEFA-STUDYCASE-1/database"
	"TEFA-STUDYCASE-1/models"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const TasksCollection = "tasks"

type TasksRepository interface {
	CreateTask(task *models.Task) error
	UpdateTask(task *models.Task) error
	GetTaskById(id string) (task *models.Task, err error)
	// GetAllTasks() (tasks []*models.Task, err error)
	DeleteTask(id string) error
}

type tasksRepository struct {
	c *mgo.Collection
}

func NewTaskRepository(conn database.Connection) TasksRepository {
	return &tasksRepository{conn.DB().C(TasksCollection)}
}

func (r *tasksRepository) CreateTask(task *models.Task) error {
	return r.c.Insert(task)
}

func (r *tasksRepository) UpdateTask(task *models.Task) error {
	return r.c.UpdateId(task.Id, task)
}

func (r *tasksRepository) GetTaskById(id string) (task *models.Task, err error) {
	err = r.c.FindId(bson.ObjectIdHex(id)).One(&task)
	return task, err
}

// func (r *tasksRepository) GetAllTasks() (tasks []*models.Task, err error) {
// 	err = r.c.Find(bson.M{}).All(&tasks)
// 	return tasks, err
// }

func (r *tasksRepository) DeleteTask(id string) error {
	return r.c.RemoveId(bson.ObjectIdHex(id))
}
