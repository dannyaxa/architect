# convox/architect

Create a CloudFormation stack for a Convox application.

## Usage

    $ docker run convox/architect -processes web,worker -balancers web

## Parameters

The resulting stack will expect the following parameters:

| Name                | Description                                                                |
|---------------------|----------------------------------------------------------------------------|
| `AMI`               | Application AMI. (See [convox/builder](https://github.com/convox/builder)) |
| `AllowSSHFrom`      | Allow SSH from this CIDR block                                             |
| `AvailabilityZones` | A comma-delimited list of availability zones to use (specify 3)            |
| `Environment`       | URL to an  environment for this app (`.env` format)                        |


The stack will also expect these parameters for each process type:

| Name                | Description                                   |
|---------------------|-----------------------------------------------|
| `WebCommand`        | Override the default command for this process |
| `WebPorts`          | Port mappings for this process                |
| `WebScale`          | Number of instances of this process to run    |

## License

Apache 2.0 &copy; 2015 Convox, Inc.
