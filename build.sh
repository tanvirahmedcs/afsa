#!/bin/bash

# Build script for DeathNote on macOS and Linux
# This script builds the deathnote tool for the current system

set -e

# Colors for output
GREEN='\033[0;32m'
BLUE='\033[0;34m'
RED='\033[0;31m'
NC='\033[0m' # No Color

echo -e "${RED}========================================${NC}"
echo -e "${RED}DeathNote Build Script${NC}"
echo -e "${RED}========================================${NC}"

# Check if Go is installed
if ! command -v go &> /dev/null; then
    echo "❌ Go is not installed. Please install Go 1.21 or later."
    echo "   Visit: https://golang.org/dl"
    exit 1
fi

# Display Go version
GO_VERSION=$(go version | awk '{print $3}')
echo -e "${GREEN}✓ Go version: $GO_VERSION${NC}"

# Create output directory if needed
mkdir -p build

echo -e "\n${GREEN}Downloading dependencies...${NC}"
go mod download

echo -e "\n${GREEN}Building DeathNote...${NC}"
go build -o deathnote .

echo -e "\n${GREEN}✓ Build successful!${NC}"
echo -e "${GREEN}✓ Binary created: ./deathnote${NC}"

# Make it executable
chmod +x deathnote

echo -e "\n${GREEN}Quick Test:${NC}"
./deathnote --help | head -5

echo -e "\n${RED}========================================${NC}"
echo -e "${RED}Installation Instructions:${NC}"
echo -e "${RED}========================================${NC}"
echo ""
echo "1. Use locally:"
echo "   ./deathnote dns example.com"
echo ""
echo "2. Install globally (optional):"
echo "   sudo cp deathnote /usr/local/bin/"
echo "   deathnote dns example.com"
echo ""
echo "3. Run tests:"
echo "   ./deathnote dns google.com"
echo "   ./deathnote ip 8.8.8.8"
echo "   ./deathnote firewall status"
echo "   ./deathnote waf github.com"
echo ""
