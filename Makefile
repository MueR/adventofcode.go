YEAR := $$(shell /bin/date +'%Y')
DAY := $$(shell /bin/date +'%d')


# https://gist.github.com/prwhite/8168133
help: ## Show this help
	@ echo 'Usage: make <target>'
	@ echo
	@ echo 'Available targets:'
	@ grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "  \033[36m%-20s\033[0m %s\n", $$1, $$2}'

skeleton: ## make skeleton main(_test).go files, optional: $DAY and $YEAR
	@ if [ -n $$DAY && -n $$YEAR ]; then \
		go run scripts/cmd/skeleton/main.go -day $(DAY) -year $(YEAR) ; \
	elif [ -n $$DAY ]; then \
		go run scripts/cmd/skeleton/main.go -day $(DAY); \
	else \
		go run scripts/cmd/skeleton/main.go; \
	fi

test: ## run tests
	@ go test ./$$(date +'%Y')/day$$(date +'%d') && echo "Looks good to me!"

run: ## run current day
	@ go run ./$$(date +'%Y')/day$$(date +'%d')
