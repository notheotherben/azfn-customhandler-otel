# Azure Functions Custom Handler OpenTelemetry Repro
This repository contains an extremely simple Azure Function app which leverages the
["HTTP-only" custom handler][httponly] model and demonstrates that there is a problem with
OpenTelemetry trace context propagation in the production Azure Functions runtime.

This project is currently deployed at <https://azfncustomhandlerotel.azurewebsites.net/api/hello>
should you wish to validate the behaviour without deploying your own function.

## Requirements
You will need [Go](https://go.dev) installed on your machine, as well as the Azure Functions
Core Tools. Once you have these, you will be able to build the application and deploy it (or
test it locally).

**NOTE**: I'm building and running this on Windows (both local dev and Azure Functions OS).

## Deploy to Azure
When deploying to Azure, you can confirm this behaviour by enabling/disabling the AppInsights
integration on the function app. When AppInsights is enabled, the trace context is propagated
correctly, while when it is disabled, the trace context is not propagated (i.e. it will show `<not set>`).

```bash
# Build the main handler executable
go build main.go

# Upload the function to Azure (note that you will need to have created the function app before doing this)
func azure functionapp publish $FUNCTION_APP_NAME

# Invoke the function
curl https://$FUNCTION_APP_NAME.azurewebsites.net/api/hello

# You should see the following output (while this issue remains unpatched):
# Hello World, your OpenTelemetry TraceParent is: <not set>
```

## Test Locally
When testing locally, you should build the project for your local platform and then start it using the
`func start` command. Note that you may need to update the `host.json` file to match the name of the
handler executable (`main` on Linux/MacOS systems and `main.exe` on Windows systems).

```bash
# Build the main handler executable for your local platform
go build main.go

# Start the local function app
func start

# Invoke the function
curl http://localhost:7071/api/hello

# You should see the following output:
# Hello World, your OpenTelemetry TraceParent is: 00-f70e7847ea52af240e5efff633e57cd5-baf07326de5b0eda-00
```

[httponly]: https://learn.microsoft.com/en-us/azure/azure-functions/functions-custom-handlers#http-only-function