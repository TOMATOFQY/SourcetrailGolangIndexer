ABSDIR=`pwd`/example
if [ "$1" != "" ]
then
	ABSDIR=$1
fi

echo "Indexing" $ABSDIR "..."

export LD_LIBRARY_PATH=./bindings_golang:$LD_LIBRARY_PATH
go run ./indexer/ -pkgPath=$ABSDIR

sourcetrail $ABSDIR/cg.srctrlprj