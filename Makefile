
# Variables
BINARY_NAME := main
BUILD_FLAGS := -gcflags="all=-N -l"

# Default target
all: build

# Build with debug flags
build:
	go build $(BUILD_FLAGS) -o $(BINARY_NAME) .

# Run the debug binary directly
run: build
	./$(BINARY_NAME)

dbg:
	make clean && make build && make run

# Clean up
clean:
	rm -f $(BINARY_NAME)
