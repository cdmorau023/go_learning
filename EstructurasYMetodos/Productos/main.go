package main

import "fmt"

type Product struct {
	ID          int    `json:"id,omitempty"`
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	Category    string `json:"category,omitempty"`
}

//Función que dado slice de Product los imprime en consola con fprintf

func (pr *Product) GetAll(p []Product) {
	for _, product := range p {
		fmt.Printf("ID: %d\nName: %s\nDescription: %s\nCategory: %s\n", product.ID, product.Name, product.Description, product.Category)
	}

}

func GetById(id int) Product {

	for _, product := range Products {
		if product.ID == id {
			fmt.Printf("ID: %d\nName: %s\nDescription: %s\nCategory: %s\n", product.ID, product.Name, product.Description, product.Category)
			return product
		}
	}
	return Product{}

}

// Función que recibe un apuntador a un slice de Product y guarda un nuevo producto en el slice

func (pr *Product) Save(p *[]Product) {
	*p = append(*p, *pr)
}

// Constructor de productos
func NewProduct(id int, name, description, category string) Product {
	return Product{
		ID:          id,
		Name:        name,
		Description: description,
		Category:    category,
	}
}

var Products = []Product{
	{
		ID:          1,
		Name:        "Producto 1",
		Description: "Descripción del producto 1",
		Category:    "Categoría 1",
	},
	{
		ID:          2,
		Name:        "Producto 2",
		Description: "Descripción del producto 2",
		Category:    "Categoría 2",
	},
	{
		ID:          3,
		Name:        "Producto 3",
		Description: "Descripción del producto 3",
		Category:    "Categoría 3",
	},
}

func main() {
	p := NewProduct(4, "Producto 4", "Descripción del producto 4", "Categoría 4")
	p.Save(&Products)
	p.GetAll(Products)
	GetById(2)
}
