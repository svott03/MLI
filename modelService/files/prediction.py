import pickle
import pandas as pd
from sklearn.metrics import mean_squared_error
import numpy as np

# Prepare Data
predict_data = pd.read_csv('./predict.csv')
center_info = pd.read_csv('./fulfilment_center_info.csv')
meal_info = pd.read_csv('./meal_info.csv')

merge1 = pd.merge(predict_data, center_info, how='inner', on='center_id')
df = pd.merge(merge1, meal_info, how='inner', on='meal_id')

cat_var = ['center_type',
 'category',
 'cuisine']

df_ = df.copy()
for i in cat_var:
    df_[i] = pd.factorize(df_[i])[0]

X = df_.drop(['num_orders'], axis=1).values
y_test = df_['num_orders'].values

# Predict
filename = './finalized_model.sav'
loaded_model = pickle.load(open(filename, 'rb'))
y_pred = loaded_model.predict(X)
mse=mean_squared_error(y_test, y_pred)
print(y_pred)