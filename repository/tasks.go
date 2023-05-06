package repository

import (
	"TEFA-STUDYCASE-1/database" // O(1)
	"TEFA-STUDYCASE-1/models"   // O(1)

	"gopkg.in/mgo.v2"      // O(1)
	"gopkg.in/mgo.v2/bson" // O(1)
)

const TasksCollection = "tasks" // O(1)

type TasksRepository interface {
	CreateTask(task *models.Task) error                   // O(1)
	UpdateTask(task *models.Task) error                   // O(1)
	GetTaskById(id string) (task *models.Task, err error) // O(1)
	GetAllTasks() (tasks []*models.Task, err error)       // O(1)
	DeleteTask(id string) error                           // O(1)
}

type tasksRepository struct {
	c *mgo.Collection // O(1)
}

func NewTaskRepository(conn database.Connection) TasksRepository {
	return &tasksRepository{conn.DB().C(TasksCollection)} // O(1)
}

func (r *tasksRepository) CreateTask(task *models.Task) error {
	return r.c.Insert(task) // O(1)
}

func (r *tasksRepository) UpdateTask(task *models.Task) error {
	return r.c.UpdateId(task.Id, task) // O(1)
}

func (r *tasksRepository) GetTaskById(id string) (task *models.Task, err error) {
	err = r.c.FindId(bson.ObjectIdHex(id)).One(&task) // O(1)
	return task, err
}

func (r *tasksRepository) GetAllTasks() (tasks []*models.Task, err error) {
	err = r.c.Find(bson.M{}).All(&tasks) // O(n)
	return tasks, err
}

func (r *tasksRepository) DeleteTask(id string) error {
	return r.c.RemoveId(bson.ObjectIdHex(id)) // O(1)
}
