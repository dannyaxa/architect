# convox/architect

Build a Convox stack.

## Usage

    $ docker run convox/architect \
      -processes web,worker
      -service POSTGRES_URL=postgres://user:pass@example.org/db
      -service REDIS_URL=redis://:pass@example.org/0

## Userdata

The resulting stack will expect the following parameters:

<table>
  <tr>
    <td><code>AMI</code></td>
    <td>Application AMI. See <a href="https://github.com/convox/builder">convox/builder</a>
  </tr>
  <tr>
    <td><code>AllowSSHFrom</code></td>
    <td>Allow SSH from this CIDR block</td>
  </tr>
  <tr>
    <td><code>AvailabilityZones</code></td>
    <td>A comma-delimited list of availability zones to use (specify 3)</td>
  </tr>
  <tr>
    <td><code>Environment</code></td>
    <td>URL to an  environment for this app (<code>.env</code> format)</td>
  </tr>
</table>

The stack will also expect these parameters for each process type:

<table>
  <tr>
    <td><code><i>Process</i>Command</code></td>
    <td>Override the default command for this process</a>
  </tr>
  <tr>
    <td><code><i>Process</i>Scale</code></td>
    <td>Number of instances of this process to run</td>
  </tr>
</table>

## License

Apache 2.0 &copy; 2015 Convox, Inc.