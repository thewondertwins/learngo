# Creating Useful Interfaces

In this section we’ll look at Go’s interfaces and learn how to use them for maximum flexibility and code reuse.  We’ll single out several interfaces from the standard library and explore what makes them so powerful.  You’ll learn how to make your abstractions count with small and concise interfaces.

- What Makes a Good Interface
- Standard Library Interface Examples
- Interface Size versus Interface Utility

# What Makes a Good Interface

- Interfaces should represent behaviors
- Interfaces should wrap the behavior of a dependency
- Good interfaces represent the smallest behavior possible

## Philosophy: Go and Interfaces

Proverb:

	The bigger the interface, the weaker the abstraction

 _Rob_Pike_ at [Gopherfest SV, 2015](https://www.youtube.com/watch?v=PAAkCSZUG1c)

## What Makes a Good Interface

To define your interfaces, start by defining the behaviors of your types.

	// A Product is an item that our company stocks
	type Product struct {
		ID      int
		Name    string
		SKU     string
		LotSize int
	}

	// ProductStorage defines the behaviors required to
	// perform CRUD operations on an Order
	type ProductStorage interface {
		Get(id int) (*Product, error)
		Create(p Product) (*Product, error)
		Update(p *Product) error
		Delete(p *Product) error
	}

An interface should represent one type of behavior. In the previous example it was _Storage_.

An interface should not mix different types behaviors.  Don't have a ProductInterface that mixes _Storage_ with _Display_ for example.

## Practical Interfaces

We're going to use an Inventory application as an example that models something that you might write for your business.

*Business Domain Information*

- Products -- things we sell
- Suppliers -- vendors that provide Products
- Orders -- requests from us to our supplier to get more Products

We store our data in *postgres*

We have one Supplier *Acme*, with an HTTP API for placing orders

## Practical Interfaces - Dependencies

Our fictional inventory application has two data sources, a PostgreSQL database and a vendor API.

Suppliers, Products, and Orders are stored in PostgreSQL, but the interactions with the supplier are done through the supplier's HTTP API.

That means we have two dependencies which should be in their own packages: 

- postgres
- acme


The *postgres* package will have implementations of the domain interfaces for each of the domain types that are stored in the PostgreSQL database.

The *acme* package will have an implementation of the domain interface for interacting with the supplier.

## Practical Interfaces - Transport Mechanisms

For compatibility with other systems, the inventory application will offer consumers of the data two choices in network transport:

- HTTP/REST API
- net/rpc calls 

Therefore we'll have a package *transport* containing an *rpc* package and an *http* package.

## Package Organization - Interface

**PRO TIP:** 

Use https://github.com/josharian/impl to generate your implementation stubs

	go get github.com/josharian/impl

	USAGE:
	impl 'r *Receiver' gopath/to/pkg.Interface

	Example:
	impl 's *Source' golang.org/x/oauth2.TokenSource > mocks/token.go


To make it easy to isolate and test each individual package you should create mock implementations of each Interface you define in the domain package.  Keep all your mock implementations in a *mocks* package in the root of your application directory. 

*impl* requires that the package with the interfaces can compile.  Make sure you have a compiling package before using it.

## Practical Interfaces - Exercise 1

In the *$GOPATH/src/github.com/gophertrain/material/usefulinterfaces/exercises* directory there's a folder called *impl*

	cd impl
	cat interface.go

Now install the *impl* tool:
	
	go get github.com/josharian/impl

Try it out:

	$ impl 'c *CustomerService' github.com/gophertrain/material/usefulinterfaces/exercises/impl.CustomerGetter

You should see output like this:

	func (c *CustomerService) Get(id int) (*impl.Customer, error) {
		panic("not implemented")
	} 

That's very useful for both implementing the actual service and implementing the mocks.  But it outputs to *stdout*.  Let's make a postgres implementation of the CustomerGetter service:

	impl 'c *CustomerService' {full gopath}/impl.CustomerGetter > postgres/customer.go

The postgres folder already exists, if it didn't you'd have to create it first.

Now make an implementation for a mock of the CustomerGetter:

	mkdir mocks
	impl 'c *CustomerService' {full gopath}impl.CustomerGetter > mocks/customer.go

*impl* generates the methods for a type that doesn't exist.  To finish the implementation, you need to add a type called *CustomerService*.

	type CustomerService struct{}

Now your code should compile.

## Practical Interfaces - Example Application

	$ cd $GOPATH/src/github.com/gophertrain/material/usefulinterfaces/demos/inventory

In the root package *inventory* there are three domain objects:

- *Order* in  order.go
- *Product* in product.go
- *Supplier* in supplier.go

The *inventory* package also defines five *interfaces*:

- *OrderStorage* in  order.go
- *ProductStorage* in product.go
- *SupplierStorage* in supplier.go
- *SupplierService* in supplier.go
- *InventoryService* in inventory.go

Our application has *storage* interfaces, so there are (incomplete) implementations of those *storage* interfaces in the *postgres* package.

	type OrderService struct {
		db *sql.DB
	}

	func NewOrderService(db *sql.DB) inventory.OrderStorage {
		return &OrderService{db: db}
	}

	func (s *OrderService) Get(id int) (*inventory.Order, error) {
		panic("not implemented")
	}

	func (s *OrderService) Create(o inventory.Order) (*inventory.Order, error) {
		panic("not implemented")
	}

	func (s *OrderService) Cancel(o *inventory.Order) error {
		panic("not implemented")
	} 

Similarly there's a *SupplierService* that represents the interactions over an external API to our supplier "ACME Widgets".  The implementation of that service is in the *acme* package.

	type AcmeClientService struct {
		URL string
	}

	func NewClient(url string) inventory.SupplierService {
		return &AcmeClientService{URL: url}
	}

	func (a *AcmeClientService) PlaceOrder(o *inventory.Order) error {
		panic("not implemented")
	}

	func (a *AcmeClientService) GetStatus(o *inventory.Order) error {
		panic("not implemented")
	}

The *Service* interface defines the external API of our application to its consumers.  

	type Service interface {
		GetOrder(GetOrderRequest, *GetOrderResponse) error
		CreateOrder(CreateOrderRequest, *CreateOrderResponse) error

		OrderStatus(OrderStatusRequest, *OrderStatusResponse) error
		CancelOrder(CancelOrderRequest, *CancelOrderResponse) error

		GetProduct(GetProductRequest, *GetProductResponse) error
		CreateProduct(CreateProductRequest, *CreateProductResponse) error
		UpdateProduct(UpdateProductRequest, *UpdateProductResponse) error
		DeleteProduct(DeleteProductRequest, *DeleteProductResponse) error

		GetSupplier(GetSupplierRequest, *GetSupplierResponse) error
		CreateSupplier(CreateSupplierRequest, *CreateSupplierResponse) error
		UpdateSupplier(UpdateSupplierRequest, *UpdateSupplierResponse) error
		DeleteSupplier(DeleteSupplierRequest, *DeleteSupplierResponse) error
	}

We have customers who demand a REST style interface over HTTP, and we also have internal users in our own company who can use something faster like Go's *net/rpc*

It makes sense to create a new interface to represent the delivery of our service over any type of connection.  In the *transport* package, there's an interface just for this:


	type InventoryTransporter interface {
		inventory.Service
		Serve(net.Listener) error
	}


Now we can have transport specific implementations in packages under *transport*:

- *http*
- *rpc*

Let's browse through those two implementations quickly.  

Of note, they both implement the *transport.InventoryTransporter* interface, which *embeds* the *inventory.Service*.  

Now our *http* and *rest* implementations of the *transport.InventoryTransporter* interface are required to also implement the *inventory.Service* interface.


**PROTIP:**

Prove that your type implements an interface at compile time with this declaration:

	// Compile-time proof of interface implementation
	var _ transport.InventoryTransporter = (*RESTService)(nil) 

This variable declaration assigns a concrete type *RESTService* to the interface it should implement, then discards the result.

By doing this, your code won't compile if the *RESTService* doesn't implement the *transport.InventoryTransporter* interface. 

	Type Safety -- even when using Interfaces for abstraction

Now with interfaces all defined and (mostly) implemented, we can tie it all together in *main.go*

`src/github.com/gophertrain/material/usefulinterfaces/demos/inventory/cmd/main.go`

	func main() {
		// ommited: DB initialization

		// create data storage
		orderStore := postgres.NewOrderService(db)
		productStore := postgres.NewProductService(db)
		supplierStore := postgres.NewSupplierService(db)

		// supplier service
		acmeClient := acme.NewClient("https://acme.com/api") 

		// ommited: listener setup error checks
		net.Listen("tcp", ":8080")
		net.Listen("tcp", ":8081")

		restServer := http.NewRESTService(orderStore, supplierStore, acmeClient, productStore)
		rpcServer := rpc.NewRPCService(orderStore, supplierStore, acmeClient, productStore) 

		// ommited: server start error checks
		restServer.Serve(hl)
		rpcServer.Serve(rl)

Because the network services take interfaces, we can easily swap out storage drivers when changing databases, or more commonly, for testing.

Every interface in this application represents a behavior that is tied to a dependency.

- Storage: postgres
- Transport: http & rpc
- SupplierService: vendor API

With clearly defined interfaces, we can create new implementations of behaviors as business requirements change.

- Migrate to mysql? - implement the Storage interfaces
- New supplier?  Implement the SupplierService interface
- New service delivery method?  Implement the Transport interface

Your domain types (Order, Supplier, Product, etc) don't know or *care* how they're stored or delivered.

- No circular dependencies
- Easy testing with mocks that implement the interfaces


# Standard Library Interface Examples

## encoding/TextMarshaler

[Implementations](https://sourcegraph.com/github.com/golang/go/-/info/GoPackage/encoding/-/TextMarshaler)


*crypto/Signer*

[Implementations](https://sourcegraph.com/github.com/golang/go/-/info/GoPackage/crypto/-/Signer)


## *database/sql/Driver*

database/sql may be the best example of many interfaces used to represent a single aggregate dependency

[Implementations](https://godoc.org/database/sql/driver)

	The *REAL* power of these interfaces is that many of them are so inter-related.

*sql/driver.Rows* is an iterator over a query's results

*rows.Next(dest*[]Value)* populates a slice of *sql/driver.Value*

*sql/driver.Value* represents base types that any *sql/driver* must be able to retrieve and store.

Convert to and from a *sql/driver.Value* by implementing the *sql/driver.Valuer* and *sql/driver.ValueConverter* interfaces.

# Interface Size vs Interface Utility

Rob Pike's quote at the beginning of the module stated clearly that the larger the interface, the weaker the abstraction.

## Interface Guidelines:

- Start big 

	Abstract out your dependencies first

- Use interfaces for ONE and only ONE behavior

- Don't get sucked into Interface hell!

	Create interfaces only where multiple types can implement the same behavior
	*AND*	
	It is necessary to use that behavior in multiple places without coupling the
	types that implement the behavior.



# Exercise

*$GOPATH/src/github.com/gophertrain/material/usefulinterfaces/exercises/inventory* contains a *mocks* folder.

Create a file in the mocks directory that implements the OrderStorage interface.

- OrderStorage = mocks/orders.go 

*Hint*: use a type that has a map[int]*DomainModel  (e.g. map[int]*inventory.Order)

*Bonus*: Implement Supplier and Product

*Double Bonus Round*: implement a mock for SupplierService in mocks/acme.go
