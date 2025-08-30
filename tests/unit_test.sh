#!/bin/sh

docker run --rm -d \
	--network=host \
	fib-service:latest | tee result.log | tail -20

sleep 3

curl -s 'http://localhost:8080/fibonacci?number=78' | jq -c '.sequence' | grep -E '^\[0,1,1,2,3,5,8,13,21,34,55\]$'
if [ $? -eq 0 ]; then 
	echo "Test passed" 
	docker stop $(docker ps -a -q --filter ancestor=fib-service:latest --format="{{.ID}}")
else 
	echo "Fibonacci service did not return the right sequence" 
	docker stop $(docker ps -a -q --filter ancestor=fib-service:latest --format="{{.ID}}")
	exit 1 
fi

