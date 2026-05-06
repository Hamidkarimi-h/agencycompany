package model

import "time"

type Agency struct{
	ID uint
	Name string
	Address string
	Phone string
	MembershipDate time.Time
	EmployeeCount *uint32
	Region string
}
