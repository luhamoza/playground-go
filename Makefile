build:
	@go build -o example
run: build
	@./example
git:
	@git add . && git commit -m "feat: learning" && git push