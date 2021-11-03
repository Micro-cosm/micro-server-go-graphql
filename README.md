# micro-server-go-graphql

A Go-based Graphql service configured to run on GCP/Docker which connects to Google Sheets(1..n)
as the primary backend datasource delivering roster data via JSON payload to microservice clients.

Although currently configured with no authentication required(other than sharing the spreadsheet with the 
service account you need to provide), it was built with Google OIDC support(Firebase/ Cloud Identity)
in mind  