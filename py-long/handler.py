import os
import time
import sys

def handle(event, context):

    s = os.getenv("handler_wait_duration_secs", "5")
    dur = int(s)

    sys.stderr.write('>> handler waiting for {}\n'.format(s))
    sys.stderr.flush()

    time.sleep(dur)

    sys.stderr.write('<< handler waited for {}\n'.format(s))
    sys.stderr.flush()

    return {
        "statusCode": 200,
        "body": "Hello from OpenFaaS!"
    }
