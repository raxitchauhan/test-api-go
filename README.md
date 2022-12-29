# test-api

docker run --rm -it -d -p 4566:4566 -p 4510-4559:4510-4559 localstack/localstack

GET | http://localhost:8082/Get
GET | http://localhost:8082/Get/{uuid}

POST | http://localhost:8082/Add
{
    "name": "test",
    "author": "test-author",
    "publication": "solaris"
}

PATCH | http://localhost:8082/Update
{
    "uuid": "{uuid}",
    "name":"test"
}

--get list of tables in Dynamodb (from localstack:4566) 

GET | http://localhost:8082/get-tables

--create a table in DynamoDb with uuid and name properties

POST | http://localhost:8082/create-table
{
    "name": "test1"
}
