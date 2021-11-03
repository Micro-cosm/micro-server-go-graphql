# micro-server-go-graphql

A Go-based Graphql service configured to run on GCP/Docker and connect to Google Sheets(1..n)
as the primary backend datasource.  Great for delivering roster or user data via JSON payload to
mifedom clients.

Although currently configured with no authentication required(other than sharing the spreadsheet with the 
service account you need to provide), it was built with Google OIDC support(Firebase/ Cloud Identity)
in mind.