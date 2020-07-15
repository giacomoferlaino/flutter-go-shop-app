package product

// Products contains a slice of Product
type Products []Product

func (products *Products) toGenericSlice() []interface{} {
	genericSlice := []interface{}{}
	for _, product := range *products {
		genericSlice = append(genericSlice, product)
	}
	return genericSlice
}
