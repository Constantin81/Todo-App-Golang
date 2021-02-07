package todo

type TodoList struct {
	ID          int    `json: "id"`
	Title       string `json: "title"`
	Description string `json: "description"`
}

type UserList struct {
	Id     int
	UserId int
	ListID int
}

type TodoItem struct {
	ID          int    `json: "id"`
	Title       string `json: "title"`
	Description string `json: "description"`
	Done        bool   `json: "done"`
}

type ListItem struct {
	Id     int
	ListID int
	ItemId int
}
