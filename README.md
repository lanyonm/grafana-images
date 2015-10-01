# grafana-images [![Build Status](https://travis-ci.org/lanyonm/grafana-images.svg)](https://travis-ci.org/lanyonm/grafana-images) [![Coverage Status](https://coveralls.io/repos/lanyonm/grafana-images/badge.svg)](https://coveralls.io/r/lanyonm/grafana-images)
This program interacts with [Grafana](http://grafana.org/) and [hubot-grafana](https://github.com/criticalmass/hubot-grafana) to provide facility to copy/save Grafana panel images to a location on disk. The idea is that this location is then shared by a web server so the images can be publically available. The rough system call diagram is as follows:

![grafana-images http call diagram](http://blog.lanyonm.org/images/grafana-images-diagram-no-numbers.svg)

The HTTP post expected by grafana-images should have two headers and a json payload:

```
"Accept: application/json"
"Authorization: Bearer grafana-token-goes-here"
```

```javascript
{
  "imageUrl": "https://grafana.test.com/render/dashboard-solo/db/sample-dashboard/?panelId=5&width=1000&height=500&from=now-6h&to=now&var-server=test-server"
}
```

The returned json will have a url that can be used to publically access the Grafana panel image:

```javascript
{
  "pubImg": "http://grafana.example.com/d494b123e0c40229ca3f1e9015390578.png"
}
```

In addition to this application, it is assumed that you have a web server set up to serve the saved images at the url returned. All the necessary settings should be configurable:

```
$ grafana-images --help
Usage of grafana-images:
  -imageHost="http://grafana.example.com/saved-images": host for the saved images
  -imagePath="/opt/saved-images": location on disk where images will be saved
  -port=8080: grafana-images listening port
```

For more information on how this fits together have a look at [ChatOps: Hubot Grafana Images in HipChat](http://blog.lanyonm.org/articles/2015/09/30/chatops-hubot-grafana-images-hipchat.html).

## Building
This will run tests as well.

	make

## Running

	make run

If you want to test the running system, you'll need to send it a json payload along with a couple of headers:

```bash
curl -d '{"imageUrl":"https://grafana.test.com/render/dashboard-solo/db/sample-dashboard/?panelId=5&width=1000&height=500&from=now-6h&to=now&var-server=test-server"}' -H "Accept: application/json" -H "Authorization: Bearer 1234567543ewsfdgdh432345awdf=" http://localhost:8080/grafana-images
```

## Test Coverage

	make cover

