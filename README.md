# PSSS
PSSS is a set of Golang implemented API and command line utility reading system, socket, process information.

## Status
psss package API only supports Linux platform. Socket information reading API is implemented by means of sending Linux netlink and scanning the proc file system. Process and system information reading API is implemented by scanning the proc file system.

## Installation
    go get github.com/buck119br/psss

## cmd
A CLI utility, whose output format and argument are like Linux route utility ss.

## topo
topo package provides API reading relations between local services.