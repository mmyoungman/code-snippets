import boto3

bucketName = 'mark-test-bucket-2983289'

s3 = boto3.client('s3')
resp = s3.create_bucket(Bucket=bucketName, CreateBucketConfiguration={'LocationConstraint': 'eu-west-2'})

print(str(resp))
