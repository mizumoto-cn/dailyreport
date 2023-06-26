# if windows CP=copy else CP=cp
ifeq ($(OS), Windows_NT)
# copy template\*.template to .
	CP=for %i in (template\*.template) do copy /Y "%i" "%~ni" & mkdir configs & for %i in (template\*.tmp) do copy /Y "%i" "configs\%~ni"
else
	CP= mkdir configs && cd template && for file in ./*.template; do cp "$$file" "../${file%.template}"; done && for file in ./*.tmp; do cp "$$file" "../configs/${file%.tmp}"; done && cd ..
endif

$(info $$CP is [${CP}])

# generate conf.proto
.PHONY: conf
conf:
	cd conf && protoc --go_out=paths=source_relative:. conf.proto

.PHONY: setup
setup:
	$(CP)
	$(MAKE) conf

.PHONY: run-cmd
run-cmd:
	cd cmd && go run main.go && cd ..