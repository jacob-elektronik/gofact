package segments

type ContactInformation struct {
	ContactFunctionCode       string
	DepartmentEmployeeDetails DepartmentEmployeeDetails
}

type DepartmentEmployeeDetails struct {
	DepartmentEmployeeNameCode string
	DepartmentEmployeeName     string
}
