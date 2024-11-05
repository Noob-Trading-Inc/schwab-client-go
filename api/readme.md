brew install swagger-codegen

swagger-codegen generate -i trader.json -l go -Dmodels -o ./generated
cat ./generated/model_*.go > models.go
rm -r ./generated

swagger-codegen generate -i marketdata.json -l go -Dmodels -o ./generated
cat ./generated/model_*.go > models.go
rm -r ./generated
