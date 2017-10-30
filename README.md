# Go Twitter Thanks Bot

Super easy-to-use Twitter thanks bot written in GO that send direct message on follow request. Built on the popular go twitter package [go-twitter](https://github.com/dghubble/go-twitter). Please read Twitter best practices before setting up a bot.

## Prerequisites

This bot uses the go-twitter client library for the Twitter API, you will need to install it before you can use the Twitter bot.

```
go get github.com/dghubble/go-twitter/twitter
go get github.com/dghubble/oauth1
```

## Features

You can use *Go Twitter Thanks Bot* to:

* Reply to any twitter follow request with a direct thank you message.

## Install

You can install the Twitter Follow Bot using `go get`

```bash
go get github.com/ceccode/go-twitter-thanks-bot
```


## Usage

### Configuring the bot

Before running the bot, you must first set it up so it can connect to the Twitter API. Create a .env file and fill in the following information:

```
CONSUMER_KEY=
CONSUMER_SECRET=
ACCESS_TOKEN=
ACCESS_TOKEN_SECRET=
```

Create your keys here: [https://apps.twitter.com/](https://apps.twitter.com/).


### Running the bot

Compile and run

```
go run main.go
```

## Contributing

Feel free to contribute.

## License

This project is licensed under the [MIT License](LICENSE)

## Acknowledgments

* Dalton Hubble, author of go-twitter.
