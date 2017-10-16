Netuitive Event Handler
=======================

The Netuitive Event Handler is a command line interface (CLI) for posting external events to [Netuitive](https://www.netuitive.com). It is designed to work with Nagios, Icinga, Sensu, and other monitoring systems.

For more information on Netuitive external events, see our [help docs](https://help.netuitive.com/Content/Events/external_events_intro.htm), or contact Netuitive support at [support@netuitive.com](mailto:support@netuitive.com).

Installing and Building the Netuitive Event Handler
---------------------------------------------------

### Build

1. Setup a functioning [Go](https://golang.org) environment.
1. Pass the following code into your Go environment:

        $GOPATH/src/github.com/netuitive/netuitive-event-handler

1. Run `make`.

### Install

1. Download the event handler [here](https://github.com/Netuitive/netuitive-event-handler/releases) and copy it to `/bin/netuitive-event-handler`.
1. Change permissions on the downloaded file to `755`:

        chmod 755 [the-downloaded-file]

1. Configure the netuitive-event-handler.yaml file (defaul location is `/etc/netuitive/`) as desired. Example:

        apikey: "yourapikey"
        url: "https://api.app.netuitive.com/ingest/events"

Configuration
--------------

### Nagios

Create Nagios command definitions as follows:

#### Host Notification
    
    define command{
            command_name    notify-host-by-netuitive-event
            command_line    /bin/netuitive-event-handler -s Nagios -e "$HOSTALIAS$" -t "Host $HOSTALIAS$ is $HOSTSTATE$" -l "$HOSTSTATE$"  -m "Host $HOSTALIAS$ is $HOSTSTATE$ - Info: $HOSTOUTPUT$"
        }
        
#### Service Notification
    
    define command{
        command_name    notify-service-by-netuitive-event
        command_line    /bin/netuitive-event-handler -s Nagios -e "$HOSTALIAS$" -t "Service $SERVICEDESC$ is $SERVICESTATE$" -l "$SERVICESTATE$"  -m "Service $SERVICEDESC$ is $SERVICESTATE$ - Info: $SERVICEOUTPUT$"
    }

### Sensu Configuration

Create Sensu handler as follows:

    {
      "handlers": {
        "netuitive-event-handler": {
          "type": "pipe",
          "command": "/bin/netuitive-event-handler stdin -s Sensu",
          "severities": [
            "critical",
            "ok"
          ]
        }
      }
    }

Additional Information
-----------------------

### Command Line Options

`netuitive-event-handler [flags] [command]`

Commands available:

| Command Name | Description |
|--------------|-------------|
| stdin | Post events to Netuitive from the stdin pipe. |
| version | Print the version. |

Use `netuitive-event-handler [command] --help` for more information about a command.

Flags available:

| Flag Name | Description | Global? |
|-----------|-------------|---------|
| -a, --apikey="" | API key if not specified otherwise (optional). | Y |
| -c, --config="" | Configuration file location. The default is `/etc/netuitive/netuitive-event-handler.yaml`. | Y |
| -d, --debug="" | Enable debug. The default is false. | Y |
| -e, --element="" | Element fully qualified name (FQN) for the event. | N |
| -l, --level="" | Level of the event (INFO, WARN, CRIT). | N |
| -m, --message="" | Message text for the event. | N |
| -s, --source="netuitive-event-handler" | Source of the event (optional). | Y |
| -t, --title="" | Title of the event. | N |
| -u, --url="https://api.app.netuitive.com/ingest/events" | API URL if not specified otherwise (optional). | Y |
