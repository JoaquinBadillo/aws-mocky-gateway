# aws-mocky-gateway
Docker compose service that mocks lambda invocations behind an api gateway.
Lambdas must be containerized and included in the compose file.
To register paths add a NGINX location configuration.
