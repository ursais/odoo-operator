# Odoo Operator

## Table of Contents
* [Deployment](#Deployment)
  * [Create Custom Resource Definition](#Create-Custom-Resource-Definition)
  * [Deploy Operator](#Deploy-Operator)
* [Usage](#Usage)
* [Support](#Support)

## Deployment
### Create Custom Resource Definition

Create the CRD:
```shell
oc create -f ./deploy/crds/odoo-community.org_odoos_crd.yaml
```

Create the cluster role:
```shell
oc create -f ./deploy/cluster_role.yaml
```

### Deploy Operator

Operators can be deployed on a cluster-wide basis or on a namespace/project basis.

Create the project:
```shell
oc new-project odoo-operator --display-name="Odoo"
```

Create the service account:
```shell
oc create -f ./deploy/service_account.yaml
```

Create the role and role binding:
```shell
oc create -f ./deploy/role.yaml
oc create -f ./deploy/role_binding.yaml
oc create -f ./deploy/cluster_role_binding.yaml
```

Create the deployment:
```shell
oc create -f ./deploy/operator.yaml
```

## Usage

Create the Odoo instance:
```shell
oc create -f deploy/odoo.yaml
```

Examine the custom resource:
```shell
oc describe odoo example
```

Get the route of your Odoo instance:
```shell
oc get route
```

## Support

Report any problem or question by creating an issue on the
[GitHub project](https://github.com/ursais/odoo-operator/issues).
