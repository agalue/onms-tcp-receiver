# onms-tcp-receiver

A simple server that receives data from the OpenNMS TCP Exporter

To use it, you should be using Horizon 19.1.x or newer, and add the following to a file inside `$OPENNMS_HOME/etc/opennms.properties.d/`:

```
org.opennms.timeseries.strategy=tcp

org.opennms.rrd.tcp.host=192.168.205.1
org.opennms.rrd.tcp.port=6123
```

The above will forward the metrics to `192.168.205.1` via TCP on port `6123`. The payload is Protobuf and emulates the parameters from a `rrdtool update` command (plus strings properties, if any). Please note that OpenNMS won't store the metrics to disk, so there will be no RRD/JRB files.