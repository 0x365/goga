import matplotlib.pyplot as plt
import numpy as np
import os
import csv


def csv_input(file_name):
    with open(file_name, "r") as f:
        read_obj = csv.reader(f)
        output = []
        for row in read_obj:
            try:
                temp = []
                for item in row:
                    temp.append(eval(row))
                output.append(temp)
            except:
                try:
                    output.append(eval(row))
                except:
                    try:
                        output.append(row)
                    except:
                        output = row    
    f.close()
    return output


# Create save location if not already exist
save_location = os.path.join(os.path.dirname(os.path.abspath(__file__)), "data")

print("Graphing GA Fitness")

ga_data = np.asarray(np.array(csv_input(save_location+"/ga_results.csv")), dtype="float")

plt.plot(ga_data[:,0])
plt.yscale("log")
# plt.axis('equal')
plt.savefig(save_location+"/ga_fitness.png")
# plt.show()
plt.show()
