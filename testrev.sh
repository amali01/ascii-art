echo "..........................................................."

echo '1-Try passing as arguments "--reverse example00.txt"'
go run main.go "--reverse example00.txt"
ech | cat -eo "..........................................................."
cat -e test/example00.txt
echo '2-Try passing as arguments --reverse=test/example00.txt'
go run main.go --reverse=test/example00.txt | cat -e
echo "..........................................................."

cat -e test/example01.txt
echo '3-Try passing as arguments --reverse=test/example01.txt'
go run main.go --reverse=test/example01.txt | cat -e
echo "..........................................................."

cat -e test/example02.txt
echo '4-Try passing as arguments --reverse=test/example02.txt'
go run main.go --reverse=test/example02.txt | cat -e
echo "..........................................................."

cat -e test/example03.txt
echo '5-Try passing as arguments --reverse=test/example03.txt'
go run main.go --reverse=test/example03.txt | cat -e
echo "..........................................................."

cat -e test/example04.txt
echo '6-Try passing as arguments --reverse=test/example04.txt'
go run main.go --reverse=test/example04.txt | cat -e
echo "..........................................................."

cat -e test/example05.txt
echo '7-Try passing as arguments --reverse=test/example05.txt'
go run main.go --reverse=test/example05.txt | cat -e
echo "..........................................................."

cat -e test/example06.txt
echo '8-Try passing as arguments --reverse=test/example06.txt'
go run main.go --reverse=test/example06.txt | cat -e
echo "..........................................................."

cat -e test/example07.txt
echo '9-Try passing as arguments --reverse=test/example07.txt'
go run main.go --reverse=test/example07.txt | cat -e
echo "..........................................................."

cat -e test/example08.txt
echo '10-Try passing as arguments --reverse=test/example08.txt'
go run main.go --reverse=test/example08.txt | cat -e
echo "..........................................................."

