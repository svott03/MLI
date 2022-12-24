# Machine Learning Interface 🤖📊
Access and train your deployed model with our nice Interface!

## Usage:
If you want to train your model on new data points without interacting directly with source code, use our website interface.

## Setup:
1. Clone our Repository and install python dependencies
```bash
cd~
git clone https://github.com/svott03/MLI.git
cd MLI
make
```

2. We require 2 source code files: model.py for training and prediction.py for prediction.

- Update model.py to only output training results (txt).
- Update predcition.py to only output prediction results (txt).
- Place your prediction.py in ~/MLI/modelService/files.

3. We are using MongoDB. Follow a quickstart guide to create a cluster and a database (<databaseName>) with your local IP address. To connect with your database, go to Connect, choose Connect with your application and copy your Mongo URI. Replace <password> with your account password. This URI string will be used later. Create a collection with name <trainingCollectionName>.

Now create the following environment variables for your URI, database name, and training collection.
```bash
export MONGO_KEY="<URI>"
export MongoDB="<databaseName>"
export MongoTrainData="<trainingCollectionName>"
```

Install the MongoDB Database-tools suite on your local machine.
```bash
brew install mongodb-database-tools
```

Import your training data (from <training_file_path>) with the following:
```bash
mongoimport --uri $MONGO_KEY -d $MongoDB --collection $MongoTrainData --type=csv --headerline --file <training_file_path>
```

## Run:
1. Open 2 terminal instances and export env variables to both instances. Now start both servers.
```bash
cd~/MLI/server
go run main.go
```

```bash
cd~/MLI/modelService
go run main.go
```

2. Interact with website

Open "http://localhost:8080". Explanations for each piece of functionality are on the website. Enjoy!


## How:
We have the standard client-server interaction between your browser and our server directory. Our server dynamically updates html, sends data to the db, and sends workload requests to our modelService. Our modelService then accesses the db, runs your source code to train and use the model and sends results back to the server. We are using the gin-gonic framework to send requests in golang.