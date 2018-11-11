package sql

type SqlServerFactory struct {

}

func NewSqlServerFactory() (s *SqlServerFactory) {
	s = &SqlServerFactory{}
	return
}

func (ssf *SqlServerFactory) CreateUser() (i IUser) {
	i = NewSqlServerUser()
	return
}
func (ssf *SqlServerFactory) CreateDepartment() (i IDepartment) {
	i = NewSqlServerDepartment()
	return
}

type AccessFactory struct {

}
func NewAccessFactory() (s *AccessFactory) {
	s = &AccessFactory{}
	return
}
func (ssf *AccessFactory) CreateUser() (i IUser) {
	i = NewAccessUser()
	return
}
func (ssf *AccessFactory) CreateDepartment() (i IDepartment) {
	i = NewAccessDepartment()
	return
}