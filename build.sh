PWD=`pwd`
rm $PWD/example/cg.*
go run . -pkgPath=$PWD/example
sourcetrail ./example/cg.srctrlprj
