YEAR ?= $(shell date +'%Y')
DAY ?= $(shell date +'%d')

# https://gist.github.com/prwhite/8168133
help: ## Show this help
	@ echo 'Usage: make <target>'
	@ echo
	@ echo 'Available targets:'
	@ grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "  \033[36m%-20s\033[0m %s\n", $$1, $$2}'

skeleton: ## make skeleton main(_test).go files, optional: $DAY and $YEAR
	go run scripts/cmd/skeleton/main.go -day $(DAY) -year $(YEAR) ; \

test: ## run tests
	@ go test -v ./$(YEAR)/day$(DAY) && echo "\n✅  Looks good to me!"

run: ## run current day
	@ echo "              _                 _            __    _____          _      "
	@ echo "     /\      | |               | |          / _|  / ____|        | |     "
	@ echo "    /  \   __| |_   _____ _ __ | |_    ___ | |_  | |     ___   __| | ___ "
	@ echo "   / /\ \ / _\` \ \ / / _ \ '_ \| __|  / _ \|  _| | |    / _ \ / _\` |/ _ \\"
	@ echo "  / ____ \ (_| |\ V /  __/ | | | |_  | (_) | |   | |___| (_) | (_| |  __/"
	@ echo " /_/    \_\__,_| \_/ \___|_| |_|\__|  \___/|_|    \_____\___/ \__,_|\___|"
	@ echo "                                                                         "
	@ echo "                              $(YEAR) day $(DAY)                                 "
	@ echo ""

	@ go run ./$(YEAR)/day$(DAY)
