import seaborn as sns
import matplotlib.pyplot as plt
import pandas as pd



data_path = ['output/inst_exponent_bb_0.8_1.0.csv', 'output/inst_exponent_dynamic_weight_0.8_1.0.csv', 'output/inst_exponent_heuristic_0.8_1.0.csv', 'output/inst_exponent_dynamic_price_0.8_1.0.csv']
#data_path = ['output/inst_exponent_bb_0.8_1.0.csv', 'output/inst_exponent_bf_0.8_1.0.csv', 'output/inst_exponent_dynamic_weight_0.8_1.0.csv', 'output/inst_exponent_heuristic_0.8_1.0.csv', 'output/inst_exponent_dynamic_price_0.8_1.0.csv']
#data_path = ['output/inst_price_bb_0.8_1.0.csv', 'output/inst_price_bf_0.8_1.0.csv', 'output/inst_price_dynamic_weight_0.8_1.0.csv', 'output/inst_price_heuristic_0.8_1.0.csv', 'output/inst_price_dynamic_price_0.8_1.0.csv']
#data_path = ['output/inst_weight_bb_0.8_1.0.csv', 'output/inst_weight_bf_0.8_1.0.csv', 'output/inst_weight_dynamic_weight_0.8_1.0.csv', 'output/inst_weight_heuristic_0.8_1.0.csv', 'output/inst_weight_dynamic_price_0.8_1.0.csv']
#data_path = ['output/inst_ratio_bb_0.8_1.0.csv', 'output/inst_ratio_bf_0.8_1.0.csv', 'output/inst_ratio_dynamic_weight_0.8_1.0.csv', 'output/inst_ratio_heuristic_0.8_1.0.csv', 'output/inst_ratio_dynamic_price_0.8_1.0.csv']

df_bb = pd.read_csv(data_path[0])
#df_bf = pd.read_csv(data_path[1])
df_weight = pd.read_csv(data_path[1])
df_heuristic = pd.read_csv(data_path[2])
df_price = pd.read_csv(data_path[3])


sns.set(style="darkgrid")

sns_plot = sns.lineplot(x='exponent', y='duration', label='bb', data=df_bb).get_figure()

#sns_plot = sns.lineplot(x='exponent', y='duration', label='bf', data=df_bf).get_figure()

sns_plot = sns.lineplot(x='exponent', y='duration', label='dynamic-weight', data=df_weight).get_figure()

sns_plot = sns.lineplot(x='exponent', y='duration', label='dynamic-price', data=df_price).get_figure()

sns_plot = sns.lineplot(x='exponent', y='duration', label='heuristic', data=df_heuristic).get_figure()

#sns_plot = sns.lineplot(x='exponent', y='error', label='heuristic', data=df_heuristic).get_figure()


sns_plot.savefig('exponent.png')