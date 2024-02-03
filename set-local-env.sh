#!/bin/bash

# List of folder names
folders=("backend" "frontend")

# Iterate through each folder
for folder in "${folders[@]}"; do
  # Copy .env.local to .env
  cp "$folder/.env.local" "$folder/.env"
done

echo "Copied .env.local to .env in ${folders[@]}"