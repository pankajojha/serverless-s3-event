make build
#local start
serverless offline start

#aws or #azure deploy
serverless deploy


curl -X POST https://abcCloud.execute-api.us-east-1.amazonaws.com/dev/world -d 'Hello, world!'
curl -X POST https://abcCloud.execute-api.us-east-1.amazonaws.com/dev/event -d 'Hello, world!'
curl -X POST http://127.0.0.1:3000/upload -H "X-Autherization:8C478DB221879DD93D1EA0F7488CEA4D" -d '{Hello: world}'

curl -X POST http://127.0.0.1:3000/upload -H "X-Auth:Tester12345" -d '{Hello: world}'


curl -X POST https://vkdozk2nbb.execute-api.us-east-1.amazonaws.com/dev/upload -H "x-auth:Tester12345" -H "X-Auth:Tester12345" -d '{Hello: world}'


https://www.powershell.amsterdam/2018/05/22/beginnings-in-golang-and-aws-part-iii-uploading-to-s3-contd/
https://medium.com/@tuomovee/go-serverless-with-sam-860c62f63ba4

dynamodDB
    https://yos.io/2018/02/08/getting-started-with-serverless-go/

# local deploy
npm install -g aws-sam-local
# create template.yaml
sam local start-api


## template to deploy in S3 
<!-- aws s3 mb s3://feelings-lambdas
sam package --template-file template.yaml --s3-bucket feelings-lambdas --output-template-file packaged.yaml
sam deploy --template-file packaged.yaml --stack-name feelings --capabilities CAPABILITY_IAM -->

sam package \
    --template-file template.yaml \
    --output-template-file packaged-output.yaml \
    --s3-bucket pci-1

sam publish \
    --template packaged-output.yaml \
    --region us-east-1    