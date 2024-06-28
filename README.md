### Description
#

<br />

![Screenshot from 2024-05-18 15-11-30](https://github.com/kentlouisetonino/tempscale/assets/69438999/531d6b51-0b1b-4e52-b4a4-0bb90b36003a)


<br />

> - A CLI tool for converting temperature scale.

> - This supports `Celsius` °C, `Fahrenheit` °F, `Kelvin` °K, and `Rankine` °R scales.

> - If you want to check how I made this project from scratch, you can view it [here](https://www.youtube.com/playlist?list=PLPks-uiro_XJdp-MQMQTXjPkyzqQK4jUD). <br />
    Spoiler alert, it's pretty boring.

<br />
<br />



### System Design
#

![System Design](https://github.com/kentlouisetonino/tempscale/assets/69438999/1eb9642c-447c-45ee-a20c-d574250701cc)

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
