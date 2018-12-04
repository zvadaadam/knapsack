import subprocess
import pandas as pd
import seaborn as sns

sns.set(style="darkgrid")

#subprocess.call("generator.sh", shell=True)

df = pd.read_csv('data.csv')

sns_plot = sns.lineplot(x='x', y='y', data=df).get_figure()
sns_plot.savefig("plot.png")