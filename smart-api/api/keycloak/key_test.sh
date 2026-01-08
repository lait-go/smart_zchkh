#!/bin/bash

set -e

KEYCLOAK_URL="http://localhost:9080"
REALM_NAME="smart-hcs"
CLIENT_ID="admin-cli"
USERNAME="pushkin"
PASSWORD="userpass"

# Получаем токен через Direct Access Grant
JSON_RESPONSE=$(curl -s --location --request POST "$KEYCLOAK_URL/realms/$REALM_NAME/protocol/openid-connect/token" \
  --header "Content-Type: application/x-www-form-urlencoded" \
  --data-urlencode "client_id=$CLIENT_ID" \
  --data-urlencode "username=$USERNAME" \
  --data-urlencode "password=$PASSWORD" \
  --data-urlencode "grant_type=password")

# Проверяем ошибки
if echo "$JSON_RESPONSE" | jq -e '.error' >/dev/null; then
  ERROR_MSG=$(echo "$JSON_RESPONSE" | jq -r '.error + ": " + .error_description')
  echo "❌ Ошибка получения токена: $ERROR_MSG"
  exit 1
fi

TOKEN=$(echo "$JSON_RESPONSE" | jq -r '.access_token')

# Проверяем наличие пользователя
USER_RESPONSE=$(curl -s --location --request GET \
  "$KEYCLOAK_URL/admin/realms/$REALM_NAME/users?username=$USERNAME&exact=true" \
  --header "Authorization: Bearer $TOKEN")

if echo "$USER_RESPONSE" | jq -e 'type == "array" and length > 0' >/dev/null; then
  echo "✅ Пользователь '$USERNAME' найден."
else
  echo "❌ Пользователь '$USERNAME' НЕ найден."
fi
