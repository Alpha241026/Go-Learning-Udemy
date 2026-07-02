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

type Customer struct {
	CustomerID      int
	Name            string
	Email           string
	BillingAddress  Address //composition
	ShippingAddress Address //composition
}

func (c Customer) PrintDetails() { //displays customer & delegates address formatting to address type
	fmt.Printf("Customer ID: %d\n", c.CustomerID)
	fmt.Printf("Customer Name: %s\n", c.Name)
	fmt.Printf("Customer Email: %s\n", c.Email)
	fmt.Println("Billing Address:", c.BillingAddress.FullAddress())
	fmt.Println("Shipping Address:", c.ShippingAddress.FullAddress())
}

func main() {

	fmt.Println("---------- Composition ----------")

	cus1 := Customer{
		CustomerID: 1001,
		Name:       "Gadget Corp",
		Email:      "sales@gadgetcorp.com",
		BillingAddress: Address{
			Street:  "123 Tech Road",
			City:    "Innovateville",
			State:   "CA",
			ZipCode: "90210",
		},
		ShippingAddress: Address{
			Street:  "456 Factory Lane",
			City:    "Manufacturicity",
			State:   "NV",
			ZipCode: "89101",
		},
	}

	cus1.PrintDetails()

	fmt.Println("---------- customer with same billing and shipping address ----------")

	mainAddress := Address{
		Street:  "789 Main Street",
		City:    "Anytown",
		State:   "TX",
		ZipCode: "75001",
	}

	cus2 := Customer{
		CustomerID:      1002,
		Name:            "John Doe",
		Email:           "john.doe@example.com",
		BillingAddress:  mainAddress,
		ShippingAddress: mainAddress,
	}

	cus2.PrintDetails()
}
