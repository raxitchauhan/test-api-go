# test-api

http://localhost:8082/Get
http://localhost:8082/Get/{uuid}

http://localhost:8082/Add
{
    "name": "test",
    "author": "test-author",
    "publication": "solaris"
}

http://localhost:8082/Update
{
    "uuid": "{uuid}",
    "name":"test"
}

http://localhost:8082/get-tables --get list of tables in Dynamodb (from localstack:4566)

http://localhost:8082/create-table --create a table in DynamoDb with uuid and name properties
{
    "name": "test1"
}
