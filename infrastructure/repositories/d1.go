package repositories

type UserWithRepoMapper interface {
	AddUser(*UserWithRepoDO, *UserDO) error
	Get() (UserWithRepoDO, error)
}

type userWithRepo struct {
	mapper UserWithRepoMapper
}

type UserDO struct {
	UserName string
	UpdateAt int64
}

type UserWithRepoDO struct {
	Type     string
	Users    []UserDO
	Counts   int
	UpdateAt int64
}

type UserWithBigModelDO struct {
	Type         string
	BigModelType string
	Users        []UserDO
	Counts       int
	UpdateAt     int64
}
