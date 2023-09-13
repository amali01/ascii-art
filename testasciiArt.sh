echo 'Try passing as no arguments '
go run main.go 

echo 'Try passing as empty argument  "" '
go run . "" | cat -e

echo 'Try passing as arguments "hello"'
go run main.go "hello" | cat -e

echo 'Try passing as arguments "HELLO"'
go run main.go "HELLO" | cat -e

echo 'Try passing as arguments "HeLlo HuMaN"'
go run main.go "HeLlo HuMaN" | cat -e

echo 'Try passing as arguments "1Hello 2There"'
go run main.go "1Hello 2There" | cat -e

echo 'Try passing as arguments "Hello\nThere"'
go run main.go "Hello\nThere" | cat -e

echo 'Try passing as arguments "Hello\n\nThere"'
go run main.go "Hello\n\nThere" | cat -e

echo 'Try passing as arguments "{Hello & There #}"'
go run main.go "{Hello & There #}" | cat -e

echo 'Try passing as arguments 'hello There 1 to 2!''
go run main.go 'hello There 1 to 2!' | cat -e

echo 'Try passing as arguments "MaD3IrA&LiSboN"'
go run main.go "MaD3IrA&LiSboN" | cat -e

echo 'Try passing as arguments "1a\"#FdwHywR&/()="'
go run main.go "1a\"#FdwHywR&/()=" | cat -e

echo 'Try passing as arguments "{|}~"'
go run main.go "{|}~" | cat -e

echo 'Try passing as arguments [\]^_ 'a""''
go run main.go "[\]^_ 'a" | cat -e

echo 'Try passing as arguments "RGB"'
go run main.go "RGB" | cat -e

echo 'Try passing as arguments ":;<=>?@"'
go run main.go ":;<=>?@" | cat -e

echo 'Try passing as arguments "\!" #$%&'"'"'()*+,-./" '
go run main.go '\!" #$%&'"'"'()*+,-./' | cat -e

echo 'Try passing as arguments "ABCDEFGHIJKLMNOPQRSTUVWXYZ" '
go run main.go "ABCDEFGHIJKLMNOPQRSTUVWXYZ" | cat -e

echo 'Try passing as arguments "abcdefghijklmnopqrstuvwxyz"'
go run main.go "abcdefghijklmnopqrstuvwxyz" | cat -e
