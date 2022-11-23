package student

import "fmt"

type Address struct {
	Street string
	City   string
	Pin    int
}

type Student struct {
	name    string
	age     int
	Address []Address
}

func New() []Student {
	return []Student{
		{
			name: "John",
			age:  20,
			Address: []Address{
				{
					Street: "abc lane",
					City:   "mumbai",
					Pin:    422221,
				},
				{
					Street: "xyz lane",
					City:   "mumbai",
					Pin:    422220,
				},
			},
		},
		{
			name: "ABC",
			age:  20,
			Address: []Address{
				{
					Street: "abc lane",
					City:   "mumbai",
					Pin:    422221,
				},
				{
					Street: "xyz lane",
					City:   "mumbai",
					Pin:    422220,
				},
			},
		},
		{
			name: "PQR",
			age:  20,
			Address: []Address{
				{
					Street: "abc lane",
					City:   "mumbai",
					Pin:    422221,
				},
				{
					Street: "xyz lane",
					City:   "mumbai",
					Pin:    422220,
				},
			},
		},
	}
}

func (a Address) String() string {
	return fmt.Sprintf("%s, %s - %d", a.Street, a.City, a.Pin)
}
