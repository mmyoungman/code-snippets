import os
import logging
import json
import hashlib
import hmac
import boto3
from botocore.vendored import requests

# Other AWS services used with this:
# - API Gateway with integration response mapping
# - System Manager Parameter Store for signing secret and bot token (encrypted with KMS key)

def lambda_handler(event, context):
    if "challenge" in event:
        return event['body']['challenge']
        
    if "bot_id" in event['body']['event']:
        logging.warn("Ignore bot event")
        return

    # Get parameters
    parameters = boto3.client('ssm').get_parameters(
        Names= ['slack-illeist-signing-secret', 'slack-illeist-bot-token'],
        WithDecryption=True)

    for parameter in parameters['Parameters']:
        if parameter['Name'] == 'slack-illeist-signing-secret':
            slack_signing_secret = bytes(parameter['Value'], 'utf-8')
        elif parameter['Name'] == 'slack-illeist-bot-token':
            bot_token = parameter['Value']

    # Create basestring with no spaces    
    timestamp = event['headers']['X-Slack-Request-Timestamp']
    bodystring = json.dumps(event['body'], separators=(',', ':'))
    basestring = f'v0:{timestamp}:{bodystring}'.encode('utf-8')

    # Create and verify signiture
    my_sig = 'v0=' + hmac.new(slack_signing_secret, basestring, hashlib.sha256).hexdigest()
    if not hmac.compare_digest(my_sig, event['headers']['X-Slack-Signature']):
        logging.warn(f"Verification failed. My sig: {my_sig}")
        return
    
    # Send request back to Slack
    headers = { 
        'Content-Type': 'application/json; charset=utf-8',
        'Authorization': f'Bearer {bot_token}'
    }
    data = { 
        'channel': event['body']['event']['channel'],
        'text': event['body']['event']['text'][::-1]
    }
    r = requests.post("https://slack.com/api/chat.postMessage", json=data, headers=headers)

    # TODO: Verify response is as expected
    #print(r.content)
