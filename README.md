.. netuitive-event-handler
.. README.rst

Netuitive-event-handler
=======================

*Netuitive Event Handler is a CLI for posting external events
to Netuitive. It is designed to be used from Nagios, Icinga, Sensu, and other monitoring systems.


Installation
------------
* download from someplace to someplace useful
* chmod 755 the thing
* configure the /etc/netuitive/netuitive-event-handler.yaml

/etc/netuitive/netuitive-event-handler.yaml example
-------------------------------------------
    apikey: DEMOab681D46bf5616dba8337c85DEMO
    url: "https://api.app.netuitive.com/ingest/events"


Nagios Configuration
--------------------
* Be one with the Nagios documentation.
* Create Nagios commands definitions as followings:

### Host Notification
    define command{
        command_name    notify-host-by-netuitive-event
        command_line    /bin/netuitive-event-handler -s Nagios -e "$HOSTALIAS$" -t "Host $HOSTALIAS$ is $HOSTSTATE$" -l "$HOSTSTATE$"  -m "Host $HOSTALIAS$ is $HOSTSTATE$ - Info: $HOSTOUTPUT$"
    }

### Service Notification
    define command{
        command_name    notify-service-by-netuitive-event
        command_line    /bin/netuitive-event-handler -s Nagios -e "$HOSTALIAS$" -t "Service $SERVICEDESC$ is $SERVICESTATE$" -l "$SERVICESTATE$"  -m "Service $SERVICEDESC$ is $SERVICESTATE$ - Info: $SERVICEOUTPUT$"
    }

