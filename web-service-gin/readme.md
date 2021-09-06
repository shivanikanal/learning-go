This module spawns an http server and adds some simple CRUD apis to it to fetch albums data

Run the server using:
`go run .`

The server runs on port 8081

APIs available
`GET /albums`
`GET /albums/:id`
`POST /albums --data {"id": "4","title": "The Modern Sound of Betty Carter","artist": "Betty Carter","price": 49.99}`