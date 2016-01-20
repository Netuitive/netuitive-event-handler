Netuitive-event-handler
=======================
Netuitive Event Handler is a CLI for posting external events to Netuitive.
It is designed to be used from Nagios, Icinga, Sensu, and other monitoring systems.


Build
-----
* Get a functioning [Go](https://golang.org) environment
* Put this code in $GOPATH/src/github.com/netuitive/netuitive-event-handler
* Run make


Installation
------------
* download from [our repo](http://repos.app.netuitive.com/cli-agent/index.html) and copy it to /bin/netuitive-event-handler
* chmod 755 the thing
* configure the /etc/netuitive/netuitive-event-handler.yaml

/etc/netuitive/netuitive-event-handler.yaml example
-------------------------------------------
    apikey: DEMOab681D46bf5616dba8337c85DEMO
    url: "https://demoapi.app.netuitive.com/ingest/events"


Nagios Configuration
--------------------
* Be one with the Nagios documentation.
* Create Nagios commands definitions as follows:

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

Sensu Configuration
--------------------
* Be one with the Sensu documentation.
* Create Sensu handler as followings:

### Sensu Handler
    {
      "handlers": {
        "netuitive-event-handler": {
          "type": "pipe",
          "command": "/bin/netuitive-event-handlernetuitive-event-handler-linux stdin -s Sensu",
          "severities": [
            "critical",
            "ok"
          ]
        }
      }
    }


Command Line Options
--------------------

    Usage:
      netuitive-event-handler [flags]
      netuitive-event-handler [command]

    Available Commands:
      stdin       Post events to Netuitive from the stdin pipe
      version     print the version

    Flags:
      -a, --apikey="":  API Key if not otherwise specified (optional)
      -c, --config="": config file (default is /etc/netuitive/netuitive-event-handler.yaml)
      -d, --debug[=false]: enable debug
      -e, --element="": Element FQN for the event
      -l, --level="": Level of the event
      -m, --message="": Message text of the event
      -s, --source="netuitive-event-handler": Source of the event (optional)
          --tags="": Tags for the event (optional) Example: tag1:value1,tag2:value2
      -t, --title="": Title of the event
      -u, --url="https://api.app.netuitive.com/ingest/events":  API URL if not otherwise specified (optional)

    Use "netuitive-event-handler [command] --help" for more information about a command.



    Usage:
      netuitive-event-handler stdin [flags]

    Global Flags:
      -a, --apikey="":  API Key if not otherwise specified (optional)
      -c, --config="": config file (default is /etc/netuitive/netuitive-event-handler.yaml)
      -d, --debug[=false]: enable debug
      -s, --source="netuitive-event-handler": Source of the event (optional)
      -u, --url="https://api.app.netuitive.com/ingest/events":  API URL if not otherwise specified (optional)
