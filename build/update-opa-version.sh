#!/usr/bin/env bash
# Script to revendor OPA. Add and Commit changes if needed

# Update the OPA verison in glide.yaml
sed -i '' "/opa/{N;s/version: .*/version: $1/;}" glide.yaml

# Check if OPA version has changed
git diff-index --quiet HEAD --
if [ $? -ne 0 ]; then 
  # run glide update
  glide up -v

  # add and commit changes
  git add . && git commit -s -m "Updated OPA version to $1" 
fi 
