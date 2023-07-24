## Deleting Table

table_name="$1"

local_dynamodb="http://localhost:8000"

echo -e "\nTables: "

aws dynamodb list-tables --endpoint-url "${local_dynamodb}"

echo -e "Deleting Tables.. \n "

aws dynamodb delete-table --table-name "${table_name}" --endpoint-url "${local_dynamodb}"
