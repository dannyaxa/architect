# convox/architect

Create a CloudFormation stack for a Convox application.

## Usage

    $ docker run convox/architect -processes web,worker -balancers web

## Parameters

The resulting stack will expect the following parameters:

| Name           | Description                                                            |
|----------------|------------------------------------------------------------------------|
| `AMI`          | Application AMI. (See [convox/build](https://github.com/convox/build)) |
| `AllowSSHFrom` | Allow SSH from this CIDR block                                         |
| `Environment`  | URL to an  environment for this app (`.env` format)                    |
| `Key`          | ARN for a KMS key used to encrypt `Environment`                        |
| `SSHKey`       | Key name for SSH access                                                |

The stack will also expect these parameters for each process type:

| Name         | Description                                   |
|--------------|-----------------------------------------------|
| `WebCommand` | Override the default command for this process |
| `WebScale`   | Number of instances of this process to run    |
| `WebSize`    | Size of instance to run for process           |

And these additional parameters for any process behind a balancer:

| Name           | Description                       |
|----------------|-----------------------------------|
| `WebCheckPath` | Healthcheck path for this process |
| `WebCheckPort` | Healthcheck port for this process |
| `WebPort`      | Listen port for this process      |

## See Also

* [convox/build](https://github.com/convox/build)

## License

Apache 2.0 &copy; 2015 Convox, Inc.
