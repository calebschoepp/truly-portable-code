# Deploying to SpinKube

- [Deploying to SpinKube](#deploying-to-spinkube)
  - [1. Install SpinKube on AKS Kubernetes Cluster](#1-install-spinkube-on-aks-kubernetes-cluster)
  - [2. Deploy URL Shortener with Valkey Backing Storage](#2-deploy-url-shortener-with-valkey-backing-storage)
  - [Learning Summary](#learning-summary)
  - [Navigation](#navigation)

In this module we'll explore how to deploy our URL shortening Spin application to a Kubernetes cluster. Deploying to Kubernetes allows us to easily scale our application and take advantage of Kubernetes' robust cloud native ecosystem.

We'll use [SpinKube](https://www.spinkube.dev/) to deploy our Spin application to Kubernetes. SpinKube is a project that makes it easy to run Spin applications on Kubernetes. SpinKube is the collection of a number of open source components:

- **[Spin Operator](https://github.com/spinframework/spin-operator/):** A Kubernetes operator that enables deploying and running Spin applications in Kubernetes. It houses the SpinApp and SpinAppExecutor CRDs which are used for configuring the individual workloads and workload execution configuration such as runtime class.
- **[Containerd Shim Spin](https://github.com/spinframework/containerd-shim-spin):** A shim for running Spin workloads managed by containerd. The Spin workload uses this shim as a runtime class within Kubernetes enabling these workloads to function similarly to container workloads in Pods in Kubernetes.
- **[Kwasm](https://kwasm.sh/):** An operator that automates and manages the lifecycle of containerd shims in a Kubernetes environment. This includes tasks like installation, update, removal, and configuration of shims, reducing manual errors and improving reliability in managing WebAssembly (Wasm) workloads and other containerd extensions. (This will be replaced with [Runtime Class Manager](https://github.com/spinframework/runtime-class-manager) in the future.)

## 1. Install SpinKube on AKS Kubernetes Cluster

We're going to be using [Azure Kubernetes Service (AKS)](https://azure.microsoft.com/en-us/products/kubernetes-service) for our Kubernetes cluster. We need to create an AKS cluster and then install SpinKube on it. Then we'll be able to deploy the URL shortener Spin application to the cluster.

Follow along with the guide [here](https://www.spinkube.dev/docs/install/azure-kubernetes-service/). It will walk you through creating an AKS cluster, installing SpinKube on it, deploying a sample application, and testing it works.

Once you have that working come back here and go to the next step.

## 2. Deploy URL Shortener with Valkey Backing Storage

Now it is time to deploy our URL shortener Spin application to the AKS cluster. When running a Spin app locally the key value store is backed by Sqlite. In Kubernetes however, we'll want a single persistent service running otherwise every replica of our Spin app would be using its own local Sqlite database which would lead to inconsistent results.

Let's setup [Valkey](https://valkey.io/) as the backing storage for key value.

```bash
$ helm repo add valkey https://valkey.io/valkey-helm/
$ helm repo update
$ helm install valkey --namespace valkey --create-namespace oci://registry-1.docker.io/bitnamicharts/valkey
```

We will need the URL for Valkey so that we can configure our Spin application to use it. Let's put the URL in a K8s secret:

```bash
$ kubectl create secret generic kv-secret --from-literal=valkey-url="redis://valkey.valkey.svc.cluster.local:6379"
```

With Valkey running we can turn our attention to deploying the Spin app. First, we need to push our Spin app to a registry so that it can be pulled down by Spin Operator.

Spin Apps are packaged and distributed as OCI artifacts. By leveraging OCI artifacts, Spin Apps can be distributed using any registry that implements the Open Container Initiative Distribution Specification (a.k.a. “OCI Distribution Spec”). In this module we'll use [ttl.sh](https://ttl.sh/) as our registry.

The spin CLI simplifies packaging and distribution of Spin Apps and provides an atomic command for this (`spin registry push`). You can package and distribute the url-shortener Spin app with the following command:

```bash
$ spin registry push --build ttl.sh/<your-unique-image-name>:24h
```

> **Note:** Make sure to replace `<your-unique-image-name>` with a unique name for your Spin app otherwise you may run into naming conflicts with other users.

> **Note:** It is a good practice to add the `--build` flag to `spin registry push`. It prevents you from accidentally pushing an outdated version of your Spin App to your registry of choice.

Now we need to create a SpinApp custom resource to tell Spin Operator to deploy our Spin app. Create a file named `url-shortener-spinapp.yaml` with the following contents:

```yaml
apiVersion: core.spinkube.dev/v1alpha1
kind: SpinApp
metadata:
  name: url-shortener
spec:
  image: ttl.sh/<unique-image-name>:24h
  replicas: 1
  executor: containerd-shim-spin
  runtimeConfig:
    keyValueStores:
      - name: "default"
        type: "redis"
        options:
          - name: "url"
            valueFrom:
              secretKeyRef:
                name: "kv-secret"
                key: "valkey-url"
```

We specify a few important things in this custom resource. We provide an image where the app can be found. A replica count of how many pods we want. That the containerd-shim-spin executor should be used. And finally we configure the key value store to use Valkey by referencing the K8s secret we created earlier. This configuration is done via [Spin runtime config](https://spinframework.dev/v3/dynamic-configuration).

Go ahead and apply the SpinApp custom resource to the cluster:

```bash
$ kubectl apply -f url-shortener-spinapp.yaml
```

Now we can test that it is working. First we need to get the name of the pod that is running our Spin app:

```bash
$ kubectl get pods
NAME                                 READY   STATUS    RESTARTS   AGE
url-shortener-xxxxxx-xxxxx           1/1     Running   0         2m
```

Now we can port forward to the Spin app:

```bash
$ kubectl port-forward url-shortener-xxxxxx-xxxxx 8080:80
```

In another terminal:

```bash
$ curl localhost:3000/foo -i
HTTP/1.1 404 Not Found

$ curl localhost:3000/foo -i --data 'https://wikipedia.org'
HTTP/1.1 201 Created

$ curl localhost:3000/foo -i
HTTP/1.1 302 Found
location: https://wikipedia.org
```

> **Note:** For fun you could create another separate Spin App CR with a different name that uses the same valkey. Then you could prove to yourself that the data is shared between the two Spin apps because you won't need to shorten the URL on the new app.

## Learning Summary

In this module you learned how to:

- Install SpinKube on an AKS Kubernetes cluster
- Deploy Valkey as the backing storage for key value
- Push a Spin app to an OCI registry
- Create a SpinApp custom resource to deploy a Spin app to Kubernetes

## Navigation

- Go back to [Local Spin](01-local-spin.md) if you still have questions about the previous section
- Otherwise, proceed to [Deploying to Fermyon Wasm Functions](03-fwf.md).

If you have any feedback for this module please open an issue on the GitHub repo [here](https://github.com/calebschoepp/truly-portable-code/issues/new).
