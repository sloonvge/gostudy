package main

import "./sql"

func main() {
	u := sql.NewUser()
	//d := sql.NewDepartment()

	//i := sql.NewAccessFactory()
	//
	//s := i.CreateUser()
	//s.Insert(u)
	//s.GetUser(1)
	//
	//
	//dd := i.CreateDepartment()
	//dd.Insert(d)
	//dd.GetDepartment(1)

	da := sql.NewDataAccess()
	cu := da.CreateUser()
	cu.Insert(u)
	cu.GetUser(1)
}
