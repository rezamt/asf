## Build all binaries

build:
	go build -o bin/asf-server cmd/asf/server.go


#-------------------------
# Target: swagger.validate
#-------------------------
swagger.validate:
	swagger validate pkg/swagger/swagger.yml


#-------------------------
# Target: swagger.doc
#-------------------------
swagger.doc:
	python /Users/reza/projects/nubot/swagger-yaml-to-html/swagger-yaml-to-html.py < pkg/swagger/swagger.yml > doc/index.html