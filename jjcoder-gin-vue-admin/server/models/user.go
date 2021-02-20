package models

type User struct {
	ID       int     `gorm:"column:id;primary_key" json:"id"`
	Married  int     `gorm:"column:married" json:"married"`
	Salary   float32 `gorm:"column:salary" json:"salary"`
	Card     string  `gorm:"column:card" json:"card"`
	Username string  `gorm:"column:username" json:"username"`
	Password string  `gorm:"column:password" json:"password"`
	Email    string  `gorm:"column:email" json:"email"`
	Age      int     `gorm:"column:age" json:"age"`
	Sex      string  `gorm:"column:sex" json:"sex"`
	Tel      string  `gorm:"column:tel" json:"tel"`
	Addr     string  `gorm:"column:addr" json:"addr"`
}
