import 'core-js/stable';
const runtime = require('@wailsapp/runtime');

function stopProxy() {
  window.backend.LNDProxy.Stop();
}

function startProxy() {
  var form = document.getElementById('proxy-form');
  var stopButton = document.getElementById('stop-proxy');
  var message = document.getElementById('message');
  var address = document.getElementById('address').value;
  var cert = document.getElementById('cert').value;
  var port = document.getElementById('port').value;
  if (address == '' || cert == '' || port == '') {
    alert('All fields are required');
    return;
  }
  window.backend.LNDProxy.StartProxy(address, cert, port).then(function (result, err) {
    if (err) {
      message.innerHTML = `Something went wrong. Please check your input. ${err}`;
      return false;
    }
    message.innerHTML = `<p>
        LND Proxy is running <br> http://localhost:${port} ⇨ ${address}
      </p>
      <p>Configure Joule to use REST API URL: http://localhost:${port}<p>`;
    message.style.display = 'block';
    stopButton.style.display = 'block';
    form.style.display = 'none';
  }).catch(function(err) {
    var message = document.getElementById('message');
    message.innerHTML = `Something went wrong. Please check your input. ${err}`;
  });
}

function init() {
	var app = document.getElementById('app');

  app.innerHTML = `
    <h1 id="title">Joule LND Proxy</h1>
    <div id="proxy-form">
      <p>
        LND Address: <br>
        <input type="text" id="address" class="input" placeholder="https://123.456.789:8080">
      </p>
      <p>
        LND Certificate: <br>
        <textarea type="text" id="cert" class="input" placeholder="-----BEGIN CERTIFICATE-----"></textarea>
      </p>
      <p>
        Local Port: <br>
        <input type="text" id="port" class="input" value="8080">
      </p>
      <button id="start-proxy">Run</button>
    </div>
    <div id="message"></div>
    <button id="stop-proxy" style="display:none">Stop & Close</button>
    <div class="bg-circle"></div>
  `;
  document.getElementById('start-proxy').addEventListener('click', startProxy);
  document.getElementById('stop-proxy').addEventListener('click', stopProxy);
};

runtime.Init(init);