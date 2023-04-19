import os
import openai
from dotenv import load_dotenv

# Load the API key from the .env file
load_dotenv()
api_key = os.getenv("OPENAI_API_KEY")

# Authenticate the API client
openai.api_key = api_key

# Define the prompt
prompt = "Please write me a haiku about Montreal bagels."

# Call the API
response = openai.Completion.create(
    engine="text-davinci-002",
    prompt=prompt,
    max_tokens=140,
    n=1,
    stop=None,
    temperature=0.35,
)

# Print the generated text
generated_text = response.choices[0].text
print(generated_text)
