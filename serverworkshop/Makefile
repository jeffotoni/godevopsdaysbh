# Makefile

.EXPORT_ALL_VARIABLES:	

GO111MODULE=on

statik:
	@echo "\033[0;33m################################## build statik ##################################\033[0m"
	statik -src=web/
	@echo "\033[0;32mstatik generated\033[0m  \033[0;33m[ok]\033[0m "
	@echo "--------------------------------------------------------------------------------------"
	GOOS=linux go build -ldflags="-s -w" -o serverworkshop main-statik.go
	#upx serverworkshop
	upx --brute serverworkshop
	@echo "\033[0;32mserverworkshop generated\033[0m \033[0;33m[ok]\033[0m \n"

dinamic:
	@echo "--------------------------------------------------------------------------------------"
	GOOS=linux go build -ldflags="-s -w" -o serverworkshop main.go
	upx serverworkshop
	#upx --brute serverworkshop
	@echo "\033[0;32mserverworkshop generated\033[0m \033[0;33m[ok]\033[0m \n"	