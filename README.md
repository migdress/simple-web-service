# simple web service
This is a serverless project

# architecture

![arch](res/v1.png?raw=true)

## Create user
* This microservice is a lambda function exposed as http with APIGateway whose responsibility is to create a record in a dynamodb table
* The dynamodb table must exist