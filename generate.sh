#!/bin/bash

protoc colorspb/colors.proto --go_out=plugins=grpc:.