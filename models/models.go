package models


type Employee struct {
    ID        int     `json:"id"`
    BranchID  int     `json:"branch_id"`
    FirstName string  `json:"first_name"`
    LastName  string  `json:"last_name"`
    Phone     string  `json:"phone"`
    Salary    float64 `json:"salary"`
}

type Course struct {
    ID       int    `json:"id"`
    Name     string `json:"name"`
    BranchID int    `json:"branch_id"`
}

type Student struct {
    ID        int    `json:"id"`
    FirstName string `json:"first_name"`
    LastName  string `json:"last_name"`
    Phone     string `json:"phone"`
    BranchID  int    `json:"branch_id"`
    CourseID  int    `json:"course_id"`
}

type Branch struct {
    ID       int    `json:"id"`
    Name     string `json:"name"`
    Addresses string `json:"addresses"`
}

type EmployeeBranch struct {
    ID         int `json:"id"`
    EmployeeID int `json:"employee_id"`
    BranchID   int `json:"branch_id"`
}