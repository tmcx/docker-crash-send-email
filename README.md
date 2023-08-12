<h1 align="center">DOCKER CRASH</h1>
<h3 align="center">EMAIL NOTIFICATION</h3>

<p align="center">
  <img src="assets/logo128.png" width="120px" height="120px"/>
  <br>
  <i>Email notification in the event of Docker crashes.</i>
  <br>
</p>

<hr>

## Build
Run ``go build``. This will generate a file called ``docker-crash-send-email``. This same you can run it from a console to start its operation.

## Usage
1. Create a file called ``configuration.json`` and place it along with the executable. You can use the example located in the example folder to guide you.
2. From some console, run the builded file.

<b>I recommend you execute it as a service.</b>

## License

[MIT](https://choosealicense.com/licenses/mit/)