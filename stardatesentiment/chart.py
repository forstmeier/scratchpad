import json
import matplotlib.pyplot as plt


with open('output.json') as json_file:
	data = json.load(json_file)

	x_scatter = []
	y_scatter = []

	for value in data.values():
		if bool(value):
			x_line = []
			y_line = []

			for score in value['scores']:
				x_line.append(score['place'])
				y_line.append(score['score'])

			plt.plot(x_line, y_line)

	plt.savefig('line.png')