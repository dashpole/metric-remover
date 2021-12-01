# metric-remover
Removes a google cloud monitoring metric descriptor.

## Usage

### Downloading

```sh
wget https://github.com/dashpole/metric-remover/releases/download/v0.0.1/main
sudo chmod +x main
gcloud auth application-default login
./main --project my-project --metric workload.googleapis.com/my-metric
```

### Using go run
```
gcloud auth application-default login
go run main.go --project my-project --metric workload.googleapis.com/my-metric
```
