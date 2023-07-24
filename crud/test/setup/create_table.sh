## Creating Table
local_dynamodb="http://localhost:8000"

echo "Creating Table in local dynamodb: ${local_dynamodb}"

aws dynamodb create-table \
    --table-name UserTeste \
    --attribute-definitions \
        AttributeName=ID,AttributeType=S \
    --key-schema \
        AttributeName=ID,KeyType=HASH \
    --provisioned-throughput \
        ReadCapacityUnits=5,WriteCapacityUnits=5 \
    --endpoint-url "${local_dynamodb}"



echo -e "\nTables: "

aws dynamodb list-tables --endpoint-url "${local_dynamodb}"

# aws dynamodb delete-table --table-name UserTeste "${local_dynamodb}"
