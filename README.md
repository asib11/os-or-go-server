# Os or Go Server

A minimal Go HTTP server that exposes a small in-memory products API. It ships
with a few seeded products, supports creating new ones, and can fetch products
by ID.

## Requirements

- Go 1.22+
- Standard library only

## Run

```bash
go run main.go
```

The server starts on `http://localhost:8081`.

## API

### Health check

`GET /test`

```bash
curl http://localhost:8081/test
```

### List products

`GET /products`

Returns the current in-memory product list as JSON.

```bash
curl http://localhost:8081/products
```

### Get a product by ID

`GET /products/{id}`

```bash
curl http://localhost:8081/products/1
```

### Create a product

`POST /products`

Request body (JSON):

```json
{
	"title": "Product name",
	"description": "Short description",
	"price": 19.99,
	"image_url": "https://example.com/image.jpg"
}
```

```bash
curl -X POST http://localhost:8081/products \
	-H "Content-Type: application/json" \
	-d '{"title":"Product name","description":"Short description","price":19.99,"image_url":"https://example.com/image.jpg"}'
```

Notes:

- The server assigns the `id` field automatically.
- Data is stored in memory, so restarting the server resets the product list.
- Make sure the `Content-Type` header is `application/json`.
