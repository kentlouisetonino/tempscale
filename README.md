### Description
#

<br />

![Screenshot from 2024-05-14 16-30-28](https://github.com/kentlouisetonino/tempscale/assets/69438999/35abc607-ea0a-460a-b596-6fed38743484)


<br />

> - A CLI tool for converting temperature scale.

> - This supports `Celsius` °C, `Fahrenheit` °F, `Kelvin` °K, and `Rankine` °R scales.

<br />
<br />



### System Design
#

![tempscale-system-design](https://github.com/kentlouisetonino/tempscale/assets/69438999/07170a87-049d-4719-b2eb-d2cd65e6d70b)

<br />
<br />



### Local Development
#

> - Note the root `main.go` is a symbolic link.

> - Run the following commands.

```sh
# Change permission and make it executable.
sudo chmod +x run.sh

# Run the app.
./run.sh
```

<br />
<br />



### Deployment
#

```sh
# Run the multipass.
sudo systemctl restart snap.multipass.multipassd.service

# Deploy to snap store.
snapcraft login
snapcraft clean
snapcraft --debug
snapcraft push tempscale_<version-number-in-snapcraft-yaml>_amd64.snap --release=stable
```
