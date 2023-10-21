package food

import uuid "github.com/satori/go.uuid"

func GetRecepie(id uint) (*Recepie, error) {
	return &Recepie{
			ID:   id,
			Name: "Foo FOOD",
			Users: []User{
				{
					uuid.NewV1(),
					"foobar",
					"",
					"foobar@example.com",
					nil,
				},
				{
					uuid.NewV1(),
					"foobar2",
					"",
					"foobar2@example.com",
					nil,
				},
			},
			Produces: []*Produce{
				{
					1,
					"Toast",
					nil,
				},
				{
					1,
					"Butter",
					nil,
				},
			},
			Appliances: []*Appliance{
				{
					1,
					"knife",
					nil,
				},
			},
		},
		nil
}
