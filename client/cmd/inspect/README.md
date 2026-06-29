MikroTik commands inspection tool
=================================

This tool queries MikroTik resources, including commands and arguments, and outputs result as a structured JSON.

## MikroTik client resource

To query resource or command, just run:
```sh
$ go run ./cmd/inspect -root "/ip/dns/static"
```
By default, the depth of inspection is set to `1` which gives output of the current command path `/ip/dns/static` and 1 level below (chidlren).

To get all children recursively (may be slow) pass `-depth -1` to the tool invocation.


Sample output is similar to this (might be different for your remote RouterOS setup):
```json
{
    "Self": "/ip/dns/static",
    "Name": "static",
    "Type": "dir",
    "Children": [
        {
            "Self": "/ip/dns/static/add",
            "Name": "add",
            "Type": "cmd",
            "Children": null,
            "Arguments": null
        },
        {
            "Self": "/ip/dns/static/comment",
            "Name": "comment",
            "Type": "cmd",
            "Children": null,
            "Arguments": null
        },
        {
            "Self": "/ip/dns/static/disable",
            "Name": "disable",
            "Type": "cmd",
            "Children": null,
            "Arguments": null
        },
        {
            "Self": "/ip/dns/static/edit",
            "Name": "edit",
            "Type": "cmd",
            "Children": null,
            "Arguments": null
        },
        {
            "Self": "/ip/dns/static/enable",
            "Name": "enable",
            "Type": "cmd",
            "Children": null,
            "Arguments": null
        },
        {
            "Self": "/ip/dns/static/export",
            "Name": "export",
            "Type": "cmd",
            "Children": null,
            "Arguments": null
        },
        {
            "Self": "/ip/dns/static/find",
            "Name": "find",
            "Type": "cmd",
            "Children": null,
            "Arguments": null
        },
        {
            "Self": "/ip/dns/static/get",
            "Name": "get",
            "Type": "cmd",
            "Children": null,
            "Arguments": null
        },
        {
            "Self": "/ip/dns/static/move",
            "Name": "move",
            "Type": "cmd",
            "Children": null,
            "Arguments": null
        },
        {
            "Self": "/ip/dns/static/print",
            "Name": "print",
            "Type": "cmd",
            "Children": null,
            "Arguments": null
        },
        {
            "Self": "/ip/dns/static/remove",
            "Name": "remove",
            "Type": "cmd",
            "Children": null,
            "Arguments": null
        },
        {
            "Self": "/ip/dns/static/reset",
            "Name": "reset",
            "Type": "cmd",
            "Children": null,
            "Arguments": null
        },
        {
            "Self": "/ip/dns/static/set",
            "Name": "set",
            "Type": "cmd",
            "Children": null,
            "Arguments": null
        }
    ],
    "Arguments": null
}
```
