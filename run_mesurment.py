import subprocess
import os
import configparser
import pandas as pd
import seaborn as sns
import matplotlib.pyplot as plt


ABS_PATH = os.path.abspath(os.path.dirname(__file__))


def run_instance(num_repeat, optimized, algorithm_name, num_items, num_instances, capacity_weights_ratio, max_weight, max_price, exponent, type):

    df = pd.DataFrame()

    for i in range(num_repeat):

        subprocess.call(['./generator.sh',
                     '-o', f'{optimized}',
                     '-a', f'{algorithm_name}',
                     '-n', f'{num_items}',
                     '-N', f'{num_instances}',
                     '-m', f'{capacity_weights_ratio}',
                     '-W', f'{max_weight}',
                     '-C', f'{max_price}',
                     '-k', f'{exponent}',
                     '-d', f'{type}'])

        new_df = pd.read_csv('instance_data.csv')

        if i != 0:
            df['duration'] = df['duration'] + new_df['duration']
            #df['error'] = df['error'] + new_df['error']
        else:
            df = new_df



    df['duration'] = df['duration']/num_repeat
    #df['error'] = df['error']/NUM_REPEAT


    print(df.head())

    new_filename = 'output/' + 'inst'+ '_' + str(optimized) + '_' + str(algorithm_name) + '_' + str(capacity_weights_ratio) + '_' + str(exponent)

    sns.set(style="darkgrid")
    sns_plot = sns.lineplot(x=optimized, y='duration', data=df).get_figure()
    sns_plot.savefig(new_filename + '.png')

    # clear plot
    plt.clf()

    df.to_csv(new_filename + '.csv', sep=',')


if __name__ == "__main__":

    #configs = ['config_bb_weight.cfg', 'config_bb_price.cfg', 'config_bb_ratio.cfg', 'config_bb_exp.cfg']
    configs = ['config_bb_exp.cfg']
    #configs = ['config_bb_weight.cfg']

    for i in range(len(configs)):

        cfg = configparser.ConfigParser()
        cfg.read(ABS_PATH + '/config/' + configs[i])

        num_repeat = cfg.getint('params', 'num_repeat')
        optimized = cfg.get('params', 'optimized')
        algorithm_name = cfg.get('params', 'algorithm_name')
        num_items = cfg.getint('params', 'num_items')
        num_instances = cfg.getint('params', 'num_instances')
        capacity_weights_ratio = cfg.getfloat('params', 'capacity_weights_ratio')
        max_weight = cfg.getint('params', 'max_weight')
        max_price = cfg.getint('params', 'max_price')
        exponent = cfg.getfloat('params', 'exponent')
        type = cfg.getint('params', 'type')

        run_instance(num_repeat, optimized, algorithm_name, num_items, num_instances, capacity_weights_ratio, max_weight, max_price, exponent, type)