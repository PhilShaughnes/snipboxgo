# https://cheatography.com/linux-china/cheat-sheets/justfile/

# set dotenv-load := true

default:
	@just --list

# watch and run a go file
watch:
	fd -tf . | entr -rc go run .

# watch and run a go file
wtest:
	fd -tf . | entr -c go test .
