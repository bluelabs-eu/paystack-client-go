#!/bin/bash

# Use the YQ environment variable if set, otherwise default to 'yq'
YQ_CMD=${YQ:-yq}

# Check if yq is installed
if ! command -v $YQ_CMD &> /dev/null; then
    echo "Error: yq is not installed. Install it from https://github.com/mikefarah/yq."
    exit 1
fi

# Check arguments
if [ "$#" -ne 3 ]; then
    echo "Usage: ./modify_openapi_schema.sh <file_path> <target_path> <schema_type>"
    exit 1
fi

FILE_PATH=$1
TARGET_PATH=$2
SCHEMA_TYPE=$3

# Check if the file exists
if [ ! -f "$FILE_PATH" ]; then
    echo "Error: File '$FILE_PATH' not found."
    exit 1
fi


# Check if the target path exists in the OpenAPI file
if ! $YQ_CMD eval ".paths.\"$TARGET_PATH\"" "$FILE_PATH" &> /dev/null; then
    echo "Error: Path '$TARGET_PATH' not found in the OpenAPI file."
    exit 1
fi

# Check if the POST method exists for the target path
if ! $YQ_CMD eval ".paths.\"$TARGET_PATH\".post" "$FILE_PATH" &> /dev/null; then
    echo "Error: POST method not found for path '$TARGET_PATH'."
    exit 1
fi

# Check if the schema type exists in the requestBody content
if ! $YQ_CMD eval ".paths.\"$TARGET_PATH\".post.requestBody.content.\"$SCHEMA_TYPE\"" "$FILE_PATH" &> /dev/null; then
    echo "Error: Schema type '$SCHEMA_TYPE' not found in the requestBody."
    exit 1
fi

# Keep only the specified schema type
$YQ_CMD eval -i ".paths.\"$TARGET_PATH\".post.requestBody.content = {\"$SCHEMA_TYPE\": .paths.\"$TARGET_PATH\".post.requestBody.content.\"$SCHEMA_TYPE\"}" "$FILE_PATH"

echo "Successfully modified '$FILE_PATH' to keep only '$SCHEMA_TYPE' for path '$TARGET_PATH'."