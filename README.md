# reacji-github-issue

Slack-bot create GitHub issues with a specific Reacji.

# Setup Guide

## Create a Socket-Mode Slack application and install to your workspace.

- Set up your Slack App according to [this guide](https://api.slack.com/apis/connections/socket#setup).
- Get the app-level token(`xapp-blablabla...` ).
- Install the app to your workspace and get the bot token(`xoxb-blablabla...` ).

## Create a GitHub app and install to your organisation.

- Set up your GitHub App according
  to [this guide](https://docs.github.com/en/developers/apps/building-github-apps/creating-a-github-app)
- Get your App ID from <https://github.com/settings/apps> and click `Edit` button on your app see `App Id`.
- [Generate a private key](https://docs.github.com/en/developers/apps/building-github-apps/authenticating-with-github-apps#generating-a-private-key)
  and get base64 encoded value from `.pem` file.

```shell
base64 ${your-github-app.pem}
```

- Install the app to your organisation and get `installation id`.
    - Go to <https://github.com/settings/installations/> and click `Configure` button on your app and see last section
      of the URL.(`https://github.com/settings/installations/{your_installation_id}`)

## Edit reacji.toml.

- Edit `reacjira.toml` for your setting.

## Set secrets as environment variables and run.

```shell

# set environment variables
export SLACK_APP_TOKEN=${app-level Slack token}
export SLACK_BOT_TOKEN=${bot Slack token}
export GITHUB_APP_ID=${GitHub App ID}
export GITHUB_INSTALLATION_ID=${GitHub installation ID}
export GITHUB_PEM_BINARY=${base64 encoded private key file.}

# run
go build
go run reacji-github-issue
```

# Limitations

- You need to run your own process.
    - Slack App uses `Socket Mode` and it is not currently allowed in the public Slack App Directory.
