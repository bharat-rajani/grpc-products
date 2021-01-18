protoc -I proto --go_out=gen/proto/products/ --go-grpc_out=gen/proto/products/  ./proto/products.proto 


python -m grpc_tools.protoc -I proto --python_out=client/py/ --grpc_python_out=client/py ./proto/products.proto