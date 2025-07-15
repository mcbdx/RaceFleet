package car

// BMW struct implements the Car interface

type BMW struct {
	BaseCar
}

// new BMW - creates car - simulates a constructor
// Note: In Go, we typically use a function to create a new instance of a struct
func NewBMW(id, make, model, year) {
	return BMW{
		BaseCar: BaseCar{
			id: id, 
			Make: make,
			Model: model,
			Year: year,
		}
	}
}


