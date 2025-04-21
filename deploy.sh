#!/bin/bash

set -e  # Exit on error

if [ "$#" -ne 2 ]; then
  echo "Usage: $0 <remote-user> <remote-host>"
  exit 1
fi

REMOTE_USER="$1"
REMOTE_HOST="$2"
APP_NAME="vehicleapi"
REMOTE_PATH="/opt/vehicleapi"
SERVICE_NAME="vehicleapi"
BUILD_PATH="./cmd/vehicleapi"

echo "🔨 Building Go binary..."
go build -o $APP_NAME $BUILD_PATH

scp $APP_NAME $REMOTE_USER@$REMOTE_HOST:$TEMP_REMOTE_PATH

echo "🚀 Moving binary and restarting service..."
echo "ssh command: sudo mv /home/$REMOTE_USER/$APP_NAME $REMOTE_PATH/$APP_NAME"

ssh $REMOTE_USER@$REMOTE_HOST "sudo mv /home/$REMOTE_USER/$APP_NAME $REMOTE_PATH/$APP_NAME"
echo "File was correctly moved."
echo "ssh command: sudo /usr/bin/systemctl restart $SERVICE_NAME && sudo /usr/bin/systemctl status $SERVICE_NAME --no-pager"
ssh $REMOTE_USER@$REMOTE_HOST "sudo /usr/bin/systemctl restart $SERVICE_NAME && sudo /usr/bin/systemctl status $SERVICE_NAME --no-pager"

echo "✅ Deployment complete."
