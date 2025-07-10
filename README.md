## Spark Go Example

This is an example Spark job in Go. This uses Spark Connect
to submit a job to a Spark cluster.

## Requirements

* go 1.23 or later
* GNU Make

## Installing Spark

```bash
curl -LO https://dlcdn.apache.org/spark/spark-4.0.0/spark-4.0.0-bin-hadoop3-connect.tgz && tar xvf spark-4.0.0-bin-hadoop3-connect.tgz
```

Ensure SPARK_HOME env variable is pointing to the extracted directory. Run this in your
shell or set this in your `~/.bashrc` or `~/.zshrc`:

```
export SPARK_HOME=<path-to>/spark-4.0.0-bin-hadoop3-connect
export PATH="$SPARK_HOME/bin:$SPARK_HOME/sbin:$PATH"
```

## Building

```bash
go build -o bin/spark_connect_app cmd/spark_connect_app.go
```

## Submitting Job

##### Submitting Job Using Spark Connect

First ensure the Spark Connect server is up and running:

```bash
$SPARK_HOME/sbin/start-connect-server.sh
```

Verify you can reach the UI at http://localhost:4040/.

Run the standalone app to submit the job:

```bash
chmod +x bin/spark_connect_app
cmd/spark_connect_app
```
You can also run the app directly using `go run`:

```bash
go run cmd/spark_connect_app.go
```

You should see something like so

```bash
Submitting job to sc://localhost:15002
+---+
|id |
+---+
|0  |
|1  |
|2  |
|3  |
|4  |
|5  |
|6  |
|7  |
|8  |
|9  |
+---+
```

