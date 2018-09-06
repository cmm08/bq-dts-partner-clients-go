# Getting Started with BigQuery Data Transfer API using Go

<a href="https://console.cloud.google.com/cloudshell/open?git_repo=https://github.com/GoogleCloudPlatform/java-docs-samples&page=editor&open_in_editor=bigquery/datatransfer/cloud-client/README.md">
<img alt="Open in Cloud Shell" src ="http://gstatic.com/cloudssh/images/open-btn.png"></a>

[BigQuery Data Transfer Service][BigQuery Data Transfer] features an API that
allows developers to create transfer jobs from data sources to BigQuery.
These sample Go application demonstrates how to access the BigQuery Data
Transfer API using the [BigQuery Data Transfer Service Client Library for
Go][bigquery-dts-client-go].

[BigQuery Data Transfer]: https://cloud.google.com/bigquery/docs/transfer-service-overview
[bigquery-dts-client-go]: https://github.com/cmm08/bq-dts-partner-clients-go/tree/master/gapi-cloud-bigquerydatatransfer-go 

## Quickstart

### Before you begin

Please do the following, to set up GCP project with necessary APIs, service
accounts, and permissions:

1. Install the [gcloud SDK](https://cloud.google.com/sdk/downloads#interactive)

2. Enable the following APIs in the [Google Developers Console](https://console.developers.google.com/project/_/apiui/apiview/pubsub/overview).
    * BigQuery Data Transfer API
    * Cloud Pub/Sub API

    ```
    gcloud services enable bigquerydatatransfer.googleapis.com
    gcloud services enable pubsub.googleapis.com
    ```

3. Create an operational IAM Service Account and download credentials for running your source.  Running these commands
    * Creates a new Service Account named `bq-dts-[SOURCE]@[PROJECT_ID].iam.gserviceaccount.com`
    * Grants `roles/bigquery.admin`, `roles/pubsub.subscriber`, `roles/storage.objectAdmin`
    * Downloads a Service-Account key called `.gcp-service-account.json`

    ```
    SOURCE="example-source"  # Change to your own name 
    PROJECT_ID=$(gcloud config get-value core/project)
    PARTNER_SA_NAME="bq-dts-${SOURCE}"
    PARTNER_SA_EMAIL="${PARTNER_SA_NAME}@${PROJECT_ID}.iam.gserviceaccount.com"

    # Creating a Service Account
    gcloud iam service-accounts create ${PARTNER_SA_NAME} --display-name ${PARTNER_SA_NAME}

     # Granting Service Account required roles
    gcloud projects add-iam-policy-binding ${PROJECT_ID} --member="serviceAccount:${PARTNER_SA_EMAIL}" --role='roles/bigquery.admin'
    gcloud projects add-iam-policy-binding ${PROJECT_ID} --member="serviceAccount:${PARTNER_SA_EMAIL}" --role='roles/pubsub.subscriber'
    gcloud projects add-iam-policy-binding ${PROJECT_ID} --member="serviceAccount:${PARTNER_SA_EMAIL}" --role='roles/storage.objectAdmin'

    # Create service account credentials and store it locally needed for starting/running data sources.
    gcloud iam service-accounts keys create --iam-account "${SERVICE_ACCOUNT_EMAIL}" .gcp-service-account.json
    ```

4. Create an administrative IAM Service Account and store credentials locally for creating data source.
    * Creates a new Service Account named `bq-dts-admin@[PROJECT_ID].iam.gserviceaccount.com`
    * Grants `roles/project.owner`
    * Downloads a Service-Account key called `.gcp-service-account.json`

    ```
    PROJECT_ID=$(gcloud config get-value core/project)
    PARTNER_SA_NAME="bq-dts-admin"
    PARTNER_SA_EMAIL="${PARTNER_SA_NAME}@${PROJECT_ID}.iam.gserviceaccount.com"

    # Creating a Service Account
    gcloud iam service-accounts create ${PARTNER_SA_NAME} --display-name ${PARTNER_SA_NAME}
    gcloud projects add-iam-policy-binding ${PROJECT_ID} --member="serviceAccount:${PARTNER_SA_EMAIL}" --role='roles/owner'

    # Create service account credentials and store it locally needed for creating data source
    gcloud iam service-accounts keys create --iam-account "${PARTNER_SA_EMAIL}" .gcp-service-account-owner.json
    ```

   For more information on service account authentication, please refer to [Authenticate using a service account](https://cloud.google.com/docs/authentication/getting-started).

5. Grant permissions to a GCP-managed Service Account
    * Creates a custom role - `bigquerydatatransfer.connector` with permission `clientauthconfig.clients.getWithSecret`
    * Grants project-specific role `bigquerydatatransfer.connector`

    ```
    PROJECT_ID=$(gcloud config get-value core/project)
    GCP_SA_EMAIL="connectors@bigquery-data-connectors.iam.gserviceaccount.com"

    # Creating a custom role
    gcloud iam roles create bigquerydatatransfer.connector --project ${PROJECT_ID} --title "BigQuery Data Transfer Service Connector" --description "Custom role for GCP-managed Service Account for BQ DTS" --permissions "clientauthconfig.clients.getWithSecret" --stage ALPHA

    # Granting Service Account required roles
    gcloud projects add-iam-policy-binding ${PROJECT_ID} --member="serviceAccount:${GCP_SA_EMAIL}" --role="projects/${PROJECT_ID}/roles/bigquerydatatransfer.connector"
    ```

6. Create an [OAuth Consent Screen](https://support.google.com/cloud/answer/6158849?hl=en#userconsent)

7. Join BigQuery Data Transfer Service Partner-level whitelists.  Reach out to your Google Cloud Platform contact to get whitelisted for these APIs.

8. Install [Maven](http://maven.apache.org/). The example program that makes
   BigQuery Data Transfer Service API calls is built using Maven in this
   QuickStart. If you perfer, you can also use Gradle, in which case the Gradle
   Wrapper is included and you don't need to install Gradle.

### Run the Example

1. Set up environment variables:

   ```
   export PROJECT_ID="<your project id>"
   export GOOGLE_APPLICATION_CREDENTIALS=.gcp-service-account-owner.json
   ```

2. Try the example program that calls BigQuery Data Transfer Service
   APIs:

   ```
   cd gapi-cloud-bigquerydatatransfer-go
   go run main.go "${PROJECT_ID}"
   ```

### Notes

For illustration purpose, the example calls the following two BigQuery Data Transfer API methods:
  * [ListTransferConfigs](https://cloud.google.com/bigquery/docs/reference/datatransfer/rest/v1/projects.locations.transferConfigs/list)

  * [ListDataSourceDefinitions](https://cloud.google.com/bigquery/docs/reference/data-transfer/partner/rest/v1/projects.locations.dataSourceDefinitions/list)

Note that the 2nd method is only visible to whitelisted partner projects. Plese talk to your Google Cloud Platform contact.
