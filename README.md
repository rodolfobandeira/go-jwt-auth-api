# go-jwt-auth-api

A simple API using Go Fiber to provide JWT authentication. Also a simple Kubernetes config to spin up all that

```
curl -X POST http://localhost/login \
   -H "Content-Type: application/json" \
   -d '{"user":"rodolfo","pass":"bandeira"}' \
```

```
curl --location --request GET 'http://localhost:3000/restricted' \
--header 'Authorization: Bearer eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJhZG1pbiI6dHJ1ZSw\
iZGV0YWlscyI6IkFueSBkZXRhaWwgZm9yIHRoZSB1c2VyIGNhbiBiZSBhZGRlZCBoZXJlLiA6KSIsImV4cCI6MTYz\
ODY1MzA0MCwibmFtZSI6IlJvZG9sZm8gQmFuZGVpcmEifQ.df-CWh_coFlMaShQCSv3WdscpL_Q_EEUhGJIvfsQ9Wd\
JHf9NzWwj6OmGTlHWKne-UXGcHHmKxIy8arzGAH5eZldiYTd_8STsFxXldXpFWWtsQYheOA1J7s4OtzN2LRiTNuz\
si5jf3ZByrrTUtZJEfuxzZQ24UAF22EZ5DE8P7FjAjGP1NxD44LXSAqrhZkPRESjze77q-y2HvvlUYD8K3pkgvvt\
1we6bx3dvLmOpHy_aXv76bEXX1whQ5JVDIJwSHMtmjxha2G118afZxH89Fy0xI182nfvPpKxIh6Ib-H2ZT0gcXxF\
ShimobhGoHyxkScqYUdrsqBeGI1chyEI_Qg'
```

---


#### Kubernetes in local enviroment:

```bash
minikube start

# Starting the Golang App API pod running on port 3000
kubectl apply -f k8s/deployment.yml

# Checking pod status:
kubectl get pods

# Creating a service of the type "NodePort". NodePort will
# make the external world able to access the cluster through this service
# using the "minikube ip" and the specific port
kubectl apply -f k8s/service-nodeport.yaml

# Checking which port should we use:
kubectl get services | grep NodePort
# service-nodeport   NodePort    10.107.102.51   <none>        8080:30003/TCP   70m

# Finally, accessing the pod App API passing through the NodePort service, and all this,
# from the external world (outside the cluster access):

# Get the network IP to access externally (From laptop)
minikube service service-nodeport --url
# http://192.168.49.2:30003

curl http://192.168.49.2:30003

# Scaling to 10 pods:
kubectl scale=10 -f k8s/deployment.yml

# Downgrade to 2:
kubectl scale --replicas=2 -f k8s/deployment.yaml
# deployment.apps/pod-api-app scaled


kubectl get pods
# NAME                          READY   STATUS    RESTARTS   AGE
# pod-api-app-598c8fc75-2gvh8   1/1     Running   0          4m46s
# pod-api-app-598c8fc75-5ct8z   1/1     Running   0          4m5s


# Running a MySQL (MariaDB) server:
kubectl apply -f k8s/db.yaml
kubectl exec -it db -- bash
mysql -u user -p
# user_password
# MariaDB [(none)]> show databases;
# +--------------------+
# | Database           |
# +--------------------+
# | information_schema |
# | my_db              |
# +--------------------+
# 2 rows in set (0.000 sec)
```
