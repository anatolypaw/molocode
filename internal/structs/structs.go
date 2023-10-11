package structs

/* Управление пользователями */
type User struct {
	Login string 	`bson:"login" json:"login"`
	Password string `bson:"password" json:"password"` //TODO передалать на хранение хэша
	Role string 	`bson:"role" json:"role"`	  //admin, user
}

