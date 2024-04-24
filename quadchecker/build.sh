#!/bin/bash

# Directory to move the original Go files to
target_dir="quad"

# Ensure the target directory exists
mkdir -p "$target_dir"

# List of Go files to compile
go_files=("quadA.go" "quadB.go" "quadC.go" "quadD.go" "quadE.go" "quadchecker.go")

# Loop over the Go files
for go_file in "${go_files[@]}"; do
    # Extract the base name for the executable
    base_name=$(basename "$go_file" .go)

    # Compile the Go file into an executable
    echo "Compiling $go_file to $base_name"
    go build -o "$base_name" "$go_file"
    # Set the permissions for the executable
    chmod 755 "$base_name"

    # Check if the current file is not main.go
    if [[ "$go_file" != "main.go" ]]; then
        # Move the original Go file to the target directory
        echo "Moving $go_file to $target_dir/"
        mv "$go_file" "$target_dir/"
        # Set the permissions for the source file in the target directory
        chmod 755 "$target_dir/$go_file"
    fi
done

# Specifically set permissions for main.go without moving it
if [[ -f "main.go" ]]; then
    echo "Setting permissions for main.go"
    chmod 644 "main.go"
fi

echo "Compilation and moving completed."

# Create the go.mod file in the same directory as the script
module_name="main_module"
if [[ ! -f "go.mod" ]]; then
    echo "Creating go.mod with module name $module_name"
    go mod init "$module_name"
fi

# Remove the 'quad' directory and its contents
echo "Removing $target_dir directory and its contents..."
rm -r "$target_dir"
echo "$target_dir directory removed."
echo "Process completed."
