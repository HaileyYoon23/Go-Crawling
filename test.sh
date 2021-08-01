cat ImageNameList.txt | while read line
do
    test="$line\n"
    echo $test
done
