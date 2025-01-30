package main

import (
	"demo/src/products/application"
	"demo/src/products/infraestructure"
	"demo/src/products/infraestructure/repositories"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	mysql := infraestructure.NewMySQL()
	defer mysql.Close()

	productRepo := repositories.NewProductRepository(mysql.DB)

	createProduct := application.NewCreateProduct(*productRepo)
	getProducts := application.NewGetProducts(*productRepo)
	updateProduct := application.NewUpdateProduct(*productRepo)
	deleteProduct := application.NewDeleteProduct(*productRepo) 

	createProductController := infraestructure.NewCreateProductController(createProduct)
	getProductsController := infraestructure.NewGetProductsController(getProducts)
	updateProductController := infraestructure.NewUpdateProductController(updateProduct)
	deleteProductController := infraestructure.NewDeleteProductController(deleteProduct) 

	router := gin.Default()
	productRoutes := infraestructure.NewProductRoutes(createProductController, getProductsController, updateProductController, deleteProductController) 
	productRoutes.SetupRoutes(router)

	log.Println("[Main] Servidor corriendo en http://localhost:8080")
	router.Run(":8080")
}
