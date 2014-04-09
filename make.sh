# Installs the 3rd party libraries for the go project
# Please not that the mgo libraries will require bazaar to be installed first
# Usage: chmod +x make.sh
#        ./make.sh
# Please also be aware that this will set the GOPATH to the current directory.
# If this is not desired, then I suggest installing the packages manually instead.
export GOPATH="$(pwd)"
go get code.google.com/p/gcfg
go get labix.org/v2/mgo