server:
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative proto/greet.proto
calc:
	protoc -Icalculator/proto --go_out=calculator/proto --go_opt=paths=source_relative --go-grpc_out=calculator/proto --go-grpc_opt=paths=source_relative calculator/proto/*.proto

	