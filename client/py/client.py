import logging
import asyncio
import grpc

import products_pb2
import products_pb2_grpc


GRPC_HOST_PORT = 'localhost:8080'

async def main():
    async with grpc.aio.insecure_channel(GRPC_HOST_PORT) as channel:
        stub = products_pb2_grpc.ProductServiceStub(channel)
        response = await stub.GetVendorProductTypes(products_pb2.ClientRequestType(vendor='google'),wait_for_ready=True,timeout=4)
    print("Python Product client received: " + response.productType)


if __name__ == '__main__':
    logging.basicConfig()
    asyncio.run(main())