api = "./api/base.api"
api_output = "./service"
api_template = ".goctl/1.6.3"

generate_api:
	goctl api go --home ${api_template} --api ${api} --dir ${api_output}

generate_input = "./ent/schema"
generate_output = "./ent"

# 出现root package or module was not found for: gen/entschema 必须先任意创建一个go文件占位
sql:
	go install entgo.io/ent/cmd/ent@v0.13.1
	ent generate --feature sql/upsert --target ${generate_output} ${generate_input}
