echo "Testing the Ascii-Art..."

echo "Reverse Testing against input in testfiles..."
echo "----------------------------------"
for file in ./test/*; do
echo "----------------------------------"
    go run main.go --reverse=$file
    if [ $? -ne 0 ]; then
        echo "Error running the program for file: $file"
    fi
    echo "----------------------------------"

done

echo "All tests completed."

# echo "Ascii-Art Testing against input in testfiles..."
# echo "----------------------------------"
# bash testasciiArt.sh
# echo "All tests completed."
##############################
# echo "Color Testing against input in testfiles..."
# echo "----------------------------------"
# bash testcolor.sh
# echo "All tests completed."
###############################
# echo "banner Testing against input in testfiles..."
# echo "----------------------------------"
# bash testasciiArt.sh
# echo "All tests completed."
###############################
# echo "Justify Testing against input in testfiles..."
# echo "----------------------------------"
# bash testjustify.sh
# echo "All tests completed."
###############################
# echo "Output Testing against input in testfiles..."
# echo "----------------------------------"
# bash testoutput.sh
# echo "All tests completed."