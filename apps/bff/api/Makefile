#gen-code:
#	57ctl api go -api ./def/bff.api -dir . --remote="https://github.com/v3nooonn/trytry-template.git"
#
#gen-bff:
#	57ctl api -o ./def/bff.api --remote="https://github.com/v3nooonn/trytry-template.git"
#
#gen-api:
#	57ctl api -o ./def/apis/$(api).api --remote="https://github.com/v3nooonn/trytry-template.git"

fmt:
	goctl api format --dir .

bff:
	goctl api -o ./def/bff.api

def:
	goctl api -o ./def/apis/$(api).api

code:
	goctl api go -api ./def/bff.api -dir .
