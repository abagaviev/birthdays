## Simple REST API service for saving birthdays

using Gin and Gorn

How to use: 
```$ go run main.go```

### endpoints
Create birthday
```POST /birtdays```
```body {
    name - string required
    surname - string required
    date_of_birth - string required, format "YYYY-MM-DD"
}```

Get all birthdays
```GET /birtdays```

Get birthday by ID
```GET /birtdays/{id}```

Update birthday 
```PATCH /birtdays/{id}```
```body {
    name - string 
    surname - string 
    date_of_birth - string, format "YYYY-MM-DD"
}```

Delete birthday
```DELETE /birtdays/{id}```



