.PHONY: all build clean run

# Имя исполняемого файла
APP_NAME := GoDatabase

# Путь к исходному файлу
SOURCE_FILE := cmd/app/main.go

# Опции компиляции
BUILD_FLAGS := -v

# Путь к Go
GO := go

all: build

build:
	@echo "Building $(APP_NAME)..."
	@$(GO) build $(BUILD_FLAGS) -o $(APP_NAME) $(SOURCE_FILE)

clean:
	@echo "Cleaning up..."
	@rm -f $(APP_NAME)

run: build
	@echo "Running $(APP_NAME)..."
	@./$(APP_NAME)