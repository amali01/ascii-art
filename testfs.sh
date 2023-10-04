echo 'Try passing as no arguments '
go run main.go 

echo 'Try passing as arguments "hello" standard aaa '
go run main.go "hello" standard aaa

echo 'Try passing as arguments "hello" standard | cat -e '
go run main.go "hello" standard | cat -e

echo 'Try passing as arguments "hello world" shadow | cat -e '
go run main.go "hello world" shadow | cat -e

echo 'Try passing as arguments "nice 2 meet you" thinkertoy | cat -e '
go run main.go "nice 2 meet you" thinkertoy | cat -e

echo 'Try passing as arguments "you & me" standard | cat -e '
go run main.go "you & me" standard | cat -e

echo 'Try passing as arguments "123" shadow | cat -e '
go run main.go "123" shadow | cat -e

echo 'Try passing as arguments "/(\")" thinkertoy | cat -e '
go run main.go "/(\")" thinkertoy | cat -e

echo 'Try passing as arguments 'hello There 1 to 2!''
go run main.go 'hello There 1 to 2!'

echo 'Try passing as arguments "ABCDEFGHIJKLMNOPQRSTUVWXYZ" shadow | cat -e'
go run main.go "ABCDEFGHIJKLMNOPQRSTUVWXYZ" shadow | cat -e

echo 'Try passing as arguments "\"#$%&/()*+,-./" thinkertoy | cat -e '
go run main.go "\"#$%&/()*+,-./" thinkertoy | cat -e

echo 'Try passing as arguments "It's Working" thinkertoy | cat -e '"
go run main.go "It's Working" thinkertoy | cat -e

echo 'Try passing as arguments "And We're DONE\n AMJAD" doom ''"
go run main.go "And were DONE\n \tAMJAD" doom | cat -e

