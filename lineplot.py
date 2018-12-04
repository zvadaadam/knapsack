import subprocess
import pandas as pd
import seaborn as sns

sns.set(style="darkgrid")

NUM_REPEAT = 50
algorithm_name = 'bb'
num_items = 25
num_instances = 15
capacity_weights_ratio = 0.2
max_weight = 500
max_price = 500
exponent = 1
type = 0



for i in range(NUM_REPEAT):

    subprocess.call(['./generator.sh',
                     '-a', f'{algorithm_name}',
                     '-n', f'{num_items}',
                     '-N', f'{num_instances}',
                     '-m', f'{capacity_weights_ratio}',
                     '-W', f'{max_weight}',
                     '-C', f'{max_price}',
                     '-k', f'{exponent}',
                     '-d', f'{type}'])

    new_df = pd.read_csv('data_1.csv')

    if i != 0:
        df['duration'] = df['duration'] + new_df['duration']
        #df['error'] = df['error'] + new_df['error']
    else:
        df = new_df



df['duration'] = df['duration']/NUM_REPEAT
#df['error'] = df['error']/NUM_REPEAT



print(df.head())

sns_plot = sns.lineplot(x='instance', y='duration', data=df).get_figure()
sns_plot.savefig("plot.png")