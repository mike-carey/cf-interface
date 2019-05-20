#!/usr/bin/env bash

###
# This file modifies the generated file from ifacemaker to import cfclient and modify the classes that needed to be referenced  from that package.
##

file=${1:cf_interface.go}

sed -i.bak 's|\(import (\)|\1\
'$'\t"github.com/cloudfoundry-community/go-cfclient"|g' $file
sed -i.bak '/^[type|/\/\/]/! s/\([][( \*]\)\([A-Z]\)/\1cfclient.\2/g' $file
