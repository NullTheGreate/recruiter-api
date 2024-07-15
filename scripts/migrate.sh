#!/bin/bash

# migrate.sh for MongoDB

# Default values
MIGRATIONS_DIR="./migrations"
DATABASE_URL="mongodb://localhost:27017/rec"
ACTION="up"

# Function to print usage
usage() {
    echo "Usage: $0 [-p <migrations_path>] [-d <database_url>] [-a <action>]"
    echo "  -p  Path to migrations directory (default: $MIGRATIONS_DIR)"
    echo "  -d  MongoDB connection URL (default: $DATABASE_URL)"
    echo "  -a  Migration action: up, down, or version (default: $ACTION)"
    exit 1
}

# Parse command line options
while getopts ":p:d:a:" opt; do
    case $opt in
        p) MIGRATIONS_DIR="$OPTARG" ;;
        d) DATABASE_URL="$OPTARG" ;;
        a) ACTION="$OPTARG" ;;
        \?) echo "Invalid option -$OPTARG" >&2; usage ;;
    esac
done

# Verify golang-migrate is installed
if ! command -v migrate &> /dev/null; then
    echo "Error: golang-migrate is not installed. Please install it first."
    echo "You can install it using: go install -tags 'mongodb' github.com/golang-migrate/migrate/v4/cmd/migrate@latest"
    exit 1
fi

# Perform migration action
case $ACTION in
    up)
        echo "Applying migrations..."
        migrate -path "$MIGRATIONS_DIR" -database "$DATABASE_URL" up
        ;;
    down)
        echo "Reverting last migration..."
        migrate -path "$MIGRATIONS_DIR" -database "$DATABASE_URL" down 1
        ;;
    version)
        echo "Checking current migration version..."
        migrate -path "$MIGRATIONS_DIR" -database "$DATABASE_URL" version
        ;;
    *)
        echo "Invalid action: $ACTION"
        usage
        ;;
esac

# Check for errors
if [ $? -eq 0 ]; then
    echo "Migration action '$ACTION' completed successfully."
else
    echo "Migration action '$ACTION' failed."
    exit 1
fi