#!/bin/bash

set -e
#Use the registData.json to register in EdgeXFoundry,this file contains some necessary information for registration.
registData=`cat ./registData.json`

#EdgeXFoundry register client url
url=http://edgex-export-client:48071/api/v1/registration

curl -X POST $url -H 'Content-Type: application/json' -H 'cache-control: no-cache' -d "${registData}"

#Startup htservice program
exec $1
