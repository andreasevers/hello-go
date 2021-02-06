# hello-go
Simple app for demos

## The app itself

```bash
cd src
go run hello-server.go world
```

## Dockerfile

```bash
cd ..
bat Dockerfiles/*
docker build -t hello -f Dockerfiles/Dockerfile-6 src
docker run hello
```

## kbld

```bash
bat kbld/deployment.yml
bat kbld/config.yml
kbld -f kbld/ | kubectl apply -f -
```

## pack

```bash
cd src
pack set-default-builder gcr.io/paketo-buildpacks/builder:base-platform-api-0.3
pack build hello-go # --no-pull
docker images | grep hello-go
pack inspect-image hello-go
pack inspect-image hello-go --bom | jq
pack inspect-builder
```

### Rebasing

```bash
docker pull gcr.io/paketo-buildpacks/run:0.0.97-base-cnb
docker tag gcr.io/paketo-buildpacks/run:0.0.97-base-cnb paketobuildpacks/run:base-cnb
docker images | grep run
docker images | grep hello-go
pack rebase hello-go --no-pull
catd <(docker inspect <oldsha>) <(docker inspect hello-go) | tail -n20
```

### Publishing

```bash
docker tag hello-go andreasevers/hello-go:pack-1.0.0
pack build andreasevers/hello-go:pack-1.0.0 --publish
```

## envsubst

```bash
cd ../envsubst
catd deployment-dev.yml deployment-prod.yml
bat deployment.yml
NAMESPACE=dev
envsubst < deployment.yml | kubectl apply -f -
```

## kustomize

```bash
cd ../kustomize
cd ops/base
bat deployment.yaml

cd ../overlays/dev
bat kustomization.yaml
kustomize build --load_restrictor none .
kustomize build --load_restrictor none . | kubectl apply -f -
k get all -n dev

cd ../prod
bat *
kustomize build --load_restrictor none . | kubectl apply -f -
kns prod
k get all
```

## ytt

### Overlays

```bash
cd ../../../ytt/overlays
bat *
ytt -f deployment.yml -f dev.yml | kubectl apply -f-
ytt -f deployment.yml -f prod.yml
```

### Pythonic

```bash
cd ../pythonic
bat *
ytt -f base/ -f envs/dev.yml | kubectl apply -f-
ytt -f base/ -f envs/prod.yml
```

## Helm

```bash
cd ../../helm
tree .
bat all-the-things
kns dev
helm upgrade --install hello-go -f values.yml -f values-dev.yml .
k get pods
```


