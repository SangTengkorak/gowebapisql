This project is based on this tutorial: https://medium.com/@abhishekkushwaha11998/go-building-api-with-mysql-and-gin-d5a8ef70348f

Few operational tasks:

++Create people Send Request
curl -X POST http://localhost:8090/peoples/create\
 --header "Content-Type: application/json"\
 -d '{"ID":5, "first_name":"Jackie", "last_name":"chan", "birth_year":1954, "email":"jackie_chan@zmail.com"}'

++Update people Send Request
curl -X POST http://localhost:8090/peoples/upd/5\
 --header "Content-Type: application/json"\
 -d '{"ID":5, "first_name":"Master Jackie", "last_name":"chan", "birth_year":1954, "email":"jackie_chan@zmail.com"}'

++++Get all peoples Send Request
curl -X GET http://localhost:8090/peoples

++++Get specific people by ID Send Request
curl -X GET http://localhost:8090/peoples/get/5

++Delete specific people by ID Send Request
curl -X POST http://localhost:8090/peoples/del/5
