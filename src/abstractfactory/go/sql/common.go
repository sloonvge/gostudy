package sql

type IUser interface {
	Insert(u *User)
	GetUser(i int) (u *User)
}

type IDepartment interface {
	Insert(d *Department)
	GetDepartment(i int) (d *Department)
}

type IFactory interface {
	CreateUser() (i IUser)
	CreateDept() (d IDepartment)
}

type DataAccess struct {
	DB string
}

func NewDataAccess() *DataAccess{
	db := "SqlServer"
	//db := "Access"
	return &DataAccess{DB: db}
}

func (da *DataAccess) CreateUser() (i IUser){
	switch da.DB {
	case "SqlServer":
		i = NewSqlServerUser()
	default:
		i = NewAccessUser()

	}
	return
}

func (da *DataAccess) CreateDepartment() (i IDepartment) {
	switch da.DB {
	case "SqlServer":
		i = NewSqlServerDepartment()
	default:
		i = NewAccessDepartment()
	}
	return
}