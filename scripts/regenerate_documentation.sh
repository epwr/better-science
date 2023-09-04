#!/bin/bash

# Directory to store generated documentation 
doc_dir="doc"

# List all packages in the module
packages=$(go list ./...)

# Create the documentation directory if it doesn't exist
mkdir -p "$doc_dir"

# Remove all existing documentation
rm "${doc_dir}/*"

# Loop over packages and generate documentation
for pkg in $packages; do
    # Generate HTML documentation for the package and save it to the doc_dir
    go doc -all -u "$pkg" > "${doc_dir}/$(basename $pkg).html"
done

echo "Documentation generated for all packages in the module."
