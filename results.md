# Results

## Overview
nip-checker provides information on current failure rates of NIP transactions.

## Architecture
nip-checker is a http and grpc server. It runs a background job that periodically calls the NIP transactions api to get live failure rates. Rates are stored in a local json file(```data/data.json```) and are updated upon each call.  

## Assumptions
1. ```todayFailureRateInward``` and ```todayFailureRateOutward``` are the quantities of interest.
2. A threshold of 0.5% for both ```todayFailureRateInward``` and ```todayFailureRateOutward```. Effective failure rate is the maximum of the two.

