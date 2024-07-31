URL=https://esi.evetech.net/latest/swagger.json

latest: download generate

download:
	wget -O esi.json $(URL)

generate:
	swagger-codegen generate -i esi.json -l go -o swagger
