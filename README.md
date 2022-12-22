Setup MongoDB with your current IP Address. Put MongoDB API Key in .env variable called MONGO_KEY. Upload original training Data to MongoDB. Now clone repo and execute "go run main.go". Open "http://localhost:8080/index.html". Now use GUI to upload new data points, train model, and run model!



Initialize MongoDB

Install the MongoDB Database-tools suite on your local machine.
https://www.mongodb.com/docs/database-tools/installation/installation/


Import CSV Files to different Collections
mongoimport --uri $MONGO_KEY -d <database_name> --collection <collection_name> --type=csv --headerline --file <file_path>