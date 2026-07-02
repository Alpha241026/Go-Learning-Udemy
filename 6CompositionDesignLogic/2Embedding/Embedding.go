package main

import "fmt"

type Address struct {
	Street  string
	City    string
	State   string
	ZipCode string
}

func (a Address) FullAddress() string { //returns address in readable format
	if a.Street == "" && a.City == "" {
		return "No address provided"
	}
	return fmt.Sprintf("%s, %s, %s %s", a.Street, a.City, a.State, a.ZipCode)
}

type ContactInfo struct {
	Email string
	Phone string
}

func (ci ContactInfo) DisplayContact() string {
	return fmt.Sprintf("Email: %s, Phone: %s", ci.Email, ci.Phone)
}

type Company struct {
	Name string
	Address
	ContactInfo
	BusinessType string
}

func (c Company) GetProfile() {
	fmt.Printf("Company Name: %s\n", c.Name)

	fmt.Printf("Location: %s\n", c.FullAddress())
	fmt.Printf("Street (promoted): %s\n", c.Street)

	fmt.Printf("Email (promoted): %s\n", c.Email)
	fmt.Printf("Business Type: %s\n", c.BusinessType)
}

type CompanyWithOwnEmail struct {
	Name string
	Address
	ContactInfo
	Email string
}

func main() {

}
