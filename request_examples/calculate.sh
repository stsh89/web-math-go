curl -X POST localhost:8080/calculate \
    -H "content-type: application/json" \
    --data '{"term": "y=x-a", "args": {"a": 3}}'
