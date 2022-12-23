#!/usr/bin/dumb-init /bin/sh
set -euo pipefail

initialize()
{
    echo "Initializing Youtube Cron microservice"
}

run_migrations()
{
    echo "Run migrations on live db"
    su-exec appuser /bin/migrate up
}

collect_process()
{
    # Get pid for app process
    APP_PID=$!

    # Wait for app to finish.
    wait "$APP_PID"
}

start_youtube_api()
{
    echo "Starting Youtube API layer"
    ls -al
    su-exec appuser /bin/api -base_path=${SRC_DIR} &

    collect_process
}

start_youtube_worker()
{
    echo "Starting Youtube Cron Workers"
    su-exec appuser /bin/worker -base_path=${SRC_DIR} &

    collect_process
}

initialize

run_migrations

if [[ "$1" == "api" ]]; then
    echo "Youtube api pods"
    start_youtube_api
else
    echo "Youtube Cron workers"
    start_youtube_worker
fi

