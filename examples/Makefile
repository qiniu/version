BUILDCOMMENT=package KODO.2022-11-30-14-16-28.tar.gz
GITTAG=`git describe --tags`
DATE=`date`
NAME=libversion
LDFLAGS="-X 'github.com/qiniu/version/v2.BuildDate=${DATE}' -X 'github.com/qiniu/version/v2.GitTag=${GITTAG}' -X 'github.com/qiniu/version/v2.BuildComments=${BUILDCOMMENT}' -X 'github.com/qiniu/version/v2.Name=${NAME}'"

all:
	go build -ldflags ${LDFLAGS} -v ./main.go
