package sql

type User struct {
	Id int
	Name string
}

func NewUser() (u *User) {
	u = &User{}
	return
}

func (u *User) Set(i int, n string) {
	u.Id = i
	u.Name = n
}

type Department struct {
	Id int
	dept string
}

func NewDepartment() (d *Department) {
	d = &Department{}
	return
}

func (d *Department) Set(id int, dept string) {
	d.Id = id
	d.dept = dept
}


