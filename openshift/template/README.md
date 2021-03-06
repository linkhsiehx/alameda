# Install Alameda with Template

This directory contains OpenShift application templates for Alameda to install [develop](./dev), [build](./build) and [deploy](./deploy). 

The OpenShift Template documentation as [here](https://docs.okd.io/3.11/dev_guide/templates.html)

# Apply Template to OpenShift

Login as admin, then upload the templates for differe enviroment.

```
oc login -u admin
```

## Develop

The Template is used to create the alameda development enviroment

```
oc create -f ./dev/alameda-crd-rbac.yaml
oc create -f ./dev/alameda-datahub.yaml
oc create -f ./dev/alameda-operator.yaml
oc create -f ./dev/alameda-ai.yaml
```

## Build

The Template is used to create the alameda build enviroment

```
oc create -f ./build/alameda-ai.yaml
oc create -f ./build/alameda-rhel7.yaml
oc create -f ./build/alameda.yaml
```

## Deploy

The Template is used to create the alameda application enviroment

```
oc create -f ./deploy/alameda.yaml
```
