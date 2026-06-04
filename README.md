# Os or Go Server

A minimal Go HTTP server that exposes a simple products API. It supports listing
products and creating a new product via JSON.

## Requirements

- Go 1.22+ (module uses the standard library only)

## Run

```bash
go run .
```

Server starts at `http://localhost:8081`.

## API

### List products

`GET /products`

```bash
curl http://localhost:8081/products
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
- Make sure the `Content-Type` header is `application/json`.
