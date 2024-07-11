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
* differentiation  
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
