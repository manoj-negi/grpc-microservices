# grpc-microservices`

protoc -Icalculator/proto --go_opt=module=github.com/Clement-Jean/grpc-go-course --go_out=. --go-grpc_opt=module=github.com/Clement-Jean/grpc-go-course --go-grpc_out=. calculator/proto/*.proto

protoc -Icalculator/proto --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative calculator/proto/*.proto