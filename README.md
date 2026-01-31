## Description

> - A playground anything related to Databases integration with Go.

> - Below are the available routes.

```
http://localhost:3000

# Available routes.
/ready
/redis/ready
/sql/ready
/sql/create/user/table
```


<br />
<br />
<br />



## Setup

> - Install `GNU Make` and `Grafana k6`.

> - Run the following commands.

```sh
# Run the HTTP server.
make run

# (Optional) If want to run k6 load testing.
make run-k6
```
