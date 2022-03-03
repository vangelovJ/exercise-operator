# exercise-operator
K8s Operator for exercises


### Generating commands:

````
operator-sdk init --domain vangelovj --repo github.com/vangelovj/exercise-operator
operator-sdk create api --group exercise --version v1 --kind Exercise --resource --controller
````

#### Building the operator

````````
make manifests
make docker-build
make docker-push
````````
#### Deploying the operator

````
make deploy
kubectl apply -f ./static_yaml/rbac.yaml
````


#### Deploying the CR example

````
kubectl apply -f ./config/samples/exercise_v1_exercise.yaml
````

#### Removing the operator

````
kubectl apply -f ./config/samples/exercise_v1_exercise.yaml
make undeploy
kubectl delete -f ./static_yaml/rbac.yaml
````

