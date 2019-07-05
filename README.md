# _SSH_ OMG Microservice

[![Open Microservice Guide](https://img.shields.io/badge/OMG%20Enabled-👍-green.svg?)](https://microservice.guide)
[![Build Status](https://travis-ci.com/heaptracetechnology/microservice-ssh.svg?branch=master)](https://travis-ci.com/heaptracetechnology/microservice-ssh)
[![codecov](https://codecov.io/gh/heaptracetechnology/microservice-ssh/branch/master/graph/badge.svg)](https://codecov.io/gh/heaptracetechnology/microservice-ssh)

An OMG service for SSH, it  is a cryptographic network protocol for operating network services securely over an unsecured network.

## Direct usage in [Storyscript](https://storyscript.io/):

```coffee
>>> ssh exec command:'ifconfig' username:'username' password:'password' host:'serverHost' port:'portNumber' 
{"standardOutput":"Output of executed command","standardError":"Failure message if any error occurred","returnCode":"HTTPstatusCode"}
```

Curious to [learn more](https://docs.storyscript.io/)?

✨🍰✨

## Usage with [OMG CLI](https://www.npmjs.com/package/omg)

##### SSH Execute
```shell
$ omg run exec -a command=<COMMAND> -a username=<SERVER_USERNAME> -a password=<SERVER_PASSWORD> -a host=<SSH_HOST> -a port=<PORT_NUMBER> -e PRIVATE_KEY=<PRIVATE_KEY_FILE_BASE64_DATA>
```

**Note**: The OMG CLI requires [Docker](https://docs.docker.com/install/) to be installed.

## License
[MIT License](https://github.com/omg-services/awesome/blob/master/LICENSE).




