echo "..........................................................."

echo '1-Try passing as arguments "--reverse example00.txt"'
go run main.go "--reverse example00.txt"
ech | cat -eo "..........................................................."
cat -e funcs/example00.txt
echo '2-Try passing as arguments --reverse=funcs/example00.txt'
go run main.go --reverse=funcs/example00.txt | cat -e
echo "..........................................................."

cat -e funcs/example01.txt
echo '3-Try passing as arguments --reverse=funcs/example01.txt'
go run main.go --reverse=funcs/example01.txt | cat -e
echo "..........................................................."

cat -e funcs/example02.txt
echo '4-Try passing as arguments --reverse=funcs/example02.txt'
go run main.go --reverse=funcs/example02.txt | cat -e
echo "..........................................................."

cat -e funcs/example03.txt
echo '5-Try passing as arguments --reverse=funcs/example03.txt'
go run main.go --reverse=funcs/example03.txt | cat -e
echo "..........................................................."

cat -e funcs/example04.txt
echo '6-Try passing as arguments --reverse=funcs/example04.txt'
go run main.go --reverse=funcs/example04.txt | cat -e
echo "..........................................................."

cat -e funcs/example05.txt
echo '7-Try passing as arguments --reverse=funcs/example05.txt'
go run main.go --reverse=funcs/example05.txt | cat -e
echo "..........................................................."

cat -e funcs/example06.txt
echo '8-Try passing as arguments --reverse=funcs/example06.txt'
go run main.go --reverse=funcs/example06.txt | cat -e
echo "..........................................................."

cat -e funcs/example07.txt
echo '9-Try passing as arguments --reverse=funcs/example07.txt'
go run main.go --reverse=funcs/example07.txt | cat -e
echo "..........................................................."

cat -e funcs/example08.txt
echo '10-Try passing as arguments --reverse=funcs/example08.txt'
go run main.go --reverse=funcs/example08.txt | cat -e
echo "..........................................................."

