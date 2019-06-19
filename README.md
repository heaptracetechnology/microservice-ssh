# _SSH_ OMG Microservice

[![Open Microservice Guide](https://img.shields.io/badge/OMG%20Enabled-👍-green.svg?)](https://microservice.guide)

An OMG service for SSH, it  is a cryptographic network protocol for operating network services securely over an unsecured network.

## Direct usage in [Storyscript](https://storyscript.io/):

##### SSH
```coffee
>>> ssh exec command:'ifconfig' username:'username' password:'password' host:'192.168.1.01' port:'4587'
{"standardOutput": "Output for command","standardError": "Error if any occurred else empty","returnCode": "HTTPStatusCode"}
```

Curious to [learn more](https://docs.storyscript.io/)?

✨🍰✨

## Usage with [OMG CLI](https://www.npmjs.com/package/omg)

##### SSH
```sh
$ omg run exec -a command=<COMMAND> -a username=<SERVER_USERNAME> -a password=<SERVER_PASSWORD> -a host=<SSH_HOST> -a port=<PORT_NUMBER> -e PRIVATE_KEY=<PRIVATE_KEY_FILE_BASE64_DATA>
```

**Note**: the OMG CLI requires [Docker](https://docs.docker.com/install/) to be installed.

## License
[MIT License](https://github.com/omg-services/ssh/blob/master/LICENSE).
