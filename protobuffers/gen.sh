#!/bin/bash
protoc seguridad/usuario.proto --go_out=plugins=grpc:.
protoc zuvap/pago.proto --go_out=plugins=grpc:.