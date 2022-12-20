# msgteams

![](https://img.shields.io/badge/test-passing-green)
![](https://img.shields.io/badge/coverage-95%25-orange)
![](https://img.shields.io/badge/version-0.1.0-blue)

A command line tool for sending simple messages to Microsoft Teams channels.

## Overview


`msgteams` allows you to send messages to a Microsoft Teams channel from the command line. It is a convenient way to send messages as part of automated processes or scripts.

To use `msgteams`, you will need to generate a URL for your Microsoft Teams channel. This can be done by following these steps:

- Open your Microsoft Teams channel.
- Click on the three dots in the top right corner of the channel.
- From the menu that appears, select "Connectors".
- Scroll down until you see the "Incoming Webhook" connector. If it is not appearing, you may need to contact your Teams administrator to have it added to your channel.
- Click on the "Configure" or "Add" button for the Incoming Webhook connector.
- Provide a name for the incoming webhook and click "Create".
- Click "Save" to save the URL provided.

The URL provided is a unique address that you will use to send messages to your Microsoft Teams channel. You will need to save the URL so that you can use it with the command.

Please note that Microsoft may change the process for creating incoming webhooks at any time. If the steps above do not work, you can find the most up-to-date instructions on the Microsoft Teams documentation page

## Usage

```
msgteams [flags]
```
## Examples

```
msgteams --url <your URL> --message "your message"

or

msgteams -u <your URL> -m "your message"
```

## Options

```
-h, --help             display help and exit
-m, --message string   message to be sent to teams channel
-u, --url string       url to Microsoft Teams channel
-v, --version          output version information and exit
```

## Troubleshooting

If you encounter any issues when using msgteams, try the following:

- Double-check that you have entered the correct URL for your Microsoft Teams channel
- Make sure that you have followed the steps to configure an incoming webhook correctly
- Ensure that you have the necessary permissions to send messages to the channel
- If you are using msgteams as part of a script or automated process, make sure that it is being executed correctly

## Limitations

- msgteams can only send simple text messages to Microsoft Teams channels
- There may be limitations on the length of messages that can be sent using msgteams

## Best Practices

- Keep messages concise and to the point
- Use proper formatting to make messages easier to read
- Test msgteams with a small message before sending longer or more important messages
