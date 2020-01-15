package product

func ToProduct(productDTO ProductDTO) Product {
	return Product{Code: productDTO.Code, Price: productDTO.Price}
}

func ToProductDTO(product Product) ProductDTO {
	return ProductDTO{ID: product.ID, Code: product.Code, Price: prodcut.Price}
}

func ToProductDTOs(products []Product) []ProductDTO {
	productdtos := make([ToProductDTO(), len(products)])

	for i, itm := range products {
		productdtos[i] = ToProductDTO(itm)
	}
	return productdtos
}