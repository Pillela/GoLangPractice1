package Math_Functions

import "fmt"

type Employee_biodata struct{
	Name string
	Age int
	Gender string
}

type Employee_project struct{
	Employee_biodata
	Client_name string
	Vendor_name string
	Contract_duration_months int
}

func (a *Employee_biodata) Basic_info(){
	a.Name = a.Name + "-Sage_it"
	fmt.Println("Emp name:",a.Name)
	fmt.Println("Emp age:",a.Age)
	fmt.Println("Emp gend:",a.Gender)
}

func (a *Employee_project) Project_info(){
	a.Basic_info()
	fmt.Println("Client Name:",a.Client_name)
	fmt.Println("Vendor Name:",a.Vendor_name)
	fmt.Println("Contract Duration in months:", a.Contract_duration_months)
}

