### Initial
go mod init modulo  
go build  
go install  
### External Package
go get github.com/badoux/checkmail 
### Remove Unused Packages
go mod tidy  
### Variables
go is strongly typed  
go doesn´t let you import a variable and not use it  
### Data Types
always double quotes for string  
there is no char in go
in go, every data type has a value of 0  
### Functions  
more than one return per function  
### Operators  
in go there is no ternary operator
### Heritage
there is no inheritance in go
### Ponteirs
memory reference  
& put variable in pointer  
\* differentiation  
### Arrays and Slices  
if you don't specify the size of an array, it becomes a slice  
### Internal Arrays  
when slice overflows the size, go creates another internal array and doubles it  
array a list of fixed size and slice a list without fixed size  
### Maps
key value structure and not a mutable structure  
### If and Else  
start a variable directly in if  
there are no relatives in go  
### Switch  
fallthrough  
in go you don't need to use the break clause  
### Loops
go only has for  
you can't use a range in a struct  
### Advanced Functions  
functions with named return  
variatic function  
anonymous functions  
recursive functions  
defer  
panic amd recover  
closure  
functions with pointers  
init  
### Methods  
methods are functions that are associated with user-defined types  
### Interfaces  
implicitly implemented  
### Interface with generic type  
be careful not to turn int into a defective solution  
### Command line application  
receive some inputs via command line 
go mod init command_line_application  
go get github.com/urfave/cli  
go run main.go ip --host mercadolivre.com.br  
go build  
./command_line_application ip --host amazon.com.br  
### Competition  
competition != parallelism  
goroutines  
wait group  
channel  
channel with buffer  
select  
### Competition Standards  
worker pools pattern  
generator pattern  
multiplexer pattern  
### Automated Tests  
creating function for testing  
basic unit testing  
unit testing with different scenarios  
go test commands tips and coverage reports  
function with sub tests  
#### Commands
go mod init introduction  
go get golang.org/x/text/cases golang.org/x/text/language  
#### Tips
_test  
Test  
go test ./...  
go test -v  
t.Parallel()  
go test --cover  
go test --coverprofile coverage.txt  
go tool cover --func=coverage.txt  
go tool cover --html=coverage.txt  
### Advanced Test  
t.Run  
TDD  
### Additional Fundamentals  
#### JSON  
JSON concept and marshal function  
JSON function unmarshal   
#### HTTP  
HTTP protocol  
#### HTML  
creating HTML page  
rendering web page  
#### CSS  
applyng CSS  
#### JS  
introduction to javascript  
### Database  
#### Installation on linux  
sudo apt-get update  
sudo apt-get install mysql-server  
sudo mysql_secure_installation utility  
sudo systemctl start mysql  
sudo systemctl enable mysql  
sudo mysql -u root -p  
show databases;  
create database golang;  
use golang;  
CREATE TABLE users (id INT auto_increment primary key, name VARCHAR(50) NOT NULL, email VARCHAR(50) NOT NULL) ENGINE=INNODB;  
CREATE USER 'golang'@'localhost' IDENTIFIED BY 'golang';  
GRANT ALL PRIVILEGES ON golang.* TO 'golang'@'localhost';  
exit;  
mysql -u golang -p  
#### Installation on windows or mac  
dev.mysql.com  
MySQL Installer for Windows  
MySQL community server  
Workbench  
macOS select macOS  
#### Connecting to the database  
go mod init database  
go get github.com/go-sql-driver/mysql  
implicit import "_"  
SQL is a dialect that is used in several banks  
##### Running on WSL 2 with windows database  
cat /etc/resolv.conf | grep nameserver
GRANT ALL PRIVILEGES ON *.* TO 'golang'@'%' IDENTIFIED BY 'golang';  
FLUSH PRIVILEGES;  
sudo apt update  
sudo apt install mysql-client-core-8.0  
#### Installing Postman  
https://www.postman.com/  
#### Creating structure to CRUD  
go mod init crud  
CRUD = Create, Read, Update, Delete  
go get github.com/gorilla/mux  
#### Insert Data
router.HandleFunc("/users", server.CreateUser).Methods("POST")  
#### Get Data  
router.HandleFunc("/users", server.GetUsers).Methods("GET")  
#### Get Data by ID  
router.HandleFunc("/users/{id}", server.GetUser).Methods("GET")  
#### Update data  
router.HandleFunc("/users/{id}", server.UpdateUser).Methods("PUT")  
#### Delete data  
router.HandleFunc("/users/{id}", server.DeleteUser).Methods("DELETE")  
