import pandas as pd
from sklearn.feature_extraction.text import TfidfVectorizer
from sklearn.metrics.pairwise import linear_kernel
from flask import Flask, request, jsonify

app = Flask(__name__)
def load_data(file_path):
    df = pd.read_json(file_path, lines=True)
    return df

#Load the review dataset
df = load_data('data/Video_Games.jsonl')

#Prepare the TF-DF matrix for the review text
tfidf = TfidfVectorizer(stop_words='english')

#Replacing NaN with empty string to aviod errors
df['text'] = df['text'].fillna('')

#Fit and Transform the review text to a TF-DF matrix
tfidf_matrix = tfidf.fit_transform(df['text'])

# Compute the cosine similarity matrix
cosine_sim = linear_kernel(tfidf_matrix, tfidf_matrix)

# Build a reverse mapping od indices and product IDs
indices = pd.Series(df.index,index=df['asin']).drop_duplicates()

# Function to get recommendations based on product ASIN
def get_recommendations(asin, cosine_sim=cosine_sim):
    #get the index of the product that matches the ASIN
    idx = indices[asin]

    #Get the pairwise similarity scores of all products with that product
    sim_scores = list(enumerate(cosine_sim[idx]))

    # Sort the products based on similar products
    sim_scores = sorted(sim_scores, key=lambda x: x[1], reverse=True)

    # Get the scores of the 5 most similar product
    sim_scores = sim_scores[1:6]

    #Get the product indices
    product_indices = [i[0] for i in sim_scores]

    #Return the top 5 most similar products
    return df['asin'].iloc[product_indices].tolist()

@app.route('/recommend/<string:asin>', methods=['GET'])
def recommend(asin):
    try:
        recommendations = get_recommendations(asin)
        return jsonify({'recommendations':recommendations})
    except KeyError:
        return jsonify({'error':'asin not found'}), 404
if __name__ == '__main__':
    # Start the Flask app on port 5000
    app.run(debug=True, host='0.0.0.0', port=5000)