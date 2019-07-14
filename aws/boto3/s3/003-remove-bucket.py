import boto3

bucketName = 'mark-test-bucket-2983289'

s3 = boto3.resource('s3')
bucket = s3.Bucket(bucketName)
resp = bucket.delete()

print(str(resp))
