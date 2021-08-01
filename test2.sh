#!bin/bash
go run main.go
cat number.txt | while read line
do
    test="$line"
    echo "sh $test"
done
go run clear.go

