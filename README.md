# web-math-go

Sample web service for storing and calculating math equations.
For now only simple cases of math equations are supported, for example

    y=x+a
    y=x+a-b
    y=x+5+a, etc

Currently supported functionalities:

    * Basic CRUD operations for storing/removing equations inside Notion database
    * Calculate equation on a range from -5 to 5

## Examples

Also available inside `request_examples` folder

    # create
    curl -X POST localhost:8080/equations -H "content-type: application/json" --data '{"term": "y=x+a+b"}'

    # list all
    curl localhost:8080/equations

    # delete
    curl -X DELETE localhost:8080/equations -H "content-type: application/json" --data '{"term": "y=x+a+b"}'

    # Get list of points for equation on a range from -5 to 5
    curl -X POST localhost:8080/calculate -H "content-type: application/json" --data '{"term": "y=x-a", "args": {"a": 3}}'
