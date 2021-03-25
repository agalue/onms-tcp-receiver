# onms-tcp-receiver

A simple server that receives data from the OpenNMS TCP Exporter

To use it, you should be using Horizon 19.1.x or newer, and add the following to a file inside `$OPENNMS_HOME/etc/opennms.properties.d/`:

```
org.opennms.timeseries.strategy=tcp

org.opennms.rrd.tcp.host=192.168.205.1
org.opennms.rrd.tcp.port=6123
```

The above will forward the metrics to `192.168.205.1` via TCP on port `6123`. The payload is Protobuf and emulates the parameters from a `rrdtool update` command (plus strings properties, if any). Please note that OpenNMS won't store the metrics to disk, so there will be no RRD/JRB files.

The following is sample output from this client, which is the JSON representation of the Protobuf payload, assuming `storeByGroup=false` and `storeByForeignSource=true` in OpenNMS. The data comes from a node with a Foreign ID of `opennms` that belongs to a requisition called `local`:

```
2021/03/25 12:07:19 Message: {"path":"/opt/opennms/share/rrd/fs/local/opennms/eth0-5254004d77d3/ifDescr","owner":"127.0.0.1","timestamp":1616688438000,"strValue":["eth0"]}
2021/03/25 12:07:19 Message: {"path":"/opt/opennms/share/rrd/fs/local/opennms/eth0-5254004d77d3/ifSpeed","owner":"127.0.0.1","timestamp":1616688438000,"strValue":["1000000000"]}
2021/03/25 12:07:19 Message: {"path":"/opt/opennms/share/rrd/fs/local/opennms/eth0-5254004d77d3/ifInOctets","owner":"127.0.0.1","timestamp":1616688438000,"dblValue":[13414336]}
2021/03/25 12:07:19 Message: {"path":"/opt/opennms/share/rrd/fs/local/opennms/eth0-5254004d77d3/ifOutOctets","owner":"127.0.0.1","timestamp":1616688438000,"dblValue":[3339636]}
```

The first two entries are examples of string attributes, whereas the last two are examples of numeric attributes. For the numeric entries, the above is equivalent to:

```bash
rrdtool update /opt/opennms/share/rrd/fs/local/opennms/eth0-5254004d77d3/ifInOctets.rrd 1616688438000:13414336
rrdtool update /opt/opennms/share/rrd/fs/local/opennms/eth0-5254004d77d3/ifOutOctets.rrd 1616688438000:3339636
```
