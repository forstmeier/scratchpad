import json
import re
import requests

from bs4 import BeautifulSoup, Tag, NavigableString
from nltk.sentiment import SentimentIntensityAnalyzer


# http://www.chakoteya.net/NextGen/episodes.htm

output = {}

sia = SentimentIntensityAnalyzer()

episode_url = "http://www.chakoteya.net/NextGen/{0}.htm"
num_range = range(101, 278)
for i, num in enumerate(num_range, start=1):
	html_text = requests.get(episode_url.format(num)).text

	soup = BeautifulSoup(html_text, 'html.parser')

	output[i] = {}

	title_element = soup.find_all('font', color='#2867d0', size='5')
	for title in title_element:
		output[i]['title'] = title.text.replace('\n','').replace('\r',' ')

	scores = []

	results = soup.td
	if results == None:
		continue
	else:
		children_count = len(list(results.children))
		for j, child in enumerate(results.children):
			if isinstance(child, Tag):
				if child.string != None and 'log, ' in child.string:
					child_string = child.string.replace('\n','').replace('\r',' ')
					sentiment_score = sia.polarity_scores(child_string)['compound']
					score_placement = j / children_count

					score = {
						'text': child_string,
						'score': sentiment_score,
						'place': score_placement,
					}

					scores.append(score)

	output[i]['scores'] = scores


with open('output.json', 'w') as output_json:
	json.dump(output, output_json)