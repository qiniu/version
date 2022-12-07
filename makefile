BUILDCOMMENT=package KODO.2022-11-30-14-16-28.tar.gz
GITTAG=`git describe --tags`
DATE=`date`
NAME=libversion
LDFLAGS="-X 'github.com/qiniu/version.BuildDate=${DATE}' -X 'github.com/qiniu/version.GitTag=${GITTAG}' -X 'github.com/qiniu/version.BuildComments=${BUILDCOMMENT}' -X 'github.com/qiniu/version.Name=${NAME}'"

all:
	go install -ldflags ${LDFLAGS} -v ./...