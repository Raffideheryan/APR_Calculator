# APR_Calculator
This program calculates the Annual Percentage Rate (APR) for blockchain validators based on data retrieved from a beacon node. It automates the process by fetching necessary data from blockchain endpoints and performing calculations to determine the APR.

Features
Fetches data from the provided blockchain URLs.
Calculates the APR based on the head slot and other network parameters.
Simple to configure: users only need to update URLs for their specific blockchain environment.

How It Works
The program connects to the specified beacon node endpoint and retrieves the necessary data (such as the head slot or validator performance).
It processes the data to calculate the APR for validators in the network.
The output provides an accurate APR, helping users evaluate validator performance and network efficiency.

Configuration
Open the program file.
Locate the following two placeholders:
Blockchain URL: The base URL for your beacon node API.
Head Slot URL: The specific API endpoint to fetch the head slot (usually /eth/v1/beacon/headers).
Replace these URLs with the appropriate values for your blockchain network.

Running the Program
1)go build -o apr-calculator
2)./apr-calculator

