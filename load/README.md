## Load Testing Steps

1. Create the deployment

    Choose the appropriate manifest and deploy.
   
    ```bash
    kubectl apply -f quick_start_app_envoy.yaml 
    ```

1. Make the application accessible outside the cluster.

    ```bash
    kubectl expose deployment example-app --type=NodePort --name=example-app-service --port=8080
    ```
   
1. Set the `SERVICE_URL` environment variable to the service’s IP/port.

   **minikube**:

    ```bash
    export SERVICE_PORT=$(kubectl get service example-app-service -o jsonpath='{.spec.ports[?(@.port==8080)].nodePort}')
    export SERVICE_HOST=$(minikube ip)
    export SERVICE_URL=$SERVICE_HOST:$SERVICE_PORT
    echo $SERVICE_URL
    ```

   **minikube (example)**:

    ```bash
    192.168.99.100:31380
    ```
   
1. Run load test

    Specify `duration` (amount of time to issue request to the targets) and `rate` (request rate per time unit to
    issue against the targets).

    ```bash
    echo "GET http://bob:password@$SERVICE_URL/people" | vegeta attack -duration=1s -rate=10 > result.bin
    ```
   
1. View the load test results

    ```bash
    vegeta report result.bin
    ```
   
1. Generate histogram plot

    ```bash
    vegeta report -type=hdrplot result.bin > r.hgrm
    ```