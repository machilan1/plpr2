# Deployments

GCP_PROJECT_ID=yft-staging-447720

## Pre-requisites

### Install GCloud CLI

https://cloud.google.com/sdk/docs/install

### Install Pulumi

https://www.pulumi.com/docs/iac/download-install/

### Config project

```sh
export PROJECT_ID=<YOUR GCP PROJECT ID>
gcloud auth application-default login
gcloud config set project $PROJECT_ID
gcloud auth configure-docker
```

### Create a bucket as Pulumi backend

[WARNING] This should be done only once per project.

```sh
gsutil mb gs://pulumi-${PROJECT_ID}
pulumi login gs://pulumi-${PROJECT_ID}
```

### Create a new dev stack

[WARNING] This should be done only once per project.

<!-- pulumi config set env <value> -->
<!-- pulumi config set admin-email <value> -->

```sh
pulumi stack init dev
pulumi config set gcp:project $PROJECT_ID
```

## Pulumi

```sh
pulumi up --cwd ./deploy
```

Currently, no passphrase is set for the private key.

## Manual steps

### Github

1. The cloud build connection must be done manually in each GCP project. Or we need to try Cloud Build Gen 2.

### Domain verification

1. Use cloudflare for DNS.
2. Setup domain verification in GCP before adding the domain to Cloud Run.

### Domain mapping

Run `gcloud beta run domain-mappings describe --region <REGION> --domain <DOMAIN>` to get all the DNS records to add to Cloudflare.
All the records returned under `resourceRecords` should be added to Cloudflare.

## Cloud SQL Proxy

Add this to your `.env`:

```
export CLOUD_SQL_CONNECTION_NAME=yft-prod:asia-east1:yft-staging
export CLOUD_SQL_DB_PORT=26257
export CLOUD_SQL_DB_NAME=yft
```

Run the following command to start the Cloud SQL Proxy:

```sh
make cloudsql-proxy
```

To access the database, run the following command:

```sh
make cloudsql-proxy
```

## Cloud Run

```sh
gcloud builds triggers run <TRIGGER_NAME> --branch <GIT_BRANCH_NAME>
```
