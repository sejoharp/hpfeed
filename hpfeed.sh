#!/bin/sh

### BEGIN INIT INFO
# Provides:       hpfeed
# Required-Start: $local_fs $syslog
# Required-Stop:  $local_fs $syslog
# Default-Start:  2 3 4 5
# Default-Stop:   0 1 6
# Short-Description: hpfeed
### END INIT INFO

CONFDIR="/home/hpnews"
NAME="hpfeed"

USER=hpnews
PID="/tmp/hpfeed.pid"
RETVAL=0

# source function library
. /lib/lsb/init-functions

start()
{
    echo "Starting $NAME."
    cd $CONFDIR/;
    su -c "./hpfeed >> hpfeed.log 2>&1 &" $USER && echo "OK" || echo "failed";
    echo $! > $PID
}

stop()
{
    echo "Stopping $NAME"
    kill -QUIT `cat $PID` && echo "OK" || echo "failed";
}

reload()
{
    echo "Reloading $NAME:"
    if [ -f $PID ]
    then kill -HUP `cat $PID` && echo "OK" || echo "failed";
    fi
}

case "$1" in
    start)
        start
        ;;
    stop)
        stop
        ;;
    restart)
        reload
        ;;
    reload)
        reload
        ;;
    force-reload)
        stop && start
        ;;
    *)
        echo $"Usage: $0 {start|stop|restart}"
        RETVAL=1
esac
exit $RETVAL
