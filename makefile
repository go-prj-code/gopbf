all :

buildGoCode :
	protoc --go_out=. ./gp3/user.proto