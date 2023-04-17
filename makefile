MAIN_FILE = "./cmd/futaba/main.go"
OUT_FILE = "futaba.exe"

deps:
	go get github.com/joho/godotenv
	go get github.com/bwmarrin/discordgo

run:
	go run ${MAIN_FILE}

build:
	go mod tidy
	go build -o ${OUT_FILE} ${MAIN_FILE}
