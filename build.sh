#!/usr/bin/env bash

echo "Building aerospike_exporter binaries"
echo ""
echo $GO_LDFLAGS

gox -rebuild --osarch="darwin/amd64"  -ldflags "$GO_LDFLAGS" -output "dist/aerospike_exporter" && cd dist && tar -cvzf aerospike_exporter-$CIRCLE_TAG.darwin-amd64.tar.gz aerospike_exporter && rm aerospike_exporter && cd ..
gox -rebuild --osarch="darwin/386"    -ldflags "$GO_LDFLAGS" -output "dist/aerospike_exporter" && cd dist && tar -cvzf aerospike_exporter-$CIRCLE_TAG.darwin-386.tar.gz   aerospike_exporter && rm aerospike_exporter && cd ..
gox -rebuild --osarch="linux/amd64"   -ldflags "$GO_LDFLAGS" -output "dist/aerospike_exporter" && cd dist && tar -cvzf aerospike_exporter-$CIRCLE_TAG.linux-amd64.tar.gz  aerospike_exporter && rm aerospike_exporter && cd ..
gox -rebuild --osarch="linux/386"     -ldflags "$GO_LDFLAGS" -output "dist/aerospike_exporter" && cd dist && tar -cvzf aerospike_exporter-$CIRCLE_TAG.linux-386.tar.gz    aerospike_exporter && rm aerospike_exporter && cd ..
gox -rebuild --osarch="netbsd/amd64"  -ldflags "$GO_LDFLAGS" -output "dist/aerospike_exporter" && cd dist && tar -cvzf aerospike_exporter-$CIRCLE_TAG.netbsd-amd64.tar.gz aerospike_exporter && rm aerospike_exporter && cd ..
gox -rebuild --osarch="netbsd/386"    -ldflags "$GO_LDFLAGS" -output "dist/aerospike_exporter" && cd dist && tar -cvzf aerospike_exporter-$CIRCLE_TAG.netbsd-386.tar.gz   aerospike_exporter && rm aerospike_exporter && cd ..
gox -rebuild --osarch="windows/amd64" -ldflags "$GO_LDFLAGS" -output "dist/aerospike_exporter" && cd dist && zip -9    aerospike_exporter-$CIRCLE_TAG.windows-amd64.zip   aerospike_exporter.exe && rm aerospike_exporter.exe && cd ..
gox -rebuild --osarch="windows/386"   -ldflags "$GO_LDFLAGS" -output "dist/aerospike_exporter" && cd dist && zip -9    aerospike_exporter-$CIRCLE_TAG.windows-386.zip     aerospike_exporter.exe && rm aerospike_exporter.exe && cd ..

echo "Build Successful"

echo "Uploading to Github"
ghr -t $GITHUB_TOKEN -u $CIRCLE_PROJECT_USERNAME -r $CIRCLE_PROJECT_REPONAME -delete $CIRCLE_TAG dist/

echo "Upload Successful"