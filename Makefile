build:
	@go build -o example
run: build
	@./example
git:
	@del example && git add . && git commit -m "feat: learning" && git push