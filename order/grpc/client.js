const protoLoader = require('@grpc/proto-loader');
const grpcLibrary = require('@grpc/grpc-js');

const options = {
    keepCase: true,
    longs: String,
    enums: String,
    defaults: true,
    oneofs: true
}
const packageDefinition = protoLoader.loadSync('./grpc/proto/createOrder.proto', options);
const packageObject = grpcLibrary.loadPackageDefinition(packageDefinition);

const order_proto = packageObject.order;

const client = new order_proto.CreateOrderService(
    process.env.GRPC_SERVER,
    grpcLibrary.credentials.createInsecure()
);

module.exports = client