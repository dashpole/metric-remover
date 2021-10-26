package main

import(
	"context"
	"fmt"
	"flag"
	"log"

	monitoring "cloud.google.com/go/monitoring/apiv3"
	monitoringpb "google.golang.org/genproto/googleapis/monitoring/v3"
)

func main() {
	var metricName, project string
	flag.StringVar(&metricName, "metric", "", "The name of the metric to be removed")
	flag.StringVar(&project, "project", "", "The project id of the project in which to remove the metric")
	flag.Parse()
	if metricName == "" {
		log.Fatal("--metric is required")
	}
	if project == "" {
		log.Fatal("--project is required")
	}

	name := fmt.Sprintf("projects/%s/metricDescriptors/%s", project, metricName)
	ctx := context.Background()
	c, err := monitoring.NewMetricClient(ctx)
	if err != nil {
			log.Fatalf("failed to create metric client: %v", err)
	}
	defer c.Close()
	req := &monitoringpb.DeleteMetricDescriptorRequest{
			Name: name,
	}

	if err := c.DeleteMetricDescriptor(ctx, req); err != nil {
			log.Fatalf("could not delete metric: %v", err)
	}
	log.Printf("Deleted metric: %q\n", name)
}