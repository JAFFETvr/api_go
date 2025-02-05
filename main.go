package main

import (
	"demo/src/products/application"
	"demo/src/products/infraestructure"
	"demo/src/products/infraestructure/repositories"

	clients_infraestructure "demo/src/Clients/infraestructure"
	clients_repositories "demo/src/Clients/infraestructure/repositories" // Alias para evitar conflicto de nombres
	clients_application "demo/src/Clients/applications" // Alias para evitar conflicto de nombres

	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	mysql := infraestructure.NewMySQL()
	defer mysql.Close()

	// Productos
	productRepo := repositories.NewProductRepository(mysql.DB)

	createProduct := application.NewCreateProduct(*productRepo)
	getProducts := application.NewGetProducts(*productRepo)
	getProductById := application.NewGetProductById(*productRepo) //
	updateProduct := application.NewUpdateProduct(*productRepo)
	deleteProduct := application.NewDeleteProduct(*productRepo)

	createProductController := infraestructure.NewCreateProductController(createProduct)
	getProductsController := infraestructure.NewGetProductsController(getProducts)
	getProductByIdController := infraestructure.NewGetProductByIdController(getProductById) 
	updateProductController := infraestructure.NewUpdateProductController(updateProduct)
	deleteProductController := infraestructure.NewDeleteProductController(deleteProduct)

	// Clientes
	clientRepo := clients_repositories.NewClientRepository(mysql.DB)

	createClient := clients_application.NewCreateClient(clientRepo)
	getClients := clients_application.NewGetClient(clientRepo)
	updateClient := clients_application.NewUpdateClient(clientRepo)
	deleteClient := clients_application.NewDeleteClient(clientRepo)

	createClientController := clients_infraestructure.NewCreateClientController(createClient)
	getClientsController := clients_infraestructure.NewGetClientsController(getClients)
	updateClientController := clients_infraestructure.NewUpdateClientController(updateClient)
	deleteClientController := clients_infraestructure.NewDeleteClientController(deleteClient)

	router := gin.Default()

	productRoutes := infraestructure.NewProductRoutes(
		createProductController,
		getProductsController,
		updateProductController,
		deleteProductController,
		getProductByIdController, 
	)
	productRoutes.SetupRoutes(router)

	// Rutas de clientes
	clientsRoutes := clients_infraestructure.NewClientRoutes(
		createClientController,
		getClientsController,
		updateClientController,
		deleteClientController,
	)
	clientsRoutes.SetupRoutes(router)

	// Iniciar el servidor
	log.Println("[Main] Servidor corriendo en http://localhost:8080")
	router.Run(":8080")
}
