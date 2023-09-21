package entities

type Task struct {
	ID      int64  `json:"id"`
	Name    string `json:"name"`
	Content string `json:"content"`
}

type AllTask []Task

var task = AllTask{
	{
		ID:      1,
		Name:    "Task1",
		Content: "task number 1",
	},
}
