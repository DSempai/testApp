## How to Run
To start the application, run the following command in your terminal:
```
docker-compose up --build -d
```
To start the monitoring services, run the following command in your terminal:
```
docker-compose -f monitoring/docker-compose.monitoring.yml up --build -d
```

## Monitoring Access
- Grafana: http://localhost:3000 (default credentials: admin/admin)
- Prometheus: http://localhost:9090

Endpoints
1. Get User
Retrieves user information by ID.
Endpoint: `GET /users`
Query Parameters:
`id` (required): The ID of the user to retrieve
Response:

```
{
    "id": "string",
    "name": "string"
}
```
Error Response:
```
{
    "error": "error message"
}
```
Example curl:
```
curl -X GET "http://localhost:8080/users?id=user123"
```
2. Create User
Creates a new user or updates existing user.
Endpoint: `POST /users/create`
Request Body:
```
{
    "id": "string",
    "name": "string"
}
```
Response:
```
{
    "id": "string",
    "name": "string"
}
```
Error Response:
```
{
    "error": "error message"
}
```
Example curl:
```
curl -X POST "http://localhost:8080/users/create" \
     -H "Content-Type: application/json" \
     -d '{"id": "user123", "name": "John Doe"}'
```