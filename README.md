# Introduction
This app grabs new posts in hamburg privateers forum from http://www.kickern-hamburg.de/phpBB2/ and provides them as an rss feed.

# The way it works
* check whether the macos is at home. The script continues only in this case.
* wakes the remote time machine up if it is offline
* saves the data if the time machine is available
* shuts down the remote time machine if it was offline

# Setup
* clone this repo
* instantiate the backup class and call the method startBackupProcess in a starterfile. Example:  

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
* mongodb
* Genghis from http://genghisapp.com/ to manage the db