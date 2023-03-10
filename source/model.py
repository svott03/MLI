# -*- coding: utf-8 -*-
"""Meal_Processing_Final.ipynb

Automatically generated by Colaboratory.

Original file is located at
    https://colab.research.google.com/drive/1pA2aTlZ8cbeyL29WIEpk-XuC-GfTyMGr
"""

# Commented out IPython magic to ensure Python compatibility.
import pandas as pd
import numpy as np

center_info = pd.read_csv('./files/fulfilment_center_info.csv')
meal_info = pd.read_csv('./files/meal_info.csv')
test_data = pd.read_csv('./files/test.csv')
train_data = pd.read_csv('./files/train.csv')

merge1 = pd.merge(train_data, center_info, how='inner', on='Center_id')
df = pd.merge(merge1, meal_info, how='inner', on='Meal_id')
df = df.sort_values(by=['Week'])
df = df.dropna()

cat_var = ['center_type',
 'category',
 'cuisine']

num_cols=['Center_id',
'Meal_id',
'Checkout_price',
'Base_price',
'Emailer_for_promotion',
'Homepage_featured',
'Num_orders',
'city_code',
'region_code',
'op_area']
colors=['#b84949', '#ff6f00', '#ffbb00', '#9dff00', '#329906', '#439c55', '#67c79e', '#00a1db', '#002254', '#5313c2', '#c40fdb', '#e354aa']
ts_tot_orders = df.groupby(['Week'])['Num_orders'].sum()
ts_tot_orders = pd.DataFrame(ts_tot_orders)

df_ = df.copy()
for i in cat_var:
    df_[i] = pd.factorize(df_[i])[0]

from sklearn.model_selection import train_test_split 
# %matplotlib inline

X = df_.drop(['Num_orders'], axis=1).values
y = df_['Num_orders'].values

X_train, X_test, y_train, y_test = train_test_split(X, y, test_size=0.2, random_state=0)


import xgboost as xgb
from sklearn.metrics import mean_squared_error
import pickle

from sklearn.model_selection import GridSearchCV

# Various hyper-parameters to tune
xgb1 = xgb.XGBRegressor()
parameters = {'nthread':[4], #when use hyperthread, xgboost may become slower
              'objective':['reg:squarederror'],
              'learning_rate': [.03, 0.05, .07], #so called `eta` value
              'max_depth': [5, 6, 7],
              'min_child_weight': [4],
              'subsample': [0.7],
              'colsample_bytree': [0.7],
              'n_estimators': [500]}

xgb_grid = GridSearchCV(xgb1,
                        parameters,
                        cv = 2,
                        n_jobs = 5,
                        verbose=True)

xgb_grid.fit(X_train,
         y_train)

xgb_grid.best_estimator_
y_pred = xgb_grid.predict(X_test)

mse=mean_squared_error(y_test, y_pred)

filename = './files/finalized_model.sav'
pickle.dump(xgb_grid, open(filename, 'wb'))

print(np.sqrt(mse))