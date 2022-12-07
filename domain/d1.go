package domain

type User struct {
	UserName string
	UpdateAt int64
}

type UserWithRepo struct {
	Users    []User
	Counts   int
	UpdateAt int64
}

type UserWithBigModel struct {
	BigModelType string
	Users        []User
	Counts       int
	UpdateAt     int64
}

// func (u *UserWithRepo) CountUWR() (count string, err error) {
// 	c := len(u.Users)
// 	count = strconv.Itoa(c)
// 	return count, nil
// }
