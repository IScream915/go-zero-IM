package logic

type User struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Phone string `json:"phone"`
}

var users = map[string]*User{
	"1": {
		Id:    "1",
		Name:  "admin",
		Phone: "11111111",
	},
	"2": {
		Id:    "2",
		Name:  "cyyyx",
		Phone: "22222222",
	},
}
