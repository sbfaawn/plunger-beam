docker-compose -f .\configs\docker-compose.yaml -p plunger-beam up -d

go run .\cmd\apps\main.go --env=development