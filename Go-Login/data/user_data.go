package data

type User struct {
	Name     string
	Email    string
	Password string
}

type SessionData struct {
	SessionId string
	Email     string
}

var UserData = make(map[string]User)
var Sessions = make(map[string]SessionData)

func init() {
	// User Data

	UserData["shahabazsulthan4@gmail.com"] = User{"Shahabaz", "shahabazsulthan4@gmail.com", "Shahabaz@123"}
	UserData["chabbuchabbu1@gmail.com"] = User{"Sulthan", "chabbuchabbu1@gmail.com", "Sulthan@123"}
}
