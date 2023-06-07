# Any Metric 


:shipit: A go application that generates "fake" metrics on prom format. Perfect on early stages infrastructure building
to test prometheus metric-scrapping before actual workloads are deployed.

## Features

- Supports gauge and counter metric types.
- Metrics can be tailored as to its name, frequency and increment.
- Metrics are updated automatically based on the given frequency and increment.
- Exposes metrics on `/metrics` path.
- Typically uses (4+N) goroutines, where N is the number of metrics being generated.

## Typical utilization 

Any metric accepts any number of the `metric` flag: The flag format is: `--metric:<name>:<frequency>:<increment>:<type>`:
- `<name>`: Metric's name. Must be along https://prometheus.io/docs/practices/naming guidelines .
- `<frequency>`: How often the metric is increased. Accepts times using a fraction of value and a time prefix (e.g. 1ms, 2s, 3h). Please refer to https://pkg.go.dev/time#ParseDuration for further detail.
- `<increment>`: By how much the metric is increased. Accepts a 64-bit float.
- `<type>`: Type of the metric. Supports counter and gauge.

### Locally from the terminal
- `./any-metric --metric=potatoWeight:3s:0.4:gauge --metric=tomatoesInTheBag:1s:0.2:counter`
- `docker run ghcr.io/luizfnunesmarques/any-metric:latest --metric=airplanesInTheSky:1s:0.5:gauge` 

### Standalone pod
- Please refer to the [examples/standalone-pod.yaml](https://github.com/luizfnunesmarques/any-metric/blob/main/examples/standalone-pod.yaml) file.  
  
### Using the helm chart
A helm chart to help deploying the app is available at github.com/luizfnunesmarques/helm-charts.

- Add the repo: `helm repo add luizmarques https://luizfnunesmarques.github.io/helm-charts`
- Install the chart: `helm install any-metric luizmarques/any-metric`. The metrics can be passed along through the [values file](https://github.com/luizfnunesmarques/helm-charts/blob/main/charts/any-metric/values.yaml).
  
