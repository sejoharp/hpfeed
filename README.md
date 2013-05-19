# Introduction
This app grabs new posts in hamburg privateers forum from http://www.kickern-hamburg.de/phpBB2/ and provides them as an rss 2.0 feed.

# The way it works
* logs in to the forum
* grabs the forum overview from hamburg privateers forum
* logs out from the forum
* persists all new postings
* provides an rss 2.0 feed with all postings

# Setup
* install Mongodb/couchdb and create a db
* clone this repo
* call "go get" to download all dependencies 
* customise hpfeed.conf and copy it to the destination folder
** deployment and runtime management:
*** via hpfeed.sh script:
**** customise scripts/hpfeed.sh and copy the hpfeed.sh to the start/stop folder
**** customise scripts/deploy.sh and execute it.
*** via daemontools:
**** setup a new service with a run cmd like this:
	./hpfeed -config=/PATH/TO/CONFIG/hpfeed.conf 2>&1
**** customise scripts/deploy_daemontools.sh and execute it.   

# Dependencies
* golang
* goquery from https://github.com/puerkitobio/goquery
* goconf from https://code.google.com/p/goconf/
* mgo from labix.org/v2/mgo
* dsallings-couch-go from code.google.com/p/dsallings-couch-go
* moverss from github.com/baliw/moverss
* account with permission to read the privateers forum
* mongodb or couchdb

# My setup
* A raspberry pi with raspbian linux
* couchdb

