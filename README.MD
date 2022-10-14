# Go Public API Todo
This is a Todo list public API written on Golang. You can rebuild it and save it for your own project, or you can use it directly for your Frontend project. Direct API endpoint: https://todoapi.rozik.in
## How to use it
### - Use the public API
1. Read API documentation on api.json. The API endpoint is https://todoapi.rozik.in
1. Autenthicate yourself with simple authentification https://todoapi.rozik.in/login with method POST and form data username: admin and password: admin
1. Save the token.
1. Do the request with the token as an Authorization Bearer
1. If you want to try the graphQL, you can visit: https://todoapi.rozik.in/gql_play with query endpoint https://todoapi.rozik.in/gql_query

### - Rebuild the API
1. Pull or copy the repository
1. Create your own .env file from the example.env file. (don't use username root)
1. Do docker-compose "docker-compose up"
1. Access your API from http://localhost:8200
1. Read API documentation on api.json
1. Autenthicate yourself with simple authentification http://localhost:8200/login with method POST and form data username: admin and password: admin
1. Save the token.
1. Do the request with the token as an Authorization Bearer
