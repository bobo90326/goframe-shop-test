package model

// RoleCreateUpdateBase is the base struct for role create and update.
type RoleCreateUpdateBase struct {
	Name string
	Desc string
}

// RoleCreateInput is the input of the role create.
type RoleCreateInput struct {
	RoleCreateUpdateBase
}

// RoleCreateOutput is the output of the role create.
type RoleCreateOutput struct {
	RoleId int `json:"role_id"`
}
