[![Docker Repository on Quay](https://quay.io/repository/ursais/odoo-operator/status "Docker Repository on Quay")](https://quay.io/repository/ursais/odoo-operator)

# Odoo Operator

## Table of Contents
* [Prerequisites](#Prerequisites)
* [Build the operator](#Build-the-operator)
* [Build the bundle](#Build-the-bundle)
* [Usage](#Usage)
* [Support](#Support)

## Prerequisites

* Go through the [installation guide](https://sdk.operatorframework.io/docs/building-operators/ansible/installation).
* User authorized with `cluster-admin` permissions.

## Build the operator

Build and push the operator:
```shell
export VERSION=latest
export OPERATOR_IMG="docker.io/ursa/odoo-operator:$VERSION"
make docker-build docker-push IMG=$OPERATOR_IMG
```

## Build the bundle

Bundle your operator and push the bundle image:
```shell
make bundle IMG=$OPERATOR_IMG
export BUNDLE_IMG="docker.io/ursa/odoo-operator-bundle:$VERSION"
make bundle-build BUNDLE_IMG=$BUNDLE_IMG
make docker-push IMG=$BUNDLE_IMG
```

## Usage

Run your bundle:
```shell
export BUNDLE_IMG="docker.io/ursa/odoo-operator-bundle:$VERSION"
operator-sdk run bundle $BUNDLE_IMG
```

Create a sample Odoo custom resource:
```shell
kubectl apply -f config/samples/apps_v1alpha1_odoo.yaml
```

Uninstall the operator:
```shell
operator-sdk cleanup odoo-operator
```

## Support

Report any problem or question by creating an issue on the
[GitHub project](https://github.com/ursais/odoo-operator/issues).
