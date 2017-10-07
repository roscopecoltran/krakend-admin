# krakend-admin

Data:    /usr/local/var/elasticsearch/elasticsearch
Logs:    /usr/local/var/log/elasticsearch/elasticsearch
Plugins: /usr/local/opt/elasticsearch/libexec/plugins/
Config:  /usr/local/etc/elasticsearch/
plugin script: /usr/local/opt/elasticsearch/libexec/bin/elasticsearch-plugin


To have launchd start elasticsearch now and restart at login:
  brew services start elasticsearch
Or, if you don't want/need a background service you can just run:
  elasticsearch